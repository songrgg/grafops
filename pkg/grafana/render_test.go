package grafana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const body = `
{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "gnetId": null,
  "graphTooltip": 0,
  "id": 2,
  "iteration": 1584819938476,
  "links": [],
  "panels": [
    {
      "collapsed": false,
      "datasource": null,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 0
      },
      "id": 2,
      "panels": [],
      "repeat": "SERVICE_NAME",
      "scopedVars": {
        "SERVICE_NAME": {
          "selected": true,
          "text": "news",
          "value": "news"
        }
      },
      "title": "$SERVICE_NAME",
      "type": "row"
    },
    {
      "content": "This is $SERVICE_NAME dashboard.\n\nHere's the variable TEST_VAR= \"$TEST_VAR\".",
      "datasource": "myinfluxdb",
      "description": "This is $SERVICE_NAME dashboard.",
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 1
      },
      "id": 4,
      "mode": "markdown",
      "options": {},
      "scopedVars": {
        "SERVICE_NAME": {
          "selected": true,
          "text": "news",
          "value": "news"
        }
      },
      "targets": [
        {
          "groupBy": [
            {
              "params": [
                "$__interval"
              ],
              "type": "time"
            },
            {
              "params": [
                "null"
              ],
              "type": "fill"
            }
          ],
          "orderByTime": "ASC",
          "policy": "default",
          "query": "SELECT '$SERVICE_NAME', '$TEST_VAR' FROM \"http_req_duration\" WHERE $timeFilter GROUP BY time($__interval) fill(null)",
          "rawQuery": true,
          "refId": "A",
          "resultFormat": "time_series",
          "select": [
            [
              {
                "params": [
                  "value"
                ],
                "type": "field"
              },
              {
                "params": [],
                "type": "mean"
              }
            ]
          ],
          "tags": []
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "$SERVICE_NAME",
      "type": "text"
    },
    {
      "content": "Nothing on this.\n\n\n\n",
      "datasource": "myinfluxdb",
      "description": "This is $SERVICE_NAME dashboard.",
      "gridPos": {
        "h": 8,
        "w": 17,
        "x": 3,
        "y": 9
      },
      "id": 5,
      "mode": "markdown",
      "options": {},
      "scopedVars": {
        "SERVICE_NAME": {
          "selected": true,
          "text": "news",
          "value": "news"
        }
      },
      "targets": [
        {
          "groupBy": [
            {
              "params": [
                "$__interval"
              ],
              "type": "time"
            },
            {
              "params": [
                "null"
              ],
              "type": "fill"
            }
          ],
          "orderByTime": "ASC",
          "policy": "default",
          "query": "SELECT '$SERVICE_NAME' FROM \"http_req_duration\" WHERE $timeFilter GROUP BY time($__interval) fill(null)",
          "rawQuery": true,
          "refId": "A",
          "resultFormat": "time_series",
          "select": [
            [
              {
                "params": [
                  "value"
                ],
                "type": "field"
              },
              {
                "params": [],
                "type": "mean"
              }
            ]
          ],
          "tags": []
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "$SERVICE_NAME ----- 2",
      "type": "text"
    }
  ],
  "schemaVersion": 22,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": [
      {
        "allValue": null,
        "current": {
          "selected": false,
          "text": "news",
          "value": "news"
        },
        "hide": 0,
        "includeAll": false,
        "label": "Service Name",
        "multi": false,
        "name": "SERVICE_NAME",
        "options": [
          {
            "selected": true,
            "text": "news",
            "value": "news"
          },
          {
            "selected": false,
            "text": "payment",
            "value": "payment"
          },
          {
            "selected": false,
            "text": "user",
            "value": "user"
          }
        ],
        "query": "news,payment,user",
        "skipUrlSync": false,
        "type": "custom"
      },
      {
        "allValue": null,
        "current": {
          "selected": false,
          "text": "global",
          "value": "global"
        },
        "hide": 0,
        "includeAll": false,
        "label": null,
        "multi": false,
        "name": "TEST_VAR",
        "options": [
          {
            "selected": true,
            "text": "global",
            "value": "global"
          }
        ],
        "query": "global",
        "skipUrlSync": false,
        "type": "custom"
      }
    ]
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {
    "refresh_intervals": [
      "5s",
      "10s",
      "30s",
      "1m",
      "5m",
      "15m",
      "30m",
      "1h",
      "2h",
      "1d"
    ]
  },
  "timezone": "",
  "title": "**template** service monitoring",
  "uid": "VoUygmrWz",
  "version": 10
}
`

const ExpectedRendered = `{"annotations":{"list":[{"builtIn":1,"datasource":"-- Grafana --","enable":true,"hide":true,"iconColor":"rgba(0, 211, 255, 1)","name":"Annotations \u0026 Alerts","type":"dashboard"}]},"editable":true,"gnetId":null,"graphTooltip":0,"id":2,"iteration":1584819938476,"links":[],"panels":[{"collapsed":false,"datasource":null,"gridPos":{"h":1,"w":24,"x":0,"y":0},"id":1,"panels":[],"repeat":"SERVICE_NAME","scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"title":"news","type":"row"},{"content":"This is news dashboard.\n\nHere's the variable TEST_VAR= \"local_news\".","datasource":"myinfluxdb","description":"This is news dashboard.","gridPos":{"h":8,"w":12,"x":0,"y":1},"id":2,"mode":"markdown","options":{},"scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"targets":[{"groupBy":[{"params":["$__interval"],"type":"time"},{"params":["null"],"type":"fill"}],"orderByTime":"ASC","policy":"default","query":"SELECT 'news', 'local_news' FROM \"http_req_duration\" WHERE $timeFilter GROUP BY time($__interval) fill(null)","rawQuery":true,"refId":"A","resultFormat":"time_series","select":[[{"params":["value"],"type":"field"},{"params":[],"type":"mean"}]],"tags":[]}],"timeFrom":null,"timeShift":null,"title":"news","type":"text"},{"content":"Nothing on this.\n\n\n\n","datasource":"myinfluxdb","description":"This is news dashboard.","gridPos":{"h":8,"w":17,"x":3,"y":9},"id":3,"mode":"markdown","options":{},"scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"targets":[{"groupBy":[{"params":["$__interval"],"type":"time"},{"params":["null"],"type":"fill"}],"orderByTime":"ASC","policy":"default","query":"SELECT 'news' FROM \"http_req_duration\" WHERE $timeFilter GROUP BY time($__interval) fill(null)","rawQuery":true,"refId":"A","resultFormat":"time_series","select":[[{"params":["value"],"type":"field"},{"params":[],"type":"mean"}]],"tags":[]}],"timeFrom":null,"timeShift":null,"title":"news ----- 2","type":"text"},{"collapsed":false,"datasource":null,"gridPos":{"h":1,"w":24,"x":0,"y":17},"id":4,"panels":[],"repeat":"SERVICE_NAME","scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"title":"payment","type":"row"},{"content":"This is payment dashboard.\n\nHere's the variable TEST_VAR= \"local_payment\".","datasource":"myinfluxdb","description":"This is payment dashboard.","gridPos":{"h":8,"w":12,"x":0,"y":18},"id":5,"mode":"markdown","options":{},"scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"targets":[{"groupBy":[{"params":["$__interval"],"type":"time"},{"params":["null"],"type":"fill"}],"orderByTime":"ASC","policy":"default","query":"SELECT 'payment', 'local_payment' FROM \"http_req_duration\" WHERE $timeFilter GROUP BY time($__interval) fill(null)","rawQuery":true,"refId":"A","resultFormat":"time_series","select":[[{"params":["value"],"type":"field"},{"params":[],"type":"mean"}]],"tags":[]}],"timeFrom":null,"timeShift":null,"title":"payment","type":"text"},{"content":"Nothing on this.\n\n\n\n","datasource":"myinfluxdb","description":"This is payment dashboard.","gridPos":{"h":8,"w":17,"x":3,"y":26},"id":6,"mode":"markdown","options":{},"scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"targets":[{"groupBy":[{"params":["$__interval"],"type":"time"},{"params":["null"],"type":"fill"}],"orderByTime":"ASC","policy":"default","query":"SELECT 'payment' FROM \"http_req_duration\" WHERE $timeFilter GROUP BY time($__interval) fill(null)","rawQuery":true,"refId":"A","resultFormat":"time_series","select":[[{"params":["value"],"type":"field"},{"params":[],"type":"mean"}]],"tags":[]}],"timeFrom":null,"timeShift":null,"title":"payment ----- 2","type":"text"},{"collapsed":false,"datasource":null,"gridPos":{"h":1,"w":24,"x":0,"y":34},"id":7,"panels":[],"repeat":"SERVICE_NAME","scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"title":"user","type":"row"},{"content":"This is user dashboard.\n\nHere's the variable TEST_VAR= \"global\".","datasource":"myinfluxdb","description":"This is user dashboard.","gridPos":{"h":8,"w":12,"x":0,"y":35},"id":8,"mode":"markdown","options":{},"scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"targets":[{"groupBy":[{"params":["$__interval"],"type":"time"},{"params":["null"],"type":"fill"}],"orderByTime":"ASC","policy":"default","query":"SELECT 'user', 'global' FROM \"http_req_duration\" WHERE $timeFilter GROUP BY time($__interval) fill(null)","rawQuery":true,"refId":"A","resultFormat":"time_series","select":[[{"params":["value"],"type":"field"},{"params":[],"type":"mean"}]],"tags":[]}],"timeFrom":null,"timeShift":null,"title":"user","type":"text"},{"content":"Nothing on this.\n\n\n\n","datasource":"myinfluxdb","description":"This is user dashboard.","gridPos":{"h":8,"w":17,"x":3,"y":43},"id":9,"mode":"markdown","options":{},"scopedVars":{"SERVICE_NAME":{"selected":true,"text":"news","value":"news"}},"targets":[{"groupBy":[{"params":["$__interval"],"type":"time"},{"params":["null"],"type":"fill"}],"orderByTime":"ASC","policy":"default","query":"SELECT 'user' FROM \"http_req_duration\" WHERE $timeFilter GROUP BY time($__interval) fill(null)","rawQuery":true,"refId":"A","resultFormat":"time_series","select":[[{"params":["value"],"type":"field"},{"params":[],"type":"mean"}]],"tags":[]}],"timeFrom":null,"timeShift":null,"title":"user ----- 2","type":"text"}],"schemaVersion":22,"style":"dark","tags":[],"templating":{"list":[{"allValue":null,"current":{"selected":false,"text":"news","value":"news"},"hide":0,"includeAll":false,"label":"Service Name","multi":false,"name":"SERVICE_NAME","options":[{"selected":true,"text":"news","value":"news"},{"selected":false,"text":"payment","value":"payment"},{"selected":false,"text":"user","value":"user"}],"query":"news,payment,user","skipUrlSync":false,"type":"custom"},{"allValue":null,"current":{"selected":false,"text":"global","value":"global"},"hide":0,"includeAll":false,"label":null,"multi":false,"name":"TEST_VAR","options":[{"selected":true,"text":"global","value":"global"}],"query":"global","skipUrlSync":false,"type":"custom"}]},"time":{"from":"now-6h","to":"now"},"timepicker":{"refresh_intervals":["5s","10s","30s","1m","5m","15m","30m","1h","2h","1d"]},"timezone":"","title":" service monitoring","uid":"VoUygmrWz","version":10}`

func TestUpdateDashboardWithTemplate(t *testing.T) {
	err := RenderDashboardWithTemplate(UpdateConfig{
		APIUrl:       "http://localhost:3000",
		DashboardUID: "RKAQZi9Zk",
		BasicAuth:    "",
	}, []Var{
		{
			Name: "SERVICE_NAME",
			Values: []Val{
				{
					Value: "news",
					Context: map[string]string{
						"TEST_VAR": "local_news",
					},
				},
				{
					Value: "payment",
					Context: map[string]string{
						"TEST_VAR": "local_payment",
					},
				},
				{
					Value:   "user",
					Context: map[string]string{},
				},
			},
		},
		{
			Name: "TEST_VAR",
			Values: []Val{
				{
					Value:   "global",
					Context: map[string]string{},
				},
			},
		},
	})

	assert.Nil(t, err)
}

func TestRenderDashboard(t *testing.T) {
	rendered, err := RenderDashboard([]byte(body), []Var{
		{
			Name: "SERVICE_NAME",
			Values: []Val{
				{
					Value: "news",
					Context: map[string]string{
						"TEST_VAR": "local_news",
					},
				},
				{
					Value: "payment",
					Context: map[string]string{
						"TEST_VAR": "local_payment",
					},
				},
				{
					Value:   "user",
					Context: map[string]string{},
				},
			},
		},
		{
			Name: "TEST_VAR",
			Values: []Val{
				{
					Value:   "global",
					Context: map[string]string{},
				},
			},
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, []byte(ExpectedRendered), rendered)
}
