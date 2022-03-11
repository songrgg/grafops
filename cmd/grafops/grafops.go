package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
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
	DashboardUID string `json:"dashboardUID"`
	BasicAuth    string `json:"basicAuth"`
	ConfigPath   string `json:"configPath"`
}

func (o *options) validate() {
	if o.Host == "" {
		fmt.Println("hosts can't be empty")
		os.Exit(-1)
	}

	if o.DashboardUID == "" {
		fmt.Println("template UID can't be empty")
		os.Exit(-1)
	}

	if o.ConfigPath == "" {
		fmt.Println("config path can't be empty")
		os.Exit(-1)
	}
}

// NewGrafOpsCommand creates `grafops` command.
func NewGrafOpsCommand() *cobra.Command {
	options := options{}
	cmds := &cobra.Command{
		Use:   "grafops",
		Short: "grafops manages the Grafana dashboards",
		Run: func(cmd *cobra.Command, args []string) {
			options.validate()

			configBytes, err := ioutil.ReadFile(options.ConfigPath)
			if err != nil {
				log.Fatalf("Configuration file doesn't exist")
			}

			viper.SetConfigType("yaml")
			err = viper.ReadConfig(bytes.NewBuffer(configBytes))
			if err != nil {
				log.Fatalf("Failed to read configuration file: %v", err)
			}

			var vars grafana.RenderVars
			err = viper.UnmarshalKey("vars", &vars)
			if err != nil {
				fmt.Println("fail to load config file: ", err)
				os.Exit(-1)
			}

			err = grafana.RenderDashboardWithTemplate(grafana.UpdateConfig{
				APIUrl:       options.Host,
				DashboardUID: options.DashboardUID,
				BasicAuth:    options.BasicAuth,
			}, vars)
			if err != nil {
				log.Fatalf("fail to render the Grafana dashboard: %v", err)
			}
			log.Println("Render the dashboard successfully")
		},
	}

	cmds.PersistentFlags().StringVar(&options.Host, "host", "", "The host name for Grafana server")
	cmds.PersistentFlags().StringVarP(&options.DashboardUID, "dashboard_uid", "u", "",
		"The UID of the template dashboard in Grafana, it could be in the URL of Grafana dashboard, for example,"+
			"http://localhost:3000/d/RKAQZi9Zk/service-monitoring, the UID is `RKAQZi9Zk`")
	cmds.PersistentFlags().StringVar(&options.BasicAuth, "basic_auth", "",
		"Basic auth for the Grafana API")
	cmds.PersistentFlags().StringVarP(&options.ConfigPath, "config_path", "c", "",
		"Yaml configuration file path")

	return cmds
}
