# hcloud-ccm-proxy

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.0](https://img.shields.io/badge/AppVersion-0.0.0-informational?style=flat-square)

Hetzner Cloud CCM component as helm chart proxy deployment

**Homepage:** <https://github.com/hetznercloud/hcloud-cloud-controller-manager>

**This chart is not maintained by the upstream project and any issues with the
chart should be raised [here](https://github.com/cloudhippie/charts/issues/new)**

## Installing the Chart

### OCI (Recommended)

```console
helm install hcloud-ccm-proxy oci://ghcr.io/cloudhippie/charts/hcloud-ccm-proxy
```

### Traditional

```console
helm repo add cloudhippie https://cloudhippie.github.io/charts
helm repo update
helm install hcloud-ccm-proxy cloudhippie/hcloud-ccm-proxy
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
| annotations | object | `{}` | Define additional annotations |
| fullnameOverride | string | `""` | Override the fullname |
| labels | object | `{}` | Define additional labels |
| nameOverride | string | `""` | Override the name |
| target.namespace | string | `"kube-system"` | Namespace for chart in chart proxy |
| target.release | string | `"hcloud-ccm"` | Release name for chart in chart proxy |
| target.selector | object | `{"matchLabels":{}}` | Matching for the target cluster |
| upstream.version | string | `"1.26.0"` | Upstream chart version |
