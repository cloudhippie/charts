# hcloud-ccm-mgmt

![Version: 1.3.1](https://img.shields.io/badge/Version-1.3.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.29.1](https://img.shields.io/badge/AppVersion-1.29.1-informational?style=flat-square)

Hetzner Cloud CCM component for management clusters

**Homepage:** <https://github.com/hetznercloud/hcloud-cloud-controller-manager>

**This chart is not maintained by the upstream project and any issues with the
chart should be raised [here](https://github.com/cloudhippie/charts/issues/new)**

## Installing the Chart

### OCI (Recommended)

```console
helm install hcloud-ccm-mgmt oci://ghcr.io/cloudhippie/charts/hcloud-ccm-mgmt
```

### Traditional

```console
helm repo add cloudhippie https://cloudhippie.github.io/charts
helm repo update
helm install hcloud-ccm-mgmt cloudhippie/hcloud-ccm-mgmt
```

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| tboerger | <thomas@webhippie.de> | <https://github.com/tboerger> |

## Source Code

* <https://github.com/cloudhippie/charts>
* <https://github.com/hetznercloud/hcloud-cloud-controller-manager>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Affinity for the deployment |
| annotations | object | `{}` | Define additional annotations |
| config.clusterCidr | string | `"10.1.0.0/16"` | CIDR of the cluster network |
| config.kubeconfig | string | `nil` | Name of the secret containing a kubeconfig |
| config.network | string | `nil` | Name of the private cloud network |
| extraEnvVariables | list | `[{"name":"HCLOUD_INSTANCES_ADDRESS_FAMILY","value":"dualstack"},{"name":"HCLOUD_LOAD_BALANCERS_ENABLED","value":"true"},{"name":"HCLOUD_LOAD_BALANCERS_NETWORK_ZONE","value":"eu-central"},{"name":"HCLOUD_LOAD_BALANCERS_DISABLE_IPV6","value":"false"},{"name":"HCLOUD_LOAD_BALANCERS_DISABLE_PRIVATE_INGRESS","value":"true"},{"name":"HCLOUD_LOAD_BALANCERS_USE_PRIVATE_IP","value":"true"},{"name":"HCLOUD_NETWORK_DISABLE_ATTACHED_CHECK","value":"true"}]` | Extra environment variables for controller |
| extraInitContainers | list | `[]` | List of extra init containers |
| extraSidecarContainers | list | `[]` | List of extra sidecar containers |
| extraVolumeMounts | list | `[]` | List of extra volume mounts |
| extraVolumes | list | `[]` | List of extra volumes |
| fullnameOverride | string | `""` | Override the fullname |
| image.controller.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| image.controller.repository | string | `"hetznercloud/hcloud-cloud-controller-manager"` | Image repository used by deployment |
| image.controller.tag | string | `""` | Optional tag for the repository, defaults to app version |
| image.pullSecrets | list | `[]` | Optional name of pull secret if using a private registry |
| labels | object | `{}` | Define additional labels |
| metrics.enabled | bool | `true` | Enable metrics |
| metrics.internalPort | int | `8233` | Internal metrics port of the service |
| metrics.port | int | `8233` | Metrics port of the service |
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
| nodeSelector | object | `{}` | Node selector for the deployment |
| podSecurityContext.controller | object | `{}` | Security context for the controller pod |
| replicaCount | int | `1` | Replicas for the deployment |
| resources | object | `{"controller":{"limits":{"cpu":"1000m","memory":"640Mi"},"requests":{"cpu":"100m","memory":"64Mi"}}}` | Resources for the deployment |
| securityContext | object | `{}` | Security context for the deployment |
| service.annotations | object | `{}` | Additional annotations for the service |
| service.labels | object | `{}` | Additional labels for the service |
| service.type | string | `"ClusterIP"` | Type of the service |
| serviceAccount.annotations | object | `{}` | Define annotations for the service account |
| serviceAccount.create | bool | `true` | Create a new service account |
| serviceAccount.name | string | `""` | Optional name for an existing service account |
| tolerations | list | `[]` | Tolerations for the deployment |
| updateStrategy | object | `{"type":"Recreate"}` | Updaqte strategy for deployment |
