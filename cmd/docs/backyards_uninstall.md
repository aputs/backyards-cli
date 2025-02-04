## backyards uninstall

Uninstall Backyards

### Synopsis

Uninstall Backyards

The command automatically removes the resources.
It can only dump the removable resources with the '--dump-resources' option.

```
backyards uninstall [flags]
```

### Examples

```
  # Default uninstall
  backyards uninstall

  # Uninstall Backyards from a non-default namespace
  backyards uninstall -n backyards-system
```

### Options

```
  -d, --dump-resources           Dump resources to stdout instead of applying them
  -h, --help                     help for uninstall
      --istio-namespace string   Namespace of Istio sidecar injector (default "istio-system")
      --release-name string      Name of the release (default "backyards")
      --uninstall-canary         Uninstall Canary feature as well
      --uninstall-cert-manager   Uninstall cert-manager as well
      --uninstall-demoapp        Uninstall Demo application as well
  -a, --uninstall-everything     Uninstall every component at once
      --uninstall-istio          Uninstall Istio mesh as well
```

### Options inherited from parent commands

```
  -u, --base-url string     Custom Backyards base URL. Uses automatic port forwarding / proxying if empty
      --cacert string       The CA to use for verifying Backyards' server certificate
      --context string      name of the kubeconfig context to use
      --interactive         ask questions interactively even if stdin or stdout is non-tty
  -c, --kubeconfig string   path to the kubeconfig file to use for CLI requests
  -p, --local-port int      Use this local port for port forwarding / proxying to Backyards (when set to 0, a random port will be used) (default -1)
  -n, --namespace string    namespace in which Backyards is installed [$BACKYARDS_NAMESPACE] (default "backyards-system")
      --non-interactive     never ask questions interactively
  -o, --output string       output format (table|yaml|json) (default "table")
      --use-portforward     Use port forwarding instead of proxying to reach Backyards
  -v, --verbose             turn on debug logging
```

### SEE ALSO

* [backyards](backyards.md)	 - Install and manage Backyards

