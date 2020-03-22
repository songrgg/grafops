package grafana

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"strings"

	"github.com/songrgg/grafops/pkg/simplejson"
	"github.com/songrgg/sdk"
)

type UpdateConfig struct {
	APIUrl       string `json:"apiUrl"`
	TemplateSlug string `json:"templateSlug"`
	BasicAuth    string `json:"basicAuth"`
}

type Var struct {
	Name   string `json:"name"`
	Values []Val  `json:"values"`
}

type Val struct {
	Value   string            `json:"value"`
	Context map[string]string `json:"context"`
}

// RenderVars is the variables used to render the Grafana dashboard
type RenderVars []Var

func (vars RenderVars) GetValues(name string) ([]Val, error) {
	for _, v := range vars {
		if v.Name == name {
			return v.Values, nil
		}
	}
	return nil, errors.New("var not found")
}

// GetGlobalContext returns the global context made by the name-value pairs,
// if the var has multiple values, pick the first one.
func (vars RenderVars) GetGlobalContext() map[string]string {
	ctx := make(map[string]string, 0)
	for _, v := range vars {
		if len(v.Values) > 0 {
			ctx[v.Name] = v.Values[0].Value
		}
	}
	return ctx
}

func mergeContext(ctx map[string]string, overrides map[string]string) map[string]string {
	mergedCtx := make(map[string]string)
	for k, v := range ctx {
		mergedCtx[k] = v
	}

	for k, v := range overrides {
		mergedCtx[k] = v
	}
	return mergedCtx
}

// RenderDashboardWithTemplate renders the grafana dashboard with predefined variables statically.
// It's similar to the normal grafana dashboard rendering but it will support alerts with template variables.
func RenderDashboardWithTemplate(config UpdateConfig, vars RenderVars) error {
	grafcli := sdk.NewClient(config.APIUrl, config.BasicAuth, &http.Client{})
	rawJsonBytes, prop, err := grafcli.GetRawDashboard(config.TemplateSlug)
	if err != nil {
		return err
	}

	rendered, err := RenderDashboard(rawJsonBytes, vars)
	if err != nil {
		return err
	}
	return grafcli.SetRawDashboard(rendered, prop.FolderId)
}

// RenderDashboard will render the Grafana dashboard with variables.
func RenderDashboard(body []byte, vars RenderVars) ([]byte, error) {
	var err error
	if body, err = renderPanels(body, vars); err != nil {
		return nil, err
	}

	// replace the remaining variables
	rawJson := string(body)
	for k, v := range vars.GetGlobalContext() {
		rawJson = replaceVar(rawJson, k, v)
	}
	rawJson = strings.ReplaceAll(rawJson, "**template**", "")
	return []byte(rawJson), nil
}

// renderPanels will populate the repeated panels.
func renderPanels(body []byte, vars RenderVars) ([]byte, error) {
	jsonBody, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}

	panels, err := jsonBody.Get("panels").Array()
	if err != nil {
		return nil, err
	}

	// add end panel for ending
	panels = append(panels, map[string]interface{}{"end": true})

	var (
		newPanels, repeatedPanels []map[string]interface{}
		yOffset                   = 0
	)
	for _, panel := range panels {
		panelMap := panel.(map[string]interface{})
		if panelMap["type"] == "row" || panelMap["end"] == true {
			if len(repeatedPanels) > 0 {
				repeatKey := repeatedPanels[0]["repeat"].(string)
				if vals, err := vars.GetValues(repeatKey); err == nil {
					ctx := vars.GetGlobalContext()
					var baseOffset = panelsHeight(repeatedPanels)
					for _, v := range vals {
						// override the global context with local one
						mergedCtx := mergeContext(ctx, v.Context)
						mergedCtx[repeatKey] = v.Value
						for _, p := range repeatedPanels {
							x := renderMapWithVar(p, mergedCtx, yOffset)
							newPanels = append(newPanels, x)
						}
						yOffset += baseOffset
					}
				} else {
					newPanels = append(newPanels, repeatedPanels...)
				}
			}

			repeatedPanels = []map[string]interface{}{panelMap}
		} else if len(repeatedPanels) > 0 {
			repeatedPanels = append(repeatedPanels, panelMap)
		} else {
			newPanels = append(newPanels, panelMap)
		}
	}

	// update the panel ids
	for i, panel := range newPanels {
		panel["id"] = i + 1
	}

	jsonBody.Set("panels", newPanels)
	return jsonBody.Encode()
}

// panelsHeight calculates the total height of the panels
func panelsHeight(repeatedPanels []map[string]interface{}) int {
	var maxY = 0
	var minY = math.MaxInt32
	for _, p := range repeatedPanels {
		pMap := simplejson.NewFromAny(p)
		y, _ := pMap.GetPath("gridPos", "y").Int()
		h, _ := pMap.GetPath("gridPos", "h").Int()
		if y < minY {
			minY = y
		}
		if y+h > maxY {
			maxY = y + h
		}
	}
	return maxY - minY
}

func renderMapWithVar(m map[string]interface{}, ctx map[string]string, yOffset int) (res map[string]interface{}) {
	mSimple := simplejson.NewFromAny(m)
	y, _ := mSimple.GetPath("gridPos", "y").Int()
	mSimple.SetPath([]string{"gridPos", "y"}, y+yOffset)

	marshalled, _ := mSimple.Encode()
	marshalledStr := string(marshalled)
	for k, v := range ctx {
		marshalledStr = replaceVar(marshalledStr, k, v)
	}

	// recover
	mSimple.SetPath([]string{"gridPos", "y"}, y)

	_ = json.Unmarshal([]byte(marshalledStr), &res)
	return res
}

func replaceVar(template string, key string, val string) string {
	template = strings.ReplaceAll(template, "$"+key, val)
	template = strings.ReplaceAll(template, "${"+key+"}", val)
	return template
}
