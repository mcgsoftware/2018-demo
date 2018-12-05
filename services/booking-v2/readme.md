# Booking Service

## General
Profile service is written in Go and uses minimal dependencies.

## Istio MySQL egress issue
By default, Istio will block calls to mysql hosted elsewhere.
To fix this, follow the steps in mysql readme to punch a hole in
cluster for mysql access. 

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
export VER="1.5"

go install rccldemo.com/service

docker build -t gcr.io/royal-2018-demo/booking-v2:$VER .

docker push gcr.io/royal-2018-demo/booking-v2:$VER

```

### Install successful build into Istio
This installs service into GKE and Istio.

0. cd $GOPATH

1. Edit /kubernetes/booking-deploy.yaml file, changing image version to $VER

2. Run istioctl to inject an envoy proxy into the pod (for minikube only)

3. Verify deployment. 


```
// For minikube, need to do this to inject the envoy proxy during deployment
istioctl kube-inject -f ./kubernetes/booking-deploy.yaml | kubectl apply -f -

// Remove old deployment
kubectl delete deployment booking-v2 
kubectl delete service booking

// For Google GKE, it gets installed with automatic envoy proxy injector
kubectl apply -f ./kubernetes/booking-deploy.yaml




kubectl get pods -l app=booking
kubectl logs <pod-name> -c booking

// Test it out, port 8070 is port from profile service (and code)
kubectl port-forward <pod_name> 8070:8070  

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
 $ go get -t github.com/go-sql-driver/mysql
 $ go get -t github.com/pkg/errors

```

# Deployment Instructions for GKE

# Project Structure
- /src => go code
- /kubernetes => K8 and Istio config files
