---
title: Vault-Operator Version
menu:
  product_vault-operator_0.1.0:
    identifier: vault-operator-version
    name: Vault-Operator Version
    parent: reference
product_name: vault-operator
left_menu: product_vault-operator_0.1.0
section_menu_id: reference
---
## vault-operator version

Prints binary version number.

### Synopsis

Prints binary version number.

```
vault-operator version [flags]
```

### Options

```
  -h, --help    help for version
      --short   Print just the version number.
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --bypass-validating-webhook-xray   if true, bypasses validating webhook xray checks
      --enable-analytics                 Send analytical events to Google Analytics (default true)
      --log-flush-frequency duration     Maximum number of seconds between log flushes (default 5s)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files (default true)
      --stderrthreshold severity         logs at or above this threshold go to stderr
      --use-kubeapiserver-fqdn-for-aks   if true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [vault-operator](/docs/reference/vault-operator.md)	 - Vault Operator by AppsCode - HashiCorp Vault Operator for Kubernetes

