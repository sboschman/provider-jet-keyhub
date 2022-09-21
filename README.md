# Terrajet KeyHub Provider

`provider-jet-keyhub` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for the
KeyHub API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/crossplane-contrib/provider-jet-keyhub/releases):
```
kubectl crossplane install provider crossplane/provider-jet-keyhub:v0.1.0
```

Alternatively, you can use declarative installation:
```
kubectl apply -f examples/install.yaml
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-jet-keyhub).

## Developing

Clone this repository and cd into the repository directory. Fetch the upbound/build submodule by running the following:

```console
make submodules
```

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/topicuskeyhub/provider-jet-keyhub/issues).

## Governance and Owners

provider-jet-keyhub is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-keyhub adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-keyhub is under the Apache 2.0 license.
