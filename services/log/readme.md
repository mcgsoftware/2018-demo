# Logger Service

## General
Takes input posted to it and logs to stdout.

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

$ export GOPATH=`pwd`

// Set the version we want to deploy
export VER="1.2"

go build rccldemo.com/service

docker build -t gcr.io/royal-2018-demo/logger:1.0 .

docker push gcr.io/royal-2018-demo/logger:1.0

```

### Install successful build into Istio
This installs service into GKE and Istio.

0. cd $GOPATH

1. Edit /kubernetes/logger-deploy.yaml file, changing image version to $VER

2. Run istioctl to inject an envoy proxy into the pod (for minikube only)

3. Verify deployment. 


```
// For minikube, need to do this to inject the envoy proxy during deployment
istioctl kube-inject -f ./kubernetes/booking-deploy.yaml | kubectl apply -f -

// Remove old deployment
kubectl delete deployment logger-v1 
kubectl delete service logger

// For Google GKE, it gets installed with automatic envoy proxy injector
kubectl apply -f ./kubernetes/logger-deploy.yaml

kubectl get pods

// Test it out, port 8070 is port from profile service (and code)
kubectl port-forward <pod_name> 8090:8090  

// Point browser at: http://127.0.0.1:8070/royal/api/bookings/vdsId

```

### Test on a kubernetes deployment without Istio
This will deploy into GKE on port 8070 so we can play with service without
Istio stuff. 

```
kubectl run booking --image gcr.io/royal-2018-demo/booking:1.0 --port 8070

kubectl get pods

kubectl expose deployment booking --type LoadBalancer --port 80 --target-port 8070
 
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
