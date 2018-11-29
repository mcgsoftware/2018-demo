# Kubernetes Notes
These are generic kubernetes notes, useful for minikube or google GKE.

# Gotchas

## Issue: Cached docker images
For some reason minikube (google gke too?) caches docker images. 
Even if you build a new image and push it to dockerhub, it will reuse the old image
in kubernetes unless you give it a new version tag. The side effect is that
code changes redeployed do not take effect. The work around is to create
a new docker version tag every time you deploy a new version of a service. 

## Issue: Gateway / Ingress problems

todo write it here.

