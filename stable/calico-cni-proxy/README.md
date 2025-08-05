# calico-cni-proxy

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.0](https://img.shields.io/badge/AppVersion-0.0.0-informational?style=flat-square)

Calico CNI component for as helm chart proxy deployment

**Homepage:** <https://github.com/projectcalico/calico>

**This chart is not maintained by the upstream project and any issues with the
chart should be raised [here](https://github.com/cloudhippie/charts/issues/new)**

## Installing the Chart

### OCI (Recommended)

```console
helm install calico-cni-proxy oci://ghcr.io/cloudhippie/charts/calico-cni-proxy
```

### Traditional

```console
helm repo add cloudhippie https://cloudhippie.github.io/charts
helm repo update
helm install calico-cni-proxy cloudhippie/calico-cni-proxy
```

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| tboerger | <thomas@webhippie.de> | <https://github.com/tboerger> |

## Source Code

* <https://github.com/cloudhippie/charts>
* <https://github.com/projectcalico/calico>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| annotations | object | `{}` | Define additional annotations |
| fullnameOverride | string | `""` | Override the fullname |
| labels | object | `{}` | Define additional labels |
| nameOverride | string | `""` | Override the name |
| target.namespace | string | `"tigera-operator"` | Namespace for chart in chart proxy |
| target.release | string | `"calico-cni"` | Release name for chart in chart proxy |
| target.selector | object | `{"matchLabels":{}}` | Matching for the target cluster |
| upstream.version | string | `"v3.30.2"` | Upstream chart version |
