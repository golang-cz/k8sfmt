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

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: game-demo
data:
  # property-like keys; each key maps to a simple value
  player_initial_lives: "3" # comment
  ui_properties_file_name: "user-interface.properties"

  # file-like keys
  game.properties: |
    enemy.types=aliens,monsters
    player.maximum-lives=5    
  user-interface.properties: |
    color.good=purple
    color.bad=yellow
    allow.textmode=true
```
