---
date: 2017-012-17T00:00:00+00:00
title: koki
author: koki
tags: [ kubernetes, koki, manifests ]
repo: kubeciio/koki
image: kokster/kubeci-plugin:v0.3.0
---

#### Introduction

The koki plugin for Drone/KubeCI can be used to convert Short manifests into Kubernetes manifests on the fly in the CI pipeline. 

This will enable teams to exclusively use the Short format for representing their Kubernetes API objects.

#### Options

This plugin supports the following options

| Option | Type | Description | 
|--------|------|-------------|
| files  | []string | Input files relative to root of the project which is being built using drone |
| overwrite | bool | Set to `true` to allow output files to be overwritten. (default `false`) |
| in_place | bool | Set to `true` to translate files in place. (default `false`). Should always be used with `overwrite: true` |
| prefix | string | The prefix of the output file created. (default `kube_`) |

The plugin preserves the directory of the files, as it translates them. 

#### Configuring .drone.yml for this plugin

In its simplest form, the plugin translate a short file `a.yaml` to its kubernetes equivalent `kube_a.yaml`

```yaml
workspace:
  base: /go
  path: src/github.com/kubeciio/koki

  pipeline:
    koki-short:
      image: kokster/kubeci-plugin:v0.3.0
      files:
      - deployment.yaml
```

This would result in the next steps of the pipeline having access to `kube_deployment.yaml`

###### Translating files in place

In some cases, you'd like your input yaml file and output yaml file to be the same. Since the build happens in CI, it won't affect your Github repositry if you enable this option.

```yaml
workspace:
  base: /go
  path: src/github.com/kubeciio/koki

  pipeline:
    koki-short:
      image: kokster/kubeci-plugin:v0.3.0
      files:
      - deployment.yaml
      in-place: true
      overwrite: true      # note that in-place will fail unless overwrite is set to true
```

#### Changing the prefix of the output file

If you wish to change the default prefix from `kube_` to something else, then you can achieve this using the template below

```yaml
workspace:
  base: /go
  path: src/github.com/kubeciio/koki

  pipeline:
    koki-short:
      image: kokster/kubeci-plugin:v0.3.0
      files:
      - deployment.yaml
      prefix: custom_
``` 
