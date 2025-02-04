## backyards cert-manager install

Install cert-manager

### Synopsis

Installs cert-manager.

The command automatically applies the resources.
It can only dump the applicable resources with the '--dump-resources' option.

```
backyards cert-manager install [flags]
```

### Examples

```
  # Install to the cert-manager namespace. This command will fail if cert-manager is already installed from a different source.
  backyards cert-manager install

```

### Options

```
  -d, --dump-resources   Dump resources to stdout instead of applying them
  -h, --help             help for install
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

* [backyards cert-manager](backyards_cert-manager.md)	 - Install and manage cert-manager

