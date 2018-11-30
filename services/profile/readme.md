# Profile Service

## General
Profile service is written in Go and uses minimal dependencies.

## Istio Gateway Issues

Istio's bookinfo demo specifies a path matching for gateway ingress.
See ../gateway/readme for more info. You must run the 
gateway hack there in order to make this service visible to 
callers outside the cluster.

You also must bind the server to "http://127.0.0.1" only
specifically, or it will never work in Istio!! It took days to figure
that out.


## Build and push to GKE



### Build docker image from $GOPATH and upload to google repository:
```
// Set the version we want to deploy
export VER="1.3"

go install rccldemo.com/service

docker build -t gcr.io/royal-2018-demo/profile-svc:$VER .

docker push gcr.io/royal-2018-demo/profile-svc:$VER

```

### Install successful build into Istio
This installs service into GKE and Istio.

0. cd $GOPATH

1. Edit /kubernetes/profile-deploy.yaml file, changing image version to $VER

2. Run istioctl to inject an envoy proxy into the pod (for minikube only)

3. Verify deployment. 


```
// For minikube, need to do this to inject the envoy proxy during deployment
istioctl kube-inject -f ./kubernetes/profile-deploy.yaml | kubectl apply -f -

// For Google GKE, it gets installed with automatic envoy proxy injector
kubectl apply -f ./kubernetes/profile-deploy.yaml

kubectl get pods

// Test it out, port 8082 is port from profile service (and code)
kubectl port-forward <pod_name> 8082:8082  

// Point browser at: 27.0.0.1:8082/royal/api/profile/233

```

### Test on a kubernetes deployment without Istio
This will deploy into GKE on port 8080 so we can play with service without
Istio stuff. 

```
kubectl run profile-svc --image gcr.io/royal-2018-demo/profile-svc:1.0 --port 8082

kubectl get pods

kubectl expose deployment profile-svc --type LoadBalancer --port 80 --target-port 8082
 
kubectl get services

// Hit the external ip in POSTMAN or browser
http://<external-ip>/health

```

## Removing bad or old service deployments from GKE

Don't remove pods, remove the service's deployments and services.
```
// Get deployment names (e.g. profile-svc)
$ kubectl get deployments

// Remove deployment
$ kubectl delete deployment profile-svc

// Get service names (e.g. provile-svc-v1)
$ kubectl get services

// Remove service
$ kubectl delete service provile-svc-v1
```

## Removing docker image from google's repo
Goto google cloud console's container registry for the project 
to make sure it made it. You can delete it from google docker repo with: 
```
$ gcloud container images delete gcr.io/royal-2018-demo/profile-svc:$VER --force-delete-tags
```

## Local build notes

In shell of IDE, setup GOPATH, etc. or it won't work.
```
$ export GOPATH=`pwd`
$ echo $GOPATH

// do this just once
$ gcloud auth configure-docker
```

Build artifacts. Install creates binary '/bin/service'
```
$ go build rccldemo.com/service
$ go install rccldemo.com/service

```

## Dependencies
```
 $ go get -t github.com/google/uuid
 $ go get -t github.com/gorilla/mux

```

# Deployment Instructions for GKE

# Project Structure
- /src => go code
- /kubernetes => K8 and Istio config files
