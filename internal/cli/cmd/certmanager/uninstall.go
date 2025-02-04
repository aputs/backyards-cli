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

package certmanager

import (
	"context"
	"fmt"
	"time"

	"emperror.dev/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"istio.io/operator/pkg/object"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"

	internalk8s "github.com/banzaicloud/backyards-cli/internal/k8s"

	"github.com/banzaicloud/backyards-cli/pkg/cli"
	"github.com/banzaicloud/backyards-cli/pkg/helm"
	"github.com/banzaicloud/backyards-cli/pkg/k8s"
)

type uninstallCommand struct {
	cli cli.CLI
}

type UninstallOptions struct {
	DumpResources bool
	Skip          bool
	Force         bool
}

func NewUninstallOptions() *UninstallOptions {
	return &UninstallOptions{}
}

func NewUninstallCommand(cli cli.CLI, options *UninstallOptions) *cobra.Command {
	c := &uninstallCommand{
		cli: cli,
	}

	cmd := &cobra.Command{
		Use:   "uninstall [flags]",
		Args:  cobra.NoArgs,
		Short: "Output or delete Kubernetes resources to uninstall cert-manager",
		Long: `Output or delete Kubernetes resources to uninstall cert-manager.

The command automatically removes the resources.
It can only dump the removable resources with the '--dump-resources' option.`,
		Example: `  # Default uninstall.
  backyards cert-manager uninstall

  # Uninstall cert-manager from a non-default namespace.
  backyards cert-manager uninstall --cert-manager-namespace backyards-system`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceErrors = true
			cmd.SilenceUsage = true

			return c.run(cli, options)
		},
	}

	cmd.Flags().BoolVarP(&options.DumpResources, "dump-resources", "d", options.DumpResources, "Dump resources to stdout instead of applying them")
	cmd.Flags().BoolVar(&options.Skip, "skip", false, "Skip uninstalling cert-manager")
	cmd.Flags().BoolVar(&options.Force, "force", false, "Force uninstalling cert-manager")

	return cmd
}

func (c *uninstallCommand) run(cli cli.CLI, options *UninstallOptions) error {
	if options.Skip {
		logrus.Info("Skip uninstalling cert-manager")
		return nil
	}

	err := c.validate(CertManagerNamespace, options)
	if err != nil {
		return err
	}

	objects, err := getCertManagerObjects(CertManagerNamespace)
	if err != nil {
		return err
	}
	objects.Sort(helm.UninstallObjectOrder())

	if !options.DumpResources {
		err := c.deleteResources(objects)
		if err != nil {
			return errors.WrapIf(err, "could not delete k8s resources")
		}
		return nil
	}

	yaml, err := objects.YAMLManifest()
	if err != nil {
		return errors.WrapIf(err, "could not render YAML manifest")
	}
	fmt.Fprint(cli.Out(), yaml)

	return nil
}

func (c *uninstallCommand) validate(namespace string, opts *UninstallOptions) error {
	var err error
	client, err := c.cli.GetK8sClient()
	if err != nil {
		return err
	}
	targetNamespace := &corev1.Namespace{}
	err = client.Get(context.Background(), types.NamespacedName{Name: namespace}, targetNamespace)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return errors.WrapIf(err, "failed to get target namespace for cert-manager")
	}
	if _, ok := targetNamespace.Labels[internalk8s.CLIVersionLabel]; ok {
		return nil
	}
	if opts.Force {
		logrus.Warn("Force uninstalling cert-manager")
		return nil
	}
	return errors.Errorf("cert-manager already installed but not managed by backyards, " +
		"please skip or force uninstalling cert-manager using the cli flags")
}

func (c *uninstallCommand) deleteResources(objects object.K8sObjects) error {
	client, err := c.cli.GetK8sClient()
	if err != nil {
		return err
	}

	err = k8s.DeleteResources(client, c.cli.LabelManager(), objects, k8s.WaitForResourceConditions(wait.Backoff{
		Duration: time.Second * 5,
		Factor:   1,
		Jitter:   0,
		Steps:    24,
	}, k8s.NonExistsConditionCheck))
	if err != nil {
		return errors.WrapIf(err, "could not delete k8s resources")
	}

	return nil
}
