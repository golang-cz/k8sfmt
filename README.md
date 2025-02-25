# k8sfmt - Opinionated YAML formatter for Kubernetes manifests

Like `gofmt` but for YAML files with a focus on `ConfigMap` manifests readability, e.g. in GitOps repositories.

No config or CLI options. Take it or leave it.

## Install

```
$ go install github.com/golang-cz/k8sfmt/cmd/k8sfmt@latest
```

## Usage

```
$ k8sfmt *.yml
```

## Features

- Converts multi-line strings to block scalars
```

```
