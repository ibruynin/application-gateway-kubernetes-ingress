# Greenfield Deployment

The instructions below assume Application Gateway Ingress Controller (AGIC) will be
installed in an environment with no pre-existing components.

### Required Command Line Tools

We recommend the use of [Azure Cloud Shell](https://shell.azure.com/) for all command line operations below. Launch your shell from shell.azure.com or by clicking the link:

[![Embed launch](https://shell.azure.com/images/launchcloudshell.png "Launch Azure Cloud Shell")](https://shell.azure.com)

Alternatively, launch Cloud Shell from Azure portal using the following icon:

![Portal launch](../portal-launch-icon.png)

Your [Azure Cloud Shell](https://shell.azure.com/) already has all necessary tools. Should you
choose to use another environment, please ensure the following command line tools are installed:

1. `az` - Azure CLI: [installation instructions](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)
1. `kubectl` - Kubernetes command-line tool: [installation instructions](https://kubernetes.io/docs/tasks/tools/install-kubectl)
1. `helm` - Kubernetes package manager: [installation instructions](https://github.com/helm/helm/releases/latest)
1. `jq` - command-line JSON processor: [installation instructions](https://stedolan.github.io/jq/download/)


### Deploy AKS with Azure network plugin

```bash
group="ingress-appgw"
clusterName="aks-cluster"
location="westeurope"
az group create -g $group -l "$location"

# add --disable-rbac if you want to deploy k8s without rbac
az aks create -g $group -n $clusterName --node-count 3 --network-plugin azure

# get cluster credentials
az aks get-credentials -g $group -n $clusterName
```

### Deploy User assigned Identity for AGIC with requried RBACs in Agent Pool resource group

```bash
identityName="agic-identity"
nodeResourceGroupName=$(az aks show -g $group -n $clusterName --query "nodeResourceGroup" -o tsv)
nodeResourceGroupId=$(az group show -g $nodeResourceGroupName --query "id" -o tsv)

az identity create -n $identityName -g $nodeResourceGroupName -l $location
identityPrincipalId=$(az identity show -n $identityName -g $nodeResourceGroupName --query "principalId" -o tsv)
identityResourceId=$(az identity show -n $identityName -g $nodeResourceGroupName --query "id" -o tsv)
identityClientId=$(az identity show -n $identityName -g $nodeResourceGroupName --query "clientId" -o tsv)

# if this fails with "No matches in graph database for ...", then try again.
az role assignment create \
        --role Contributor \
        --assignee "$identityPrincipalId" \
        --scope "$nodeResourceGroupId"
```

## Set up Application Gateway Ingress Controller

With the instructions in the previous section we created and configured a new AKS cluster and
an App Gateway. We are now ready to deploy a sample app and an ingress controller to our new
Kubernetes infrastructure.

### Install AAD Pod Identity (skip if already installed)
  Azure Active Directory Pod Identity provides token-based access to
  [Azure Resource Manager (ARM)](https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-overview).

  [AAD Pod Identity](https://github.com/Azure/aad-pod-identity) will add the following components to your Kubernetes cluster:
   1. Kubernetes [CRDs](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/): `AzureIdentity`, `AzureAssignedIdentity`, `AzureIdentityBinding`
   1. [Managed Identity Controller (MIC)](https://github.com/Azure/aad-pod-identity#managed-identity-controllermic) component
   1. [Node Managed Identity (NMI)](https://github.com/Azure/aad-pod-identity#node-managed-identitynmi) component


  To install AAD Pod Identity to your cluster:

   - *RBAC enabled* AKS cluster

  ```bash
  kubectl create -f https://raw.githubusercontent.com/Azure/aad-pod-identity/master/deploy/infra/deployment-rbac.yaml
  ```

   - *RBAC disabled* AKS cluster

  ```bash
  kubectl create -f https://raw.githubusercontent.com/Azure/aad-pod-identity/master/deploy/infra/deployment.yaml
  ```

### Install Helm (skip if already installed)
[Helm](https://docs.microsoft.com/en-us/azure/aks/kubernetes-helm) is a package manager for
Kubernetes. We will leverage it to install the `application-gateway-kubernetes-ingress` package:

1. Install [Helm](https://docs.microsoft.com/en-us/azure/aks/kubernetes-helm) and run the following to add `application-gateway-kubernetes-ingress` helm package:

    - *RBAC enabled* AKS cluster

    ```bash
    kubectl create serviceaccount --namespace kube-system tiller-sa
    kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller-sa
    helm init --tiller-namespace kube-system --service-account tiller-sa
    ```

    - *RBAC disabled* AKS cluster

    ```bash
    helm init
    ```

1. Add the AGIC Helm repository:
    ```bash
    helm repo add application-gateway-kubernetes-ingress https://appgwingress.blob.core.windows.net/ingress-azure-helm-package/
    helm repo update
    ```

### Install Ingress Controller Helm Chart

    ```bash
    applicationGatewayName="appgw"
    subnetPrefix="10.1.0.0/16"

    helm install application-gateway-kubernetes-ingress/ingress-azure \
      --set appgw.name=$applicationGatewayName \
      --set appgw.subnetPrefix=$subnetPrefix \
      --set armAuth.type=aadPodIdentity \
      --set armAuth.identityResourceID=$identityResourceId \
      --set armAuth.identityClientID=$identityClientId \
      --version 0.10.0-rc5
    ```

### Install a Sample App
Now that we have App Gateway, AKS, and AGIC installed we can install a sample app
via [Azure Cloud Shell](https://shell.azure.com/):

```yaml
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: aspnetapp
  labels:
    app: aspnetapp
spec:
  containers:
  - image: "mcr.microsoft.com/dotnet/core/samples:aspnetapp"
    name: aspnetapp-image
    ports:
    - containerPort: 80
      protocol: TCP

---

apiVersion: v1
kind: Service
metadata:
  name: aspnetapp
spec:
  selector:
    app: aspnetapp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 80

---

apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: aspnetapp
  annotations:
    kubernetes.io/ingress.class: azure/application-gateway
spec:
  rules:
  - http:
      paths:
      - path: /
        backend:
          serviceName: aspnetapp
          servicePort: 80
EOF
```

Alternatively you can:

1. Download the YAML file above:
```bash
curl https://raw.githubusercontent.com/Azure/application-gateway-kubernetes-ingress/master/docs/examples/aspnetapp.yaml -o aspnetapp.yaml
```

2. Apply the YAML file:
```bash
kubectl apply -f aspnetapp.yaml
```


## Other Examples
The **[tutorials](../tutorial.md)** document contains more examples on how toexpose an AKS
service via HTTP or HTTPS, to the Internet with App Gateway.
