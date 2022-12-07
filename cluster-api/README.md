## Learn the concepts about cluster api:
* https://www.youtube.com/watch?v=9H8flXm_lKk
* https://cluster-api.sigs.k8s.io/user/concepts.html
* https://cluster-api.sigs.k8s.io/user/quick-start.html
* For AWS provider: https://cluster-api-aws.sigs.k8s.io

### Keywords:
management cluster, workload cluster, <br>
machine, machineset, machinedeployment, kubeadmcontrolplane, clusterclass, cni, kubeconfig, nodes, <br>
control plane machine, workload machine, <br>

## 1. Kubernetes Cluster API Provider AWS (CAPA)
AWS cluster can be various types
* Create machines/instances of basic instance type then make control pane (with kubeadmcontrolplane) and worker nodes of your own.
* Create EKS based control plane and add worker nodes as needed. In this process you don't need to concern about the control plane it will automatically scale up and down of its own.
* Create EKS based control plane and machine pool, and you don't need to managed something of your own, AWS takes care of everything.

### 1.1 Pivoting
Pivoting basically used for changing the management cluster. So for pivoting we have to move the provider component from source cluster to target cluster https://cluster-api.sigs.k8s.io/clusterctl/commands/move.html <br>

use `cluster init` to install provider component <br>
`cluster move --to-kubeconfig=eks.kubeconfig`