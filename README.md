# argocd-playground

ArgoCD playground and POC

### POC plans

- create the appropriate cluster using flex
- deploy the argocd to a k8s cluster
- able to access the argocd ui using https: how to manage with ingress or istio (request a proper cert for this)
- create a sample apps and play around it
- learn about the rbac
- connect to SD using cd events
- multi cluster of argocd

#### Setup ingress

```
helm upgrade --install ingress-nginx ingress-nginx repo https://kubernetes.github.io/ingress-nginx --namespace ingress-nginx --create-namespace
```

##### Problems encountered

- need to change security group of the load balancer to allow port 80
- pods need to be in the same nodes with ingress controller or if not need to allow connectivity between node groups https://github.com/terraform-aws-modules/terraform-aws-eks/blob/master/docs/network_connectivity.md#security-groups
- apparently need to change the ACL on where nodes are
- need to use self-signed cert for the argocd or use lets encrypt

### Create the self-signed cert argocd example

```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -out self-signed-tls-argocd.crt -keyout self-signed-tls-argocd.key \
    -subj "/CN=argocd.example.com" \
    -reqexts SAN \
    -extensions SAN \
    -config <(cat /etc/ssl/openssl.cnf \
        <(printf "[SAN]\nsubjectAltName=DNS:argocd.example.com,DNS:*.argocd.example.com"))

kubectl create secret tls argocd-server-tls --key=self-signed-tls-argocd.key --cert=self-signed-tls-argocd.crt -n argocd
```

### Manage argocd apps

- All resources, including Application and AppProject specs, have to be installed in the Argo CD namespace (by default argocd).

#### References

- [Argo CD on Rancher Desktop as a local GitOps K8s lab](https://itnext.io/argo-cd-rancher-desktop-for-a-local-gitops-lab-8d044089f50a)
- [Kubernetes Ingress — HTTP to HTTPS with Self-Signed Certificates](https://medium.com/@muppedaanvesh/%EF%B8%8F-kubernetes-ingress-transitioning-to-https-with-self-signed-certificates-0c7ab0231e76)
- [pubsub-example](https://github.com/olivere/pubsub-example)
- [Sample apps](https://github.com/gokul0815/argocd/tree/master)
- apps of apps and applicationsets
  - https://medium.com/dzerolabs/turbocharge-argocd-with-app-of-apps-pattern-and-kustomized-helm-ea4993190e7c
  - https://medium.com/@andersondario/argocd-app-of-apps-a-gitops-approach-52b17a919a66

#### Harness and Argocd

- What is the use of delegate in this case?
- Gitops agent and argocd deployment
- How to automate the gitops agent and delegate deployment?
- We can use delegate to deploy gitops agent and argocd. But what about deployment of delegate itself? I think the first one is the manual deployment just like how we bootstrap the infra for other ci/cd tools like screwdriver.
- What about github actions and gitops model working together?

- BYOA argocd and let that be managed minimum level by Harness
- centralized gitops agent with argocd and let devs managed their own apps
- can argocd trigger by another apps? is the testing actually part of blue green or canary?
- do we need to have a UI to show all workflows together?
- argo events and argo workflows?
- argo has trigger to github and maybe this can be used as functional tests
  https://youtu.be/ag8v0Jl9n8g

#### Details ArgoCD

[Example](https://medium.com/dzerolabs/turbocharge-argocd-with-app-of-apps-pattern-and-kustomized-helm-ea4993190e7c)

#### ArgoCD Implementation

- deploy argocd but create it as apps of apps including the ingress implementation and rbac: or just install argocd using Harness
  - create connection to my personal k8s: not sure how for now
    - need to know how fleks connect to my k8s cluster
    - create delegate and connect using it
  - and then run a script to do deployment https://archive.eksworkshop.com/intermediate/290_argocd/install/
    - install via helm https://www.arthurkoziel.com/setting-up-argocd-with-helm/
  - commit back to my git repo
- what about doing it in Harness: need delegate and gitops agent
- then try out some deployment like SD instances in there

#### Implement the LambdaTest conversion module

- Ask team to identify the apps that need to migrate
- create the desired capabilities and run few test cases
- submit PR to them and enable them
- also please run tests for Screwdriver UI in LambdaTest
- deploy sample apps in EKS and allow the connections from that to the LT in AWS

### POC Github Actions and ArgoCD

- CI in Github actions
- CD in argo
  - need to build and operationalize argo
  - request the appropriate digicert for this
