package main

import (
	"fmt"
	"os"

	"github.com/songrgg/grafops/pkg/grafana"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	rootCmd := NewGrafOpsCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("fail to run the grafops: ", err)
	}
}

type options struct {
	Host         string `json:"host"`
	TemplateSlug string `json:"templateSlug"`
	BasicAuth    string `json:"basicAuth"`
}

func (o *options) validate() {
	if o.Host == "" {
		fmt.Println("hosts can't be empty")
		os.Exit(-1)
	}

	if o.TemplateSlug == "" {
		fmt.Println("template slug can't be empty")
		os.Exit(-1)
	}
}

// NewGrafOpsCommand creates `grafops` command.
func NewGrafOpsCommand() *cobra.Command {
	options := options{}
	cmds := &cobra.Command{
		Use:   "llsh",
		Short: "llsh executes the remote shell command on multiple remote servers",
		Run: func(cmd *cobra.Command, args []string) {
			options.validate()

			viper.AddConfigPath(".")
			err := viper.ReadInConfig()
			if err != nil {
				fmt.Printf("Fatal error config file: %s \n", err)
				os.Exit(-1)
			}

			var vars grafana.RenderVars
			err = viper.UnmarshalKey("vars", &vars)
			if err != nil {
				fmt.Println("fail to load config file: ", err)
				os.Exit(-1)
			}

			err = grafana.RenderDashboardWithTemplate(grafana.UpdateConfig{
				APIUrl:       options.Host,
				TemplateSlug: options.TemplateSlug,
				BasicAuth:    options.BasicAuth,
			}, vars)
			if err != nil {
				fmt.Println("fail to render the Grafana dashboard: ", err)
				os.Exit(-1)
			}
			fmt.Println("Render the dashboard successfully")
		},
	}

	cmds.PersistentFlags().StringVar(&options.Host, "host", "", "The host name for Grafana server")
	cmds.PersistentFlags().StringVarP(&options.TemplateSlug, "template_slug", "t", "",
		"The slug of the template dashboard in Grafana, it could be in the URL of Grafana dashboard, for example,"+
			"http://localhost:3000/d/RKAQZi9Zk/service-monitoring slug is service-monitoring")
	cmds.PersistentFlags().StringVar(&options.BasicAuth, "basic_auth", "",
		"Basic auth for the Grafana API")

	return cmds
}
