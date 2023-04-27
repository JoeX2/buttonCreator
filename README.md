# buttonCreator
This creates buttons ;-)

## Link to precentation:
[Kubernetes Mark II](https://docs.google.com/presentation/d/1phA2GSS3Kg7KohFzkWGx83hhvtwQJIlLqdLrFqsgv-I/edit?usp=sharing)

## Installing

What I wrote to set up this

```bash
# Install k3s
curl -sfL https://get.k3s.io | sh -

# Install Helm and connect to k3s
apt update
apt install snapd
sudo snap install helm --classic
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml

# Install cert-manager and ArgoCD so it can install everything else
kubectl create namespace argocd
kubectl -n argocd create secret generic ssh-key-secret --from-file=github-privatekey=/path/to/secret/argocd_id_rsa
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.8.2/cert-manager.crds.yaml

helm repo add argo-cd https://argoproj.github.io/argo-helm
helm dep update argocd-helm/charts/argo-cd/

helm install argo-cd argocd-helm/charts/argo-cd/ -n argocd -f ./argocd-helm/charts/argo-cd/values.yaml
helm template argocd-helm/apps/ | kubectl -n argocd apply -f -



