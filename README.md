# kubectl-patch-all

`kubectl-patch-all` is a kubectl plugin which wraps around `kubectl get` and `kubectl patch` to make it easy to patch resources as a batch

## Installing

`go install .`

This installs the binary to your `$GOPATH`, kubectl autodetects plugins based on name, provided the plugin is in your `$PATH` (e.g `export GOPATH="$HOME/go"`, `export PATH="$GOPATH/bin:$PATH"`)

## Usage

`kubectl patch-all pods --patch-file patch.yaml`
`kubectl patch-all deploys --patch-file patch.yaml --context my-cluster --namespace my-ns`
`kubectl patch-all pods --patch-file patch.yaml --dry-run`
