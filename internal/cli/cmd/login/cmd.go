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

package login

import (
	"fmt"

	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/routing/common"
	"github.com/banzaicloud/backyards-cli/pkg/cli"
	"github.com/spf13/cobra"
)

func NewLoginCmd(cli cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Aliases: []string{"l"},
		Short:   "Log in to Backyards",
		RunE: func(cmd *cobra.Command, args []string) error {
			authClient, err := common.GetAuthClient(cli)
			if err != nil {
				return err
			}
			body, err := authClient.Login()
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", body)
			return nil
		},
	}

	cmd.PersistentFlags().String("base-url", "", "Custom Backyards base URL. Use port-forwarding if empty.")

	return cmd
}
