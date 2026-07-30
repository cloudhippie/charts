# Cloudhippie: Charts

[![Release Build](https://github.com/cloudhippie/charts/actions/workflows/release.yml/badge.svg)](https://github.com/cloudhippie/charts/actions/workflows/release.yml) [![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/cloudhippie)](https://artifacthub.io/packages/search?repo=cloudhippie)

Definition and publishing of Helm charts to install tools used by Cloudhippie or
tools built within our GitHub organization.

## Prerequisites

We use [mise][mise] to manage all required tools and their versions. Install it
by following the [official installation instructions][mise-install], then run
the following commands inside the repository to activate mise and install all
tools defined in `mise.toml`:

```console
mise trust
mise install
```

## Usage

Make sure you have installed [Helm][helm], after that you can install the charts
repository and search for available charts:

```console
helm repo add cloudhippie https://cloudhippie.github.io/charts
helm search repo cloudhippie
```

## Security

If you find a security issue please contact
[security@cloudhippie.de](mailto:security@cloudhippie.de) first.

## Contributing

Generally we are following [conventional commits][commits] when we apply
changes. That way we are able to generate proper changelogs for every release.
Please use always pull requests to integrate new functionalities or to fix
issues.

For the release process we are following [semantic versioning][semver] which
clearly indicates if a new version just resolves bugs, includes new features or
even includes breaking changes.

After installing the tools via `mise install` as described above set up the
pre-commit hooks so they run automatically on every commit:

```console
pre-commit install --hook-type pre-commit --hook-type commit-msg
```

> `pre-commit` is managed by mise and will be available after `mise install`.

If you have changed something on the source you should simply commit following
the mentioned conventions:

```console
git checkout -b feat/new-feature
git add --all
git commit -m 'feat: added awesome new feature'
git push --set-upstream origin feat/new-feature
```

After pushing your changes into the Git repository you should create a pull
request on GitHub. If the pull request have been merged and everything built
fine it will also create automatically a new release at least once a week.

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2023 Cloudhippie <info@cloudhippie.de>
```

[mise]: https://mise.jdx.dev/
[mise-install]: https://mise.jdx.dev/getting-started.html
[helm]: https://helm.sh
[commits]: https://www.conventionalcommits.org/en/v1.0.0/
[semver]: https://semver.org/
