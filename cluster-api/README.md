## Learn the concepts about cluster api:
* https://www.youtube.com/watch?v=9H8flXm_lKk
* https://cluster-api.sigs.k8s.io/user/concepts.html
* https://cluster-api.sigs.k8s.io/user/quick-start.html
* For AWS provider: https://cluster-api-aws.sigs.k8s.io
* For Azure provider: https://github.com/kubernetes-sigs/cluster-api-provider-azure

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

use `cluster init` to install provider component 
```bash
cluster move --to-kubeconfig=eks.kubeconfig
```

### 1.2 VPC peering concept
https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html

With VPC peering you can connect two AWS VPC's in a same network. So any instance within a VPC can communicate with any instance of other VPC as if they are in a same cluster. With VPC peering you can connect two VPC's from same account or from different account. It also can connect VPC's across different region. The traffic remains within the AWS global backbone without exposing it to the internet.

* **prerequisite**
1. VPC's CIDR block can not overlap with each other.

* **peering process**
1. there are requester and accepter VPC, they establish connection between then using a confirmation request.
2. both VPC should update their route table so that they can send and receive traffic to/from other VPC.
3. if required, update the ingress/egress rule in your security group rule so that the traffic exchange of your vpc in not restricted.
4. with the default peering connection option instance on either side of a vpc peering connection address each other using a public DNS hostname, the hostname resolves the public IP of the instance. To change this behavior, enable DNS hostname resolution for your VPC connection. Then the hostname will resolve to the private IP address.

https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-basics.html

