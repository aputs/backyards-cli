## backyards graph

Show graph

### Synopsis

Show graph

```
backyards graph [[--service=]namespace/servicename] [flags]
```

### Options

```
  -h, --help                         help for graph
      --outbound                     Whether to show outbound or inbound metrics
  -r, --refresh-interval duration    the interval to refresh the dashboard (default 10s)
  -d, --relative-duration duration   the relative duration from now to load the graph (default 15m0s)
      --title-suffix string          Title suffix
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

