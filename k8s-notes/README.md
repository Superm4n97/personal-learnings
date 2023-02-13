3 processes must be installed in every node to make it a worker node
1. Container Runtime
2. Kubelet
3. KubeProxy

four master processes must be installed in a node to make it a master node
1. API server
2. Scheduler
3. Controller manger
4. etcd

## k8s architecture


* If a pod sent a request to another pod, the service first looks for the pod in the same node (machine) to avoid network overhead.
* 
