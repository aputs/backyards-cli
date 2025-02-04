## backyards install

Install Backyards

### Synopsis

Installs Backyards.

The command automatically applies the resources.
It can only dump the applicable resources with the '--dump-resources' option.

The command can install every component at once with the '--install-everything' option.

```
backyards install [flags]
```

### Examples

```
  # Default install.
  backyards install

  # Install Backyards into a non-default namespace.
  backyards install -n backyards-system
```

### Options

```
      --api-image string         Image for the API
  -d, --dump-resources           Dump resources to stdout instead of applying them
      --enable-auditsink         Enable deploying the auditsink service and sending audit logs over http
      --enable-auth              Enable authentication with impersonation
  -h, --help                     help for install
      --install-canary           Install Canary feature as well
      --install-cert-manager     Install cert-manager as well
      --install-demoapp          Install Demo application as well
  -a, --install-everything       Install every component at once
      --install-istio            Install Istio mesh as well
      --istio-namespace string   Namespace of Istio sidecar injector (default "istio-system")
      --release-name string      Name of the release (default "backyards")
      --run-demo                 Send load to demo application and opens up dashboard
      --web-image string         Image for the frontend
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

