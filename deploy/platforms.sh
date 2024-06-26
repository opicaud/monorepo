kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
argoPass=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
echo $argoPass
brew install argocd
kubectl apply -f argocd-helm-creds.yaml
bazel run //apps/security:security.install
#Workaround, try to apply some priority to improve
sleep 60
bazel run //apps/observability-operator:observability_operator.install