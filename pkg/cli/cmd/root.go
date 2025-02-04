// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"regexp"

	"emperror.dev/errors"
	logrushandler "emperror.dev/handler/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/login"

	"github.com/banzaicloud/backyards-cli/internal/cli/cmd"
	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/canary"
	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/certmanager"
	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/demoapp"
	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/graph"
	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/istio"
	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/routing"
	"github.com/banzaicloud/backyards-cli/pkg/cli"
)

const (
	defaultNamespace = "backyards-system"
)

var (
	backyardsNamespace string
	kubeconfigPath     string
	kubeContext        string
	verbose            bool
	outputFormat       string
	baseURL            string
	localPort          = -1
	usePortforward     bool
	ca                 string

	namespaceRegex = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
)

// RootCmd represents the root Cobra command
var RootCmd = &cobra.Command{
	Use:           "backyards",
	Short:         "Install and manage Backyards",
	SilenceErrors: true,
	SilenceUsage:  false,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		if verbose {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}

		if viper.GetInt("port-forward") < 0 {
			return errors.NewWithDetails(
				"port must be greater than or equal to zero",
				"port", viper.GetInt("port-forward"),
			)
		}

		namespaceFromEnv := os.Getenv("BACKYARDS_NAMESPACE")
		if backyardsNamespace == defaultNamespace && namespaceFromEnv != "" {
			backyardsNamespace = namespaceFromEnv
		}

		if !namespaceRegex.MatchString(backyardsNamespace) {
			return errors.NewWithDetails("invalid namespace", "namespace", backyardsNamespace)
		}

		return nil
	},
}

// Init is a temporary function to set initial values in the root cmd
func Init(version string, commitHash string, buildDate string) {
	RootCmd.Version = version

	RootCmd.SetVersionTemplate(fmt.Sprintf(
		"Backyards CLI version %s (%s) built on %s\n",
		version,
		commitHash,
		buildDate,
	))
}

// GetRootCommand returns the cli root command
func GetRootCommand() *cobra.Command {
	return RootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately
// This is called by main.main(). It only needs to happen once to the RootCmd
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		handler := logrushandler.New(log.New())
		handler.Handle(err)
		os.Exit(1)
	}
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&backyardsNamespace, "namespace", "n", defaultNamespace, "namespace in which Backyards is installed [$BACKYARDS_NAMESPACE]")
	_ = viper.BindPFlag("backyards.namespace", flags.Lookup("namespace"))
	flags.StringVarP(&kubeconfigPath, "kubeconfig", "c", "", "path to the kubeconfig file to use for CLI requests")
	_ = viper.BindPFlag("kubeconfig", flags.Lookup("kubeconfig"))
	flags.StringVar(&kubeContext, "context", "", "name of the kubeconfig context to use")
	_ = viper.BindPFlag("kubecontext", flags.Lookup("context"))
	flags.BoolVarP(&verbose, "verbose", "v", false, "turn on debug logging")

	flags.StringVarP(&outputFormat, "output", "o", "table", "output format (table|yaml|json)")
	_ = viper.BindPFlag("output.format", flags.Lookup("output"))

	_ = viper.BindPFlag("formatting.force-color", flags.Lookup("color"))
	flags.Bool("non-interactive", false, "never ask questions interactively")
	_ = viper.BindPFlag("formatting.non-interactive", flags.Lookup("non-interactive"))
	flags.Bool("interactive", false, "ask questions interactively even if stdin or stdout is non-tty")
	_ = viper.BindPFlag("formatting.force-interactive", flags.Lookup("interactive"))

	flags.StringVarP(&baseURL, "base-url", "u", baseURL, "Custom Backyards base URL. Uses automatic port forwarding / proxying if empty")
	_ = viper.BindPFlag("backyards.url", flags.Lookup("base-url"))
	flags.StringVarP(&ca, "cacert", ca, "", "The CA to use for verifying Backyards' server certificate")
	_ = viper.BindPFlag("backyards.cacert", flags.Lookup("cacert"))
	flags.IntVarP(&localPort, "local-port", "p", localPort, "Use this local port for port forwarding / proxying to Backyards (when set to 0, a random port will be used)")
	_ = viper.BindPFlag("backyards.localPort", flags.Lookup("local-port"))
	flags.BoolVar(&usePortforward, "use-portforward", usePortforward, "Use port forwarding instead of proxying to reach Backyards")
	_ = viper.BindPFlag("backyards.usePortForward", flags.Lookup("use-portforward"))

	cli := cli.NewCli(os.Stdout, RootCmd)

	RootCmd.AddCommand(cmd.NewVersionCommand(cli))
	RootCmd.AddCommand(cmd.NewInstallCommand(cli))
	RootCmd.AddCommand(cmd.NewUninstallCommand(cli))
	RootCmd.AddCommand(cmd.NewDashboardCommand(cli, cmd.NewDashboardOptions()))
	RootCmd.AddCommand(istio.NewRootCmd(cli))
	RootCmd.AddCommand(canary.NewRootCmd(cli))
	RootCmd.AddCommand(demoapp.NewRootCmd(cli))
	RootCmd.AddCommand(routing.NewRootCmd(cli))
	RootCmd.AddCommand(certmanager.NewRootCmd(cli))
	RootCmd.AddCommand(graph.NewGraphCmd(cli, "base.json"))
	RootCmd.AddCommand(login.NewLoginCmd(cli))

	RootCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		return cli.Stop()
	}
}
