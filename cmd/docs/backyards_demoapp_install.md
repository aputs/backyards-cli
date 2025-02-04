## backyards demoapp install

Install demo application

### Synopsis

Installs demo application.

The command automatically applies the resources.
It can only dump the applicable resources with the '--dump-resources' option.

```
backyards demoapp install [flags]
```

### Examples

```
  # Default install.
  backyards demoapp install

  # Install Backyards into a non-default namespace.
  backyards demoapp install -n backyards-system
```

### Options

```
  -d, --dump-resources           Dump resources to stdout instead of applying them
  -h, --help                     help for install
      --istio-namespace string   Namespace of Istio sidecar injector (default "istio-system")
```

### Options inherited from parent commands

```
  -u, --base-url string         Custom Backyards base URL. Uses automatic port forwarding / proxying if empty
      --cacert string           The CA to use for verifying Backyards' server certificate
      --context string          name of the kubeconfig context to use
      --demo-namespace string   Namespace for demo application (default "backyards-demo")
      --interactive             ask questions interactively even if stdin or stdout is non-tty
  -c, --kubeconfig string       path to the kubeconfig file to use for CLI requests
  -p, --local-port int          Use this local port for port forwarding / proxying to Backyards (when set to 0, a random port will be used) (default -1)
  -n, --namespace string        namespace in which Backyards is installed [$BACKYARDS_NAMESPACE] (default "backyards-system")
      --non-interactive         never ask questions interactively
  -o, --output string           output format (table|yaml|json) (default "table")
      --use-portforward         Use port forwarding instead of proxying to reach Backyards
  -v, --verbose                 turn on debug logging
```

### SEE ALSO

* [backyards demoapp](backyards_demoapp.md)	 - Install and manage demo application

