## backyards canary install

Install Canary feature

### Synopsis

Installs Canary feature.

The command automatically applies the resources.
It can only dump the applicable resources with the '--dump-resources' option.


```
backyards canary install [flags]
```

### Examples

```
  # Default install.
  backyards canary install

  # Install canary into a non-default namespace.
  backyards canary install --canary-namespace backyards-canary
```

### Options

```
      --canary-namespace string   Namespace for the canary operator (default "backyards-canary")
  -d, --dump-resources            Dump resources to stdout instead of applying them
  -h, --help                      help for install
      --istio-namespace string    Namespace of Istio sidecar injector (default "istio-system")
      --prometheus-url string     Prometheus URL for metrics (default "http://backyards-prometheus.backyards-system:9090/prometheus")
      --release-name string       Name of the release (default "canary-operator")
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

* [backyards canary](backyards_canary.md)	 - Install and manage Canary feature

