# Drone Plugin for Koki Short

This is a drone plugin meant to be run in the pipeline when Koki Short files are used. This plugin can read [Koki Short](https://www.koki.io) files from your repository and translate them on the fly to Kubernetes manifest files. 

Koki files provide a better user experience for users writing and reading Kubernetes manifests. 

Learn more about it here - https://docs.koki.io/short/user-guide/drone/

## Usage

This plugin supports the following options

| Option | Type | Description | 
|--------|------|-------------|
| files  | []string | Input files relative to root of the project which is being built using drone |
| overwrite | bool | Set to `true` to allow output files to be overwritten. (default `false`) |
| in_place | bool | Set to `true` to translate files in place. (default `false`). Should always be used with `overwrite: true` |
| prefix | string | The prefix of the output file created. (default `kube_`) |
| short-path | string | The path to the short binary |

The plugin preserves the directory of the files, as it translates them. 

## Examples

```bash
# Provide input files using flag
koki -f test1.yaml -f test2.yaml

# Provide input files using environment variable
PLUGIN_FILES=test1.yaml,test2.yaml koki 

# Provide output file prefix using flag
koki -f test1.yaml -p k8s_

# Provide output file prefix using environment variable
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml koki

# Translate file in-place using flag
koki -f test1.yaml -i

# Translate file in-place using environment variable. Note that in order to convert in-place, overwrite should be set to true 
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml PLUGIN_IN_PLACE=true PLUGIN_OVERWRITE=true koki

# Preserves the directory in which the original file was found
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=/path/to/dir/test1.yaml,test2.yaml PLUGIN_IN_PLACE=true koki
 >  output file will be created in /path/to/dir/ and in current directory
 
# Set overwrite=true if output file already exists (using env vars)
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml PLUGIN_OVERWRITE=true koki

# Set overwrite=true if output file already exists (using flag)
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml koki -w
```

## Setup and Usage

Get the project

```bash
go get github.com/kubeciio/koki
```

Build it

```bash
./scripts/build.sh
```

The build artifact will be found in `bin/koki-short-drone-plugin`

Build the Docker image

```bash
./scripts/release.sh
```

The docker image kokster/kubeci-plugin:{TAG} will be created.

## Testing with docker image

```bash
$ cat deployment.yaml
deployment:
  containers:
  - expose:
    - 80
    image: example/example:latest
    name: example-app
  labels:
    app: example
  name: example-deployment
  replicas: 3
  selector:
    app: example
  version: extensions/v1beta1

$ docker run -v ${PWD}:/dir kokster/kubeci-plugin:0.3.0 -f /dir/deployment.yaml
INFO[0000] Waiting for all translations to complete     
INFO[0000] Translating file=[deployment.yaml] prefix=[kube_] inPlace=[false] overwrite=[true] 
INFO[0000] Successfully translated all input files 

$ ls kube_deployment.yaml    # output file will be created in the same directory
kube_deployment.yaml

$ cat kube_deployment.yaml  # this will be the kubernets equivalent
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: example
  name: example-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: example
  strategy:
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: example
    spec:
      containers:
      - image: example/example:latest
        name: example-app
        ports:
        - containerPort: 80
          protocol: TCP
        resources: {}
status: {}
```
