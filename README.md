# kubectl-patch-all

`kubectl-patch-all` is a kubectl plugin which wraps around `kubectl get` and `kubectl patch` to make it easy to patch resources as a batch

## Installing

`go install .`

This installs the binary to your `$GOPATH`, kubectl autodetects plugins based on name, provided the plugin is in your `$PATH` (e.g `export GOPATH="$HOME/go"`, `export PATH="$GOPATH/bin:$PATH"`)

## Usage

```
# Patch all pods in the default namespace
kubectl patch-all pods --patch-file patch.yaml
```
```
# Patch all deploys which meet the label selector in the target namespace and context
kubectl patch-all deploys --patch-file patch.yaml --context some-cluster -n some-ns -l some-label=some-value
```
```
# Patch Application custom resources with the 'merge' patch type rather than the default 'strategic'
# Some CRDs do not support the default strategic patching
kubectl patch-all applications --patch-file patch.yaml --type merge
```
```
# Print configmaps which would be targeted for a patch, but do not patch
kubectl patch-all configmaps --patch-file patch.yaml --dry-run
```
