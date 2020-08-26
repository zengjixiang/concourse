## Setup with Kind + Concourse Web on Docker
1. brew install kind

1. kind-config.yml
	```
	kind: Cluster
	apiVersion: kind.x-k8s.io/v1alpha4
	networking:
	  # WARNING: It is _strongly_ recommended that you keep this the default
	  # (127.0.0.1) for security reasons. However it is possible to change this.
	  apiServerAddress: "0.0.0.0"
	  # By default the API server listens on a random open port.
	  # You may choose a specific port but probably don't need to in most cases.
	  # Using a random port makes it easier to spin up multiple clusters.
	  apiServerPort: 6443
	```

1. kind create cluster --kubeconfig kind/kubeconfig.yml --config kind/kind-config.yml

1. Determine the kind containers IP (eg. `docker network inspect kind | grep IPv4`). Update `kind/kubeconfig.yml` with the private IP of the kind container.

1. docker network connect kind concourse_web_1


## For local access to the kind cluster
 
```
	sed 's/https:\/\/.*6443/https:\/\/0.0.0.0:6443/' kind/kubeconfig.yml > /tmp/kubeconfig.yml
	export KUBECONFIG=/tmp/kubeconfig.yml
```

## To fix DNS within the pods
```
   kubectl edit -n kube-system configmap/coredns         # Replaced `forward . /etc/resolv.conf` with `forward . 8.8.8.8`
```


Note : See Makefile for additional setup like `init`