# Brownfield Deployment

The App Gateway Ingress Controller (AGIC) is a pod within your Kubernetes cluster.
AGIC monitors the Kubernetes [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)
resources, and creates and applies App Gateway config based on these.

### Outline:
- [Prerequisites](#prerequisites)
- [Azure Resource Manager Authentication (ARM)](#azure-resource-manager-authentication)
    - Option 1: [Set up aad-pod-identity](#set-up-aad-pod-identity) and [Create Azure Identity on ARM](#create-azure-identity-on-arm)
    - Option 2: [Using a Service Principal](#using-a-service-principal)
- [Install Ingress Controller using Helm](#install-ingress-controller-as-a-helm-chart)
- [Multi-cluster / Shared App Gateway](#multi-cluster--shared-app-gateway): Install AGIC in an environment, where App Gateway is
shared between one or more AKS clusters and/or other Azure components.

### Prerequisites
This documents assumes you already have the following tools and infrastructure installed:
- [AKS](https://azure.microsoft.com/en-us/services/kubernetes-service/) with [Advanced Networking](https://docs.microsoft.com/en-us/azure/aks/configure-azure-cni) enabled
- [App Gateway v2](https://docs.microsoft.com/en-us/azure/application-gateway/create-zone-redundant) in the same virtual network as AKS
- [AAD Pod Identity](https://github.com/Azure/aad-pod-identity) installed on your AKS cluster
- [Cloud Shell](https://shell.azure.com/) is the Azure shell environment, which has `az` CLI, `kubectl`, and `helm` installed. These tools are required for the commands below.

Please __backup your App Gateway's configuration__ before installing AGIC:
  1. using [Azure Portal](https://portal.azure.com/) navigate to your `App Gateway` instance
  2. from `Export template` click `Download`

The zip file you downloaded will have JSON templates, bash, and PowerShell scripts you could use to restore App
Gateway should that become necessary

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

## Setup variables

```bash
applicationGatewayName="<application-gateway-name>"
applicationGatewayGroupName="<application-gateway-group-name>"
aksClusterName="<aks-cluster-name>"
aksClusterGroupName="<aks-cluster-group-name>"
agicIdentityName="agic-identity"
```

## Setup Azure Resource Manager Authentication (ARM)

### Using User assigned Identity

1. Create a User assigned Identity in AKS Agent Pool's resource group
    ```bash
    # Create identity in agent pool's resource group
    nodeResourceGroupName=$(az aks show -n $aksClusterName -g $aksClusterGroupName --query "nodeResourceGroup" -o tsv)
    nodeResourceGroupId=$(az group show -g $nodeResourceGroupName --query "id" -o tsv)

    az identity create -n $agicIdentityName -g $nodeResourceGroupName -l $location
    identityPrincipalId=$(az identity show -n $agicIdentityName -g $nodeResourceGroupName --query "principalId" -o tsv)
    identityResourceId=$(az identity show -n $agicIdentityName -g $nodeResourceGroupName --query "id" -o tsv)
    identityClientId=$(az identity show -n $agicIdentityName -g $nodeResourceGroupName --query "clientId" -o tsv)
    ```

1. Assign `Contributor` role to the Application Gateway resource group. If this step fails with "No matches in graph database for ...", then try again after few seconds.
    ```bash
    applicationGatewayGroupId=$(az group show -g $applicationGatewayGroupName -o tsv --query "id")

    az role assignment create \
            --role Contributor \
            --assignee "$identityPrincipalId" \
            --scope "$applicationGatewayGroupId"
    ```

### Using a Service Principal
It is also possible to provide AGIC access to ARM via a Kubernetes secret.

 1. Create an Active Directory Service Principal and encode with base64. The base64 encoding is required for the JSON blob to be saved to Kubernetes.
    ```bash
    az ad sp create-for-rbac --subscription <subscription-uuid> --sdk-auth > auth.json
    spBase64Encoded=$(cat auth.json | base64 -w0)
    spAppId=$(jq -r ".appId" auth.json)
    ```

1. Assign `Contributor` role to the Application Gateway resource group.
    ```bash
    applicationGatewayGroupId=$(az group show -g $applicationGatewayGroupName -o tsv --query "id")

    az role assignment create \
            --role Contributor \
            --assignee "$spAppId" \
            --scope "$applicationGatewayGroupId"
    ```

## Set up Application Gateway Ingress Controller

With the instructions in the previous section we created and configured a new AKS cluster and
an App Gateway. We are now ready to deploy a sample app and an ingress controller to our new
Kubernetes infrastructure.

<details>
<summary><strong>Install AAD Pod Identity (skip if already installed or using service principal for authentication)</strong></summary>
Azure Active Directory Pod Identity provides token-based access to [Azure Resource Manager (ARM)](https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-overview).

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
</details>

<details>
<summary><strong>Install Helm (skip if already installed)</strong></summary>
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
</details>

### Install Ingress Controller Helm Chart

1. Add the AGIC Helm repository:
    ```bash
    helm repo add application-gateway-kubernetes-ingress https://appgwingress.blob.core.windows.net/ingress-azure-helm-package/
    helm repo update
    ```

1. Install AGIC:
    ```bash
    applicationGatewayId=$(az network application-gateway show -n $applicationGatewayName -g $applicationGatewayGroupName -o tsv --query "id")

    helm install application-gateway-kubernetes-ingress/ingress-azure \
      --set appgw.applicationGatewayID=$applicationGatewayId \
      --set armAuth.type=aadPodIdentity \
      --set armAuth.identityResourceID=$identityResourceId \
      --set armAuth.identityClientID=$identityClientId \
      --version 0.10.0-rc5
    ```

Refer to the [tutorials](../tutorial.md) to understand how you can expose an AKS service over HTTP or HTTPS, to the internet, using an Azure App Gateway.



## Multi-cluster / Shared App Gateway
By default AGIC assumes full ownership of the App Gateway it is linked to. AGIC version 0.8.0 and later can
share a single App Gateway with other Azure components. For instance, we could use the same App Gateway for an app
hosted on VMSS as well as an AKS cluster.

Please __backup your App Gateway's configuration__ before enabling this setting:
  1. using [Azure Portal](https://portal.azure.com/) navigate to your `App Gateway` instance
  2. from `Export template` click `Download`

The zip file you downloaded will have JSON templates, bash, and PowerShell scripts you could use to restore App Gateway

### Example Scenario
Let's look at an imaginary App Gateway, which manages traffic for 2 web sites:
  - `dev.contoso.com` - hosted on a new AKS, using App Gateway and AGIC
  - `prod.contoso.com` - hosted on an [Azure VMSS](https://azure.microsoft.com/en-us/services/virtual-machine-scale-sets/)

With default settings, AGIC assumes 100% ownership of the App Gateway it is pointed to. AGIC overwrites all of App
Gateway's configuration. If we were to manually create a listener for `prod.contoso.com` (on App Gateway), without
defining it in the Kubernetes Ingress, AGIC will delete the `prod.contoso.com` config within seconds.

To install AGIC and also serve `prod.contoso.com` from our VMSS machines, we must constrain AGIC to configuring
`dev.contoso.com` only. This is facilitated by instantiating the following
[CRD](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/):

```bash
cat <<EOF | kubectl apply -f -
apiVersion: "appgw.ingress.k8s.io/v1"
kind: AzureIngressProhibitedTarget
metadata:
  name: prod-contoso-com
spec:
  hostname: prod.contoso.com
EOF
```

The command above creates an `AzureIngressProhibitedTarget` object. This makes AGIC (version 0.8.0 and later) aware of the existence of
App Gateway config for `prod.contoso.com` and explicitly instructs it to avoid changing any configuration
related to that hostname.


### Enable with new AGIC installation
To limit AGIC (version 0.8.0 and later) to a subset of the App Gateway configuration modify the `helm-config.yaml` template.
Under the `appgw:` section, add `shared` key and set it to to `true`.

```yaml
appgw:
    subscriptionId: <subscriptionId>    # existing field
    resourceGroup: <resourceGroupName>  # existing field
    name: <applicationGatewayName>      # existing field
    shared: true                        # <<<<< Add this field to enable shared App Gateway >>>>>
```

Apply the Helm changes:
  1. Ensure the `AzureIngressProhibitedTarget` CRD is installed with:
      ```bash
      kubectl apply -f https://raw.githubusercontent.com/Azure/application-gateway-kubernetes-ingress/ae695ef9bd05c8b708cedf6ff545595d0b7022dc/crds/AzureIngressProhibitedTarget.yaml
      ```
  2. Update Helm:
      ```bash
      helm upgrade \
          --recreate-pods \
          -f helm-config.yaml \
          ingress-azure application-gateway-kubernetes-ingress/ingress-azure
      ```

As a result your AKS will have a new instance of `AzureIngressProhibitedTarget` called `prohibit-all-targets`:
```bash
kubectl get AzureIngressProhibitedTargets prohibit-all-targets -o yaml
```

The object `prohibit-all-targets`, as the name implies, prohibits AGIC from changing config for *any* host and path.
Helm install with `appgw.shared=true` will deploy AGIC, but will not make any changes to App Gateway.


### Broaden permissions
Since Helm with `appgw.shared=true` and the default `prohibit-all-targets` blocks AGIC from applying any config.

Broaden AGIC permissions with:
1. Create a new `AzureIngressProhibitedTarget` with your specific setup:
    ```bash
    cat <<EOF | kubectl apply -f -
    apiVersion: "appgw.ingress.k8s.io/v1"
    kind: AzureIngressProhibitedTarget
    metadata:
      name: your-custom-prohibitions
    spec:
      hostname: your.own-hostname.com
    EOF
    ```

2. Only after you have created your own custom prohibition, you can delete the default one, which is too broad:

    ```bash
    kubectl delete AzureIngressProhibitedTarget prohibit-all-targets
    ```

### Enable for an existing AGIC installation
Let's assume that we already have a working AKS, App Gateway, and configured AGIC in our cluster. We have an Ingress for
`prod.contosor.com` and are successfully serving traffic for it from AKS. We want to add `staging.contoso.com` to our
existing App Gateway, but need to host it on a [VM](https://azure.microsoft.com/en-us/services/virtual-machines/). We
are going to re-use the existing App Gateway and manually configure a listener and backend pools for
`staging.contoso.com`. But manually tweaking App Gateway config (via
[portal](https://portal.azure.com), [ARM APIs](https://docs.microsoft.com/en-us/rest/api/resources/) or
[Terraform](https://www.terraform.io/)) would conflict with AGIC's assumptions of full ownership. Shortly after we apply
changes, AGIC will overwrite or delete them.

We can prohibit AGIC from making changes to a subset of configuration.

1. Create an `AzureIngressProhibitedTarget` object:
    ```bash
    cat <<EOF | kubectl apply -f -
    apiVersion: "appgw.ingress.k8s.io/v1"
    kind: AzureIngressProhibitedTarget
    metadata:
      name: manually-configured-staging-environment
    spec:
      hostname: staging.contoso.com
    EOF
    ```

2. View the newly created object:
    ```bash
    kubectl get AzureIngressProhibitedTargets
    ```

3. Modify App Gateway config via portal - add listeners, routing rules, backends etc. The new object we created
(`manually-configured-staging-environment`) will prohibit AGIC from overwriting App Gateway configuration related to
`staging.contoso.com`.
