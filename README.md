# k8sfmt - Opinionated YAML formatter for GitOps

Like `gofmt` but for YAML files with a focus on readability of Kubernetes manifests, such as `ConfigMaps` holding configuration files.

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

- Indents by 2 spaces
- Renders multiline strings using `|2` [literal style](https://yaml.org/spec/1.2.2/#literal-style) to improve readability
    - No `\"escaped\"` strings and `\n\n` newlines
- Sorts keys alphabetically
- Keeps comments

```diff
 apiVersion: v1
 kind: ConfigMap
 metadata:
   name: game-demo
 data:
-  config.toml: "\n### Service configuration\n\n[server]\n  port          = 8000\n  read_timeout  = \"10s\"\n  write_timeout = \"10s\"\n"
+  config.toml: |2
+   
+   ### Service configuration
+   
+   [server]
+     port          = 8000
+     read_timeout  = "10s"
+     write_timeout = "10s"
+     
```
