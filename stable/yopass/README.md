# yopass

![Version: 9.13.0](https://img.shields.io/badge/Version-9.13.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 14.8.0](https://img.shields.io/badge/AppVersion-14.8.0-informational?style=flat-square)

Secure sharing of secrets, passwords and files

**Homepage:** <https://yopass.se/>

**This chart is not maintained by the upstream project and any issues with the
chart should be raised [here](https://github.com/cloudhippie/charts/issues/new)**

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
  enabled: true

  hosts:
    - host: yopass.example.com
      paths:
        - path: /
          pathType: Prefix
```

### Redis Storage

> **Warning:** Use this deployment only for testing, we recommend a proper
> deployment based on some operator to have a real lifecycle management.

```console
database:
  type: redis
  dsn: redis://redis:6379/0

redis:
  enabled: true

memcached:
  enable: false
```

### Memcached Storage

> **Warning:** Use this deployment only for testing, we recommend a proper
> deployment based on some operator to have a real lifecycle management.

```console
database:
  type: memcached
  dsn: memcached:11211

redis:
  enabled: false

memcached:
  enable: true
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
| oci://registry-1.docker.io/cloudpirates | memcached | 0.14.3 |
| oci://registry-1.docker.io/cloudpirates | redis | 0.33.2 |

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
| extraInitContainers | list | `[]` | List of extra init containers |
| extraSidecarContainers | list | `[]` | List of extra sidecar containers |
| extraVolumeMounts | list | `[]` | List of extra volume mounts |
| extraVolumes | list | `[]` | List of extra volumes |
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
| memcached.enabled | bool | `false` | Enable memcached dependency |
| memcached.fullnameOverride | string | `"memcached"` | Override fullname of memcached dependency |
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
| networkPolicy.allowExternalEgress | bool | `true` | Allow any egress traffic to external sources |
| networkPolicy.allowExternalIngress | bool | `true` | Allow any ingress traffic from external sources |
| networkPolicy.egressMemcachedMatch | list | `[{"podSelector":{"matchLabels":{"app.kubernetes.io/name":"memcached"}}}]` | Match to allow egress traffic to Memcached service |
| networkPolicy.egressMemcachedPort | int | `11211` | Egress port for Memcached service |
| networkPolicy.egressRedisMatch | list | `[{"podSelector":{"matchLabels":{"app.kubernetes.io/name":"redis"}}}]` | Match to allow egress traffic to Redis service |
| networkPolicy.egressRedisPort | int | `6379` | Egress port for Redis service |
| networkPolicy.enabled | bool | `false` | Specifies whether a NetworkPolicy should be created |
| networkPolicy.extraEgress | list | `[]` | List of extra egress rules to the NetworkPolicy |
| networkPolicy.extraIngress | list | `[]` | List of extra ingress rules to the NetworkPolicy |
| networkPolicy.generalEgress | list | `[{"ports":[{"port":53,"protocol":"UDP"},{"port":53,"protocol":"TCP"}]}]` | List of general egress rules to the NetworkPolicy |
| networkPolicy.ingressMetricsMatch | list | `[]` | Match to allow ingress traffic to metrics port |
| networkPolicy.ingressServiceMatch | list | `[]` | Match to allow ingress traffic to service port |
| nodeSelector | object | `{}` | Node selector for the deployment |
| podSecurityContext | object | `{}` | Security context for the pod |
| readOnly | object | `{"affinity":{},"enabled":false,"envFromConfigMaps":[],"envFromSecrets":[],"extraEnvSecrets":{},"extraEnvVariables":{},"ingress":{"annotations":{},"className":null,"enabled":false,"hosts":[{"host":"example.local","paths":[{"path":"/","pathType":"Prefix"}]}],"labels":{},"tls":[]},"nodeSelector":{},"replicaCount":null,"resources":{},"service":{"annotations":{},"enabled":true,"internalPort":null,"labels":{},"port":null,"type":null},"tolerations":[]}` | Optional second deployment in read-only mode |
| readOnly.affinity | object | `{}` | Affinity for the read-only deployment (defaults to .Values.affinity) |
| readOnly.enabled | bool | `false` | Enable rendering a second, read-only instance (resources suffixed with -readonly) |
| readOnly.envFromConfigMaps | list | `[]` | List of environment variables from existing configmaps for read-only instance |
| readOnly.envFromSecrets | list | `[]` | List of environment variables from existing secrets for read-only instance |
| readOnly.extraEnvSecrets | object | `{}` | Extra environment variables from secrets for read-only instance |
| readOnly.extraEnvVariables | object | `{}` | Extra environment variables for read-only instance |
| readOnly.ingress.annotations | object | `{}` | Additional annotations for the read-only ingress |
| readOnly.ingress.className | string | `nil` | Class name for the ingress resource |
| readOnly.ingress.enabled | bool | `false` | Enable ingress for read-only instance |
| readOnly.ingress.hosts | list | `[{"host":"example.local","paths":[{"path":"/","pathType":"Prefix"}]}]` | Host definition for read-only ingress |
| readOnly.ingress.labels | object | `{}` | Additional labels for the read-only ingress |
| readOnly.ingress.tls | list | `[]` | Optional TLS configuration for read-only ingress |
| readOnly.nodeSelector | object | `{}` | Node selector for the read-only deployment (defaults to .Values.nodeSelector) |
| readOnly.replicaCount | int | `nil` | Replicas for the read-only deployment (defaults to .Values.replicaCount) |
| readOnly.resources | object | `{}` | Resources for the read-only deployment (defaults to .Values.resources) |
| readOnly.service.annotations | object | `{}` | Additional annotations for the read-only service |
| readOnly.service.enabled | bool | `true` | Enable Service for read-only instance |
| readOnly.service.internalPort | int | `nil` | Internal port of the read-only service (defaults to .Values.service.internalPort) |
| readOnly.service.labels | object | `{}` | Additional labels for the read-only service |
| readOnly.service.port | int | `nil` | Port of the read-only service (defaults to .Values.service.port) |
| readOnly.service.type | string | `nil` | Type of the read-only service (defaults to .Values.service.type) |
| readOnly.tolerations | list | `[]` | Tolerations for the read-only deployment (defaults to .Values.tolerations) |
| redis.auth.enabled | bool | `false` |  |
| redis.enabled | bool | `false` | Enable redis dependency |
| redis.fullnameOverride | string | `"redis"` | Override fullname of redis dependency |
| replicaCount | int | `1` | Replicas for the deployment |
| resources | object | `{"limits":{},"requests":{"cpu":"100m","memory":"64Mi"}}` | Resources for the deployment |
| securityContext | object | `{}` | Security context for the deployment |
| service.annotations | object | `{}` | Additional annotations for the service |
| service.internalPort | int | `3000` | Internal port of the service |
| service.labels | object | `{}` | Additional labels for the service |
| service.port | int | `3000` | Port of the service |
| service.type | string | `"ClusterIP"` | Type of the service |
| serviceAccount.annotations | object | `{}` | Define annotations for the service account |
| serviceAccount.automountToken | bool | `false` | Automount service account token |
| serviceAccount.create | bool | `true` | Create a new service account |
| serviceAccount.name | string | `""` | Optional name for an existing service account |
| tolerations | list | `[]` | Tolerations for the deployment |
| updateStrategy | object | `{"type":"Recreate"}` | Updaqte strategy for deployment |

## Read-Only Mode & Split Deployments

Read-only mode disables all secret creation endpoints while keeping retrieval fully functional. This chart can deploy an optional second instance in read-only mode within the same release.

### Enable read-only instance

```yaml
readOnly:
  enabled: true
  ingress:
    enabled: true
    hosts:
      - host: yopass.example.com
        paths:
          - path: /
            pathType: Prefix
```

Resources created for the read-only instance are suffixed with `-readonly` (Deployment, Service, Ingress). Both instances share the same database configured via `.Values.database.*`.

In read-only mode the server is started with `--read-only`. Endpoint behavior:

| Endpoint | Behavior |
|----------|----------|
| `POST /create/secret` | 404 Not Found |
| `POST /create/file` | 404 Not Found |
| `GET /secret/{key}` | Active |
| `GET /file/{key}` | Active |
| `DELETE /secret/{key}` | Active |

The frontend detects `READ_ONLY: true` from the `/config` endpoint and shows a read-only landing page instead of the create form.

### Split deployment pattern (one release)

Deploy one Helm release with both instances sharing one database:

```yaml
database:
  type: memcached
  dsn: memcached:11211

memcached:
  enabled: true

ingress:
  enabled: true
  hosts:
    - host: internal.example.com
      paths:
        - path: /
          pathType: Prefix

readOnly:
  enabled: true
  ingress:
    enabled: true
    hosts:
      - host: yopass.example.com
        paths:
          - path: /
            pathType: Prefix
```

Both instances connect to the same Memcached or Redis database. Secrets created on the internal instance can be retrieved via the public read-only instance.

### OIDC on internal, public read-only

Require authentication only on the internal instance while the public instance remains open for retrieval:

```yaml
extraEnvVariables:
  REQUIRE_AUTH: "true"
  OIDC_ISSUER: https://accounts.google.com
  OIDC_CLIENT_ID: "123456789-abc.apps.googleusercontent.com"
  OIDC_CLIENT_SECRET: "GOCSPX-..."
  OIDC_REDIRECT_URL: https://internal.example.com/auth/callback
  OIDC_ALLOWED_DOMAIN: example.com

readOnly:
  enabled: true
  # no auth variables here
```

Notes:
- Deletion of one-time secrets happens automatically via `DELETE /secret/{key}` when a recipient opens a secret; this endpoint remains active in read-only mode intentionally.
- If you use a CDN or caching layer in front of the public instance, ensure retrieval endpoints (`GET /secret/*`, `GET /file/*`) are not cached.
