# yopass

![Version: 1.2.1](https://img.shields.io/badge/Version-1.2.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 11.15.1](https://img.shields.io/badge/AppVersion-11.15.1-informational?style=flat-square)

Secure sharing of secrets, passwords and files

**Homepage:** <https://yopass.se/>

## Installing the Chart

### OCI (Recommended)

```console
helm install yopass oci://ghcr.io/cloudhippie/charts/yopass
```

### Traditional

```console
helm repo add cloudhippie https://cloudhippie.github.io/charts
helm repo update
helm install yopass cloudhippie/yopass
```

## Example for Values

### Ingress Enabled

```console
ingress:
  enabled: false

  hosts:
    - host: yopass.example.com
      paths:
        - path: /
          pathType: Prefix
```

### Redis Storage

```console
database:
  type: redis

  dsn: redis://redis:6379/0

redis:
  enabled: true

memcached:
  enable: false
```

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| tboerger | <thomas@webhippie.de> | <https://github.com/tboerger> |

## Source Code

* <https://github.com/jhaals/yopass>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| oci://registry-1.docker.io/bitnamicharts | memcached | 6.6.7 |
| oci://registry-1.docker.io/bitnamicharts | redis | 18.2.1 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Affinity for the deployment |
| annotations | object | `{}` | Define additional annotations |
| config.maxLength | int | `10000` | Max length of encrypted secret |
| database.dsn | string | `"memcached:11211"` | DSN to access the database |
| database.type | string | `"memcached"` | Type of database backend |
| envFromConfigMaps | list | `[]` | List of environment variables from existing configmaps |
| envFromSecrets | list | `[]` | List of environment variables from existing secrets |
| extraEnvSecrets | object | `{}` | Extra environment variables from secrets |
| extraEnvVariables | object | `{}` | Extra environment variables from mapping |
| fullnameOverride | string | `""` | Override the fullname |
| image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| image.pullSecrets | list | `[]` | Optional name of pull secret if using a private registry |
| image.repository | string | `"jhaals/yopass"` | Image repository used by deployment |
| image.tag | string | `""` | Optional tag for the repository, defaults to app version |
| ingress.annotations | object | `{}` | Additional annotations for the ingress |
| ingress.className | string | `nil` | Class name for the ingress resource |
| ingress.enabled | bool | `false` | Enable ingress |
| ingress.hosts | list | `[{"host":"example.local","paths":[{"path":"/","pathType":"Prefix"}]}]` | Host definition for ingress |
| ingress.labels | object | `{}` | Additional labels for the ingress |
| ingress.tls | list | `[]` | Optional TLS configuration for ingress |
| labels | object | `{}` | Define additional labels |
| memcached.enabled | bool | `true` | Enable memcached dependency |
| memcached.fullnameOverride | string | `"memcached"` | Override fullname of memcached dependency |
| memcached.metrics.enabled | bool | `true` | Enable metrics for memcached |
| memcached.metrics.podAnnotations | string | `nil` | Enforce pod annotations for memcached |
| memcached.metrics.serviceMonitor.enabled | bool | `false` | Enable service monitor for memcached |
| memcached.serviceAccount.create | bool | `true` | Create service account for memcached |
| metrics.enabled | bool | `true` | Enable metrics |
| metrics.internalPort | int | `3001` | Internal metrics port of the service |
| metrics.port | int | `3001` | Metrics port of the service |
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
| podSecurityContext | object | `{}` | Security context for the pod |
| redis.enabled | bool | `false` | Enable redis dependency |
| redis.fullnameOverride | string | `"redis"` | Override fullname of redis dependency |
| redis.metrics.enabled | bool | `true` | Enable metrics for redis |
| redis.metrics.podAnnotations | string | `nil` | Enforce pod annotations for redis |
| redis.metrics.serviceMonitor.enabled | bool | `false` | Enable service monitor for redis |
| redis.serviceAccount.create | bool | `true` | Create service account for redis |
| replicaCount | int | `1` | Replicas for the deployment |
| resources | object | `{"limits":{},"requests":{"cpu":"100m","memory":"64Mi"}}` | Resources for the deployment |
| securityContext | object | `{}` | Security context for the deployment |
| service.annotations | object | `{}` | Additional annotations for the service |
| service.internalPort | int | `3000` | Internal port of the service |
| service.labels | object | `{}` | Additional labels for the service |
| service.port | int | `3000` | Port of the service |
| service.type | string | `"ClusterIP"` | Type of the service |
| serviceAccount.annotations | object | `{}` | Define annotations for the service account |
| serviceAccount.create | bool | `true` | Create a new service account |
| serviceAccount.name | string | `""` | Optional name for an existing service account |
| tolerations | list | `[]` | Tolerations for the deployment |
| updateStrategy | object | `{"type":"Recreate"}` | Updaqte strategy for deployment |
