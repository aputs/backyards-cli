## backyards istio uninstall

Output or delete Kubernetes resources to uninstall Istio

### Synopsis

Output or delete Kubernetes resources to uninstall Istio.

The command automatically removes the resources.
It can only dump the removable resources with the '--dump-resources' option.

```
backyards istio uninstall [flags]
```

### Examples

```
  # Default uninstall.
  backyards istio uninstall

  # Uninstall Istio from a non-default namespace.
  backyards istio uninstall -n custom-istio-ns
```

### Options

```
  -d, --dump-resources        Dump resources to stdout instead of applying them
  -h, --help                  help for uninstall
      --release-name string   Name of the release (default "istio-operator")
```

### Options inherited from parent commands

```
  -u, --base-url string     Custom Backyards base URL. Uses automatic port forwarding / proxying if empty
      --cacert string       The CA to use for verifying Backyards' server certificate
      --context string      name of the kubeconfig context to use
      --interactive         ask questions interactively even if stdin or stdout is non-tty
  -c, --kubeconfig string   path to the kubeconfig file to use for CLI requests
  -p, --local-port int      Use this local port for port forwarding / proxying to Backyards (when set to 0, a random port will be used) (default -1)
  -n, --namespace string    Namespace in which Istio is installed [$ISTIO_NAMESPACE] (default "istio-system")
      --non-interactive     never ask questions interactively
  -o, --output string       output format (table|yaml|json) (default "table")
      --use-portforward     Use port forwarding instead of proxying to reach Backyards
  -v, --verbose             turn on debug logging
```

### SEE ALSO

* [backyards istio](backyards_istio.md)	 - Install and manage Istio

