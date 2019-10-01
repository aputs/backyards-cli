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
	url2 "net/url"
	"os"
	"os/signal"

	"github.com/banzaicloud/backyards-cli/pkg/cli"
	"github.com/pkg/browser"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type dashboardCommand struct{}

type DashboardOptions struct {
	QueryParams map[string]string
}

func NewDashboardOptions() *DashboardOptions {
	return &DashboardOptions{
		QueryParams: make(map[string]string),
	}
}

func NewDashboardCommand(cli cli.CLI, options *DashboardOptions) *cobra.Command {
	c := dashboardCommand{}

	cmd := &cobra.Command{
		Use:   "dashboard [flags]",
		Short: "Open the Backyards dashboard in a web browser",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			err = c.run(cli, options)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}

func (c *dashboardCommand) run(cli cli.CLI, options *DashboardOptions) error {
	var err error
	url, err := cli.GetEndpointURL("")
	if err != nil {
		return err
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	defer signal.Stop(signals)
	defer func() {
		<-signals
	}()

	url, err = withQueryParams(url, options.QueryParams)
	if err != nil {
		return err
	}

	log.Infof("Backyards UI is available at %s", url)
	err = browser.OpenURL(url)
	if err != nil {
		return err
	}

	return nil
}

func withQueryParams(url string, params map[string]string) (string, error) {
	uri, err := url2.ParseRequestURI(url)
	if err != nil {
		return "", err
	}

	q := uri.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	uri.RawQuery = q.Encode()

	return uri.String(), nil
}
