# hcloud-csi-proxy

![Version: 1.0.2](https://img.shields.io/badge/Version-1.0.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.0](https://img.shields.io/badge/AppVersion-0.0.0-informational?style=flat-square)

Hetzner Cloud CSI component as helm chart proxy deployment

**Homepage:** <https://github.com/hetznercloud/csi-driver>

**This chart is not maintained by the upstream project and any issues with the
chart should be raised [here](https://github.com/cloudhippie/charts/issues/new)**

## Installing the Chart

### OCI (Recommended)

```console
helm install hcloud-csi-proxy oci://ghcr.io/cloudhippie/charts/hcloud-csi-proxy
```

### Traditional

```console
helm repo add cloudhippie https://cloudhippie.github.io/charts
helm repo update
helm install hcloud-csi-proxy cloudhippie/hcloud-csi-proxy
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
| annotations | object | `{}` | Define additional annotations |
| downstream.version | string | `"1.0.2"` | Downstream chart version |
| fullnameOverride | string | `""` | Override the fullname |
| labels | object | `{}` | Define additional labels |
| nameOverride | string | `""` | Override the name |
| target.kamajiCluster | bool | `false` | Enabled if deployed to Kamaji cluster |
| target.namespace | string | `"kube-system"` | Namespace for chart in chart proxy |
| target.release | string | `"hcloud-csi"` | Release name for chart in chart proxy |
| target.selector | object | `{"matchLabels":{}}` | Matching for the target cluster |
| upstream.version | string | `"2.16.0"` | Upstream chart version |
