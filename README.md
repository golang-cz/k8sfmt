# kubefmt - YAML formatter for K8s manifests

Opinionated formatter for GitOps repositories. Prioritizes readability of K8s manifests, e.g.`ConfigMap` holding longer configuration files.

No config or CLI options. Take it or leave it.

## Usage

```
$ kubefmt *.yaml
```

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

## Install

```
$ go install github.com/golang-cz/kubefmt/cmd/kubefmt@latest
```

## Features

- Indents by 2 spaces
- Renders multiline strings using `|2` [literal style](https://yaml.org/spec/1.2.2/#literal-style) to improve readability
    - No `\"escaped\"` strings and `\n\n` newlines
- Sorts keys alphabetically
- Keeps comments


## License

[MIT license](./LICENSE)
