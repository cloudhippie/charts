# hcloud-csi-user

![Version: 1.0.2](https://img.shields.io/badge/Version-1.0.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.16.0](https://img.shields.io/badge/AppVersion-2.16.0-informational?style=flat-square)

Hetzner Cloud CSI component for user clusters

**Homepage:** <https://github.com/hetznercloud/csi-driver>

**This chart is not maintained by the upstream project and any issues with the
chart should be raised [here](https://github.com/cloudhippie/charts/issues/new)**

## Installing the Chart

### OCI (Recommended)

```console
helm install hcloud-csi-user oci://ghcr.io/cloudhippie/charts/hcloud-csi-user
```

### Traditional

```console
helm repo add cloudhippie https://cloudhippie.github.io/charts
helm repo update
helm install hcloud-csi-user cloudhippie/hcloud-csi-user
```

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| tboerger | <thomas@webhippie.de> | <https://github.com/tboerger> |

## Source Code

* <https://github.com/cloudhippie/charts>
* <https://github.com/hetznercloud/csi-driver>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{"nodeAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":{"nodeSelectorTerms":[{"matchExpressions":[{"key":"instance.hetzner.cloud/is-root-server","operator":"NotIn","values":["true"]},{"key":"instance.hetzner.cloud/provided-by","operator":"NotIn","values":["robot"]}]}]}}}` | Affinity for the daemonset |
| annotations | object | `{}` | Define additional annotations |
| config.enableTopology | bool | `true` | Enable topology setting |
| extraEnvVariables | list | `[]` | Extra environment variables for driver |
| extraInitContainers | list | `[]` | List of extra init containers |
| extraSidecarContainers | list | `[]` | List of extra sidecar containers |
| extraVolumeMounts | list | `[]` | List of extra volume mounts |
| extraVolumes | list | `[]` | List of extra volumes |
| fullnameOverride | string | `""` | Override the fullname |
| image.driver.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| image.driver.repository | string | `"hetznercloud/hcloud-csi-driver"` | Image repository used by daemonset |
| image.driver.tag | string | `""` | Optional tag for the repository, defaults to app version |
| image.liveness.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| image.liveness.repository | string | `"registry.k8s.io/sig-storage/livenessprobe"` | Image repository used by daemonset |
| image.liveness.tag | string | `"v2.16.0"` | Optional tag for the repository |
| image.pullSecrets | list | `[]` | Optional name of pull secret if using a private registry |
| image.registrar.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| image.registrar.repository | string | `"registry.k8s.io/sig-storage/csi-node-driver-registrar"` | Image repository used by daemonset |
| image.registrar.tag | string | `"v2.14.0"` | Optional tag for the repository |
| labels | object | `{}` | Define additional labels |
| metrics.enabled | bool | `true` | Enable metrics |
| metrics.internalPort | int | `9189` | Internal metrics port of the service |
| metrics.port | int | `9189` | Metrics port of the service |
| metrics.serviceMonitor.annotations | object | `{}` | Additional annotations for the service monitor |
| metrics.serviceMonitor.enabled | bool | `false` | Enable service monitor |
| metrics.serviceMonitor.honorLabels | bool | `false` | HonorLabels chooses the metric’s labels on collisions with target labels |
| metrics.serviceMonitor.interval | string | `nil` | Interval at which metrics should be scraped |
| metrics.serviceMonitor.jobLabel | string | `nil` | Optional job label for the target service in Prometheus |
| metrics.serviceMonitor.labels | object | `{}` | Additional labels for the service monitor |
| metrics.serviceMonitor.metricRelabelings | list | `[]` | List of metric relabel configs to apply to samples before ingestion |
| metrics.serviceMonitor.namespace | string | `nil` | Namespace for ServiceMonitor, defaults to release namespace |
| metrics.serviceMonitor.relabelings | list | `[]` | List of relabel configs to apply to samples before scraping |
| metrics.serviceMonitor.scrapeTimeout | string | `nil` | Timeout after which the scrape is ended |
| nameOverride | string | `""` | Override the name |
| nodeSelector | object | `{}` | Node selector for the daemonset |
| podSecurityContext.driver | object | `{"privileged":true}` | Security context for the driver pod |
| podSecurityContext.liveness | object | `{}` | Security context for the liveness-probe pod |
| podSecurityContext.registrar | object | `{}` | Security context for the registrar pod |
| priorityClassName | string | `"system-cluster-critical"` | Priority class name for daemonset |
| resources.driver | object | `{"limits":{},"requests":{"cpu":"100m","memory":"64Mi"}}` | Resources for the driver pod |
| resources.liveness | object | `{"limits":{},"requests":{}}` | Resources for the liveness-probe pod |
| resources.registrar | object | `{"limits":{},"requests":{}}` | Resources for the registrar pod |
| securityContext | object | `{"fsGroup":1001}` | Security context for the daemonset |
| service.annotations | object | `{}` | Additional annotations for the service |
| service.labels | object | `{}` | Additional labels for the service |
| service.type | string | `"ClusterIP"` | Type of the service |
| serviceAccount.annotations | object | `{}` | Define annotations for the service account |
| serviceAccount.create | bool | `true` | Create a new service account |
| serviceAccount.name | string | `""` | Optional name for an existing service account |
| tolerations | list | `[{"effect":"NoExecute","operator":"Exists"},{"effect":"NoSchedule","operator":"Exists"},{"key":"CriticalAddonsOnly","operator":"Exists"}]` | Tolerations for the daemonset |
| updateStrategy | object | `{"type":"RollingUpdate"}` | Updaqte strategy for daemonset |
