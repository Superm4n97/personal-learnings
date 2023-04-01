# Crossplane
Crossplane connects your Kubernetes cluster to external non kubernetes resources, and allow platform teams to build custom Kubernetes APIs to consume those resources.

## Resources:
* Official site: https://www.crossplane.io/

## Notes:
* ### Provider:
Provider resource is used to connect with the cloud provider.
* ### Provider Config:
Provider config is used to as the settings for the cloud provider. It contains the provider specific information and configuration.
* ### Managed Resources:
* ### Compositions:
* ### Composite Resources:
* ### Composite Resources Definitions:
* ### Claims:

## Installation:
### AWS:
* For installation: https://docs.crossplane.io/v1.11/getting-started/provider-aws/

## Workflow:
1. install `Crossplane`
2. install provider (`upbound provider-aws`)
3. create credential `secret`
4. create `providerConfig` with credential `secret`