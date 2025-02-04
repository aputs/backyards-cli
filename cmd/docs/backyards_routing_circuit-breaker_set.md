## backyards routing circuit-breaker set

Set traffic shifting rules for a service

### Synopsis

Set traffic shifting rules for a service

```
backyards routing circuit-breaker set [[--service=]namespace/servicename] [[--version=]subset=weight] ... [flags]
```

### Options

```
      --baseEjectionTime duration           Minimum ejection duration. A host will remain ejected for a period equal to the product of minimum ejection duration and the number of times the host has been ejected (default 30s)
      --connect-timeout duration            TCP connection timeout (default 3s)
      --consecutiveErrors int32             Number of errors before a host is ejected from the connection pool (default 5)
  -h, --help                                help for set
      --interval duration                   Time interval between ejection sweep analysis (default 10s)
      --max-connections int32               Maximum number of HTTP1/TCP connections to a destination host (default 1024)
      --max-pending-requests int32          Maximum number of pending HTTP requests to a destination (default 1024)
      --max-requests int32                  Maximum number of requests to a backend (default 1024)
      --max-requests-per-connection int32   Maximum number of requests per connection to a backend. Setting this parameter to 1 disables keep alive (default 1)
      --max-retries int32                   Maximum number of retries that can be outstanding to all hosts in an envoy cluster at a given time (default 1024)
      --maxEjectionPercent int32            Maximum % of hosts in the load balancing pool for the upstream service that can be ejected (default 100)
      --service string                      Service name
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

* [backyards routing circuit-breaker](backyards_routing_circuit-breaker.md)	 - Manage circuit-breaker configurations

