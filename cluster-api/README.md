### Learn the concepts about cluster api:
* https://www.youtube.com/watch?v=9H8flXm_lKk
* https://cluster-api.sigs.k8s.io/user/concepts.html
* https://cluster-api.sigs.k8s.io/user/quick-start.html
* For AWS provider: https://cluster-api-aws.sigs.k8s.io

### Keywords:
management cluster, workload cluster, <br>
machine, machineset, machinedeployment, kubeadmcontrolplane, clusterclass, cni, kubeconfig, nodes, <br>
control plane machine, workload machine, <br>

### Kubernetes Cluster API Provider AWS (CAPA)
Two ways to achieve the kubernetes cluster using CAPA. One way by creating virtual machines for master and worker nodes and manually initializing them as cloud machine (cloud init).
Other way is by using AWS EKS support.

#### Currently, EKS based Cluster supports:
* Provisioning and managing an Amazon EKS Cluster.
* Upgrading the Kubernetes version of the EKS Cluster.
* Attaching a single or a machine pool to the EKS Cluster.
* Managing "EKS Addons" and aws-iam-authenticator configuration
* Creating an EKS fargate profile

