# Google Cloud Notes

These are notes on using Minikube for the demo. 

# Installation with Minikube

## Install base istio components

Do not install sidecar injector, it makes things more complicated for things like installing Kiali. 

```
// wipe existing stuff from minikube
minikube delete

// Follow the minikube startup directions below
minikube start...
kubectl config...

// Install istio. Download it to $istio_home
cd $istio_home

// Set path to use istio tools
vi ~/.bash_profile
==> export PATH=$PATH:/Users/131052/Brian/metrics/istio/istio-1.0.4/bin

// follow instructions from kubernetes quickstart (for istio without TLS between sidecars)
// https://istio.io/docs/setup/kubernetes/quick-start/
kubectl apply -f install/kubernetes/helm/istio/templates/crds.yaml
kubectl apply -f install/kubernetes/istio-demo.yaml

// verify install of stuff in istio-system namespace
// The get pods command will take a while until everything starts up to completed or running status.
kubectl get services -n istio-system
kubectl get pods -n istio-system

```

## Installing services

Do not install autmatic proxy injects. Use the istioctl command to decorate your yaml files:
```
istioctl kube-inject -f <your-app-spec>.yaml | kubectl apply -f -
```

Do not install BookInfo samples. 
Install my services next. This assumes the
service yaml files are pointing to legit docker images
per the docker image path in them. 


```
cd project_home/services

//
// install logger service
//

cd log
istioctl kube-inject -f ./kubernetes/logger-deploy.yaml | kubectl apply -f - 

// View logs and test it, get pod name from 'get pods' command. 
kubectl get pods
kubectl logs -l app=logger -c logger
kubectl port-forward logger-v1-5bc44b9b55-6vqwj 8090:8090  

// in postman run with POST. Put a JSON doc in 'body'
POST http://localhost:8090/royal/api/logger

// check it worked 
kubectl logs -l app=logger -c logger

//
// install rest of the services
//
cd project_home/service
istioctl kube-inject -f ./booking-v1/kubernetes/booking-deploy.yaml | kubectl apply -f -
istioctl kube-inject -f ./booking-v2/kubernetes/booking-deploy.yaml | kubectl apply -f -
istioctl kube-inject -f ./services/profile/kubernetes/profile-deploy.yaml  | kubectl apply -f -

```


# Minikub Start-up
```
// Start minikube
minikube start --memory=8192 --disk-size=30g --kubernetes-version=v1.10.0 --vm-driver=hyperkit 

// Config kubectl to look at your minikube
kubectl config use-context minikube

```

# Learning kubernetes

First, create a google project, download gcloud tools and create a cluster. Their quickstart documentation is the easiest way to do this: https://cloud.google.com/kubernetes-engine/docs/quickstart. Play around with a temporary cluster and deploy something to it. 

# Demo web tools

## View the webapp
```
export GATEWAY_URL=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

echo $GATEWAY_URL
```
Now point browser at: http://$GATEWAY_URL/productpage

## Istio Service Graph

You can view istio's service graph via: http://localhost:8088/dotviz

```
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=servicegraph -o jsonpath='{.items[0].metadata.name}') 8088:8088 &
```


## Prometheus

Show Prometheus via port-forward: http://localhost:9090/graph 
Try http_requests_total as sample metric. See some prometheus metrics samples here: https://istio.io/docs/tasks/telemetry/querying-metrics/#about-the-prometheus-add-on
```
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=prometheus -o jsonpath='{.items[0].metadata.name}') 9090:9090 &
```
Total count of all requests to the productpage service:
  istio_requests_total{destination_service="productpage.default.svc.cluster.local"}
  
Total count of all requests to v3 of the reviews service:
   istio_requests_total{destination_service="reviews.default.svc.cluster.local", destination_version="v3"}
   
This query returns the current total count of all requests to the v3 of the reviews service.
Rate of requests over the past 5 minutes to all instances of the productpage service:
  rate(istio_requests_total{destination_service=~"productpage.*", response_code="200"}[1m])


## Grafana
Setup a port-forward to view grafana with, then point browser to: http://localhost:3000 
```
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 3000:3000 &
```

Send some traffic to productpage for metrics purposes:
```
$ for i in {1..100}; do curl -o /dev/null -s -w "%{http_code}\n" http://${GATEWAY_URL}/productpage; done
```

## Jaeger

View jaeger tracing via: http://localhost:16686
```
$ kubectl port-forward -n istio-system $(kubectl get pod -n istio-system -l app=jaeger -o jsonpath='{.items[0].metadata.name}') 16686:16686 &
```

## Kiali

### Install

Follow Kubernetes install instructions here: https://www.kiali.io/gettingstarted/
Except the following:
- Extrace all the curl commands from instructions and download the 3 yaml file separatedly.
- Manually edit each yaml file making substitutions per the the intentions of instructions. 
- Apply the 3 modified yaml files to 'istio-system' namespace in kubernetes: apply -f fn.yaml -n istio-system

### Run Kiali

Run this in shell:

```
Goto google GKE console. Click into 'kiali' under "workloads".
In the workloads > Deployment Details view, find the "Ingress" endpoint. 
Point browser at endpoint:

http://35.244.168.207/

Log in to Kiali-UI as admin/admin.

```






# GKE kubectl and other useful commands


## Istio project kubectl and other commands
Show all the app deployments in default namespace:
```
$ kubectl get deployments -n default 
$ kubectl get deployments            => same as above, default namespace implied.
```
Save the gateway's current IP address in shell variable GATEWAY_URL and echo it:
```
$ export GATEWAY_URL=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
$ echo $GATEWAY_URL
```
Show all of of istio's deployments, they are in istio-system namespace:
```
$ kubectl get deployments,ing -n istio-system
```
Show istio's gateway:
```
$ kubectl get svc istio-ingressgateway -n istio-system
```

## Special gcloud and kubectl for GKE environment only
Display current kubectl's current cluster settings:
```
$ kubectl config current-context
```

Show gcloud's current settings (google project, zone, etc.)
```
$ gcloud config list
```

Set kubectl to point at a specific project and GKE cluster:
```
$ gcloud config set project [PROJECT_NAME]
$ gcloud container clusters get-credentials [CLUSTER_NAME]

// Get credentials to use new cluster
$ gcloud container clusters get-credentials royal-cluster --zone us-east4-c --project royal-2018-demo

// make sure kubectl pointing at the new cluster
kubectl config current-context
```

   
    
# Deploying services to GKE and Istio

## Deploy a service
Build docker image locally, then push it to google's cloud 'Container Registry'. Once pushed to container registry, use kubectl to deploy it to GKE container. For example:

```
// compile sources
$ go build rccldemo.com/service

// Build in local docker repo. 
$ docker build -t gcr.io/royal-2018-demo/profile-svc:1.0 .

// Push docker image to Google's Container Registry
$ docker push gcr.io/royal-2018-demo/profile-svc:1.0

// Install into GKE 
$ kubectl run profile-svc --image gcr.io/royal-2018-demo/profile-svc:1.0 --port 8082

// Verify it installed
$ kubectl get pods

// Create a kubernetes service to expose our service on port 80, mapping from service's port 8082:
$ kubectl expose deployment profile-svc --type LoadBalancer --port 80 --target-port 8082
 
// verify service creation worked. You should see external ip when it's ready to use
$ kubectl get services

// Hit the external ip in POSTMAN or browser
http://<external-ip>/health

```

## View service logs 
You can view logs from kubectl or StackDriver. For kubectl:
```
// Run get pods to get the pod name (e.g. profile-svc-7f4777597d-prn4g)
$ kubectl get pods

// View pod logs. We can look at istio's proxy or the deployment name
$ kubectl logs profile-svc-7f4777597d-prn4g profile-svc

$ kubectl logs profile-svc-7f4777597d-prn4g istio-proxy


```

## Trobleshooting deployment problems
Use kubectl get pods and kubectl describe pod for debugging deployments.
```
// Will show your pods and what state they are in
$ kubectl get pods   

// Describe pod to get more details on deployment problems
$ kubectl describe pod <pod_name_from_get_pods> 
```

### Removing old or bad microservice deployment
To remove bad microservice deployments use the commands:
```
// get list of current deployments
$ kubectl get deployments 

// remove deployment 'provile-svc' and it's pod
$ kubectl delete deployment provile-svc

// get service for deployment
$ kubectl get services

// remove service profile-svc-v1
$ kubectl remove service profile-svc-v1
```

### ImagePullBackOff unauthorized issue
If you see kubectl get pods give you a state of "ImagePullBackOff" and kubectl describe pod says "unauthorized - authentication required" it is because container registry API doesn't have right permissions. As a fast work-around, you can go to Container Registry in google console and make it 'public' anyone can read from it. 

# GKE Istio Install 

## Create a project

## Create a google project in google cloud

  create a project named 'royal-2018-demo' in cloud web console. 
  
  Also setup your local shell gcloud to point at this new project: 
  
  ```
  $ gcloud config set project royal-2018-demo
  
  // check gcloud settings are right (project, zone, etc.)
  $ gcloud config list
  ```
## Install Istio

### Base install of Istio on GKE

Follow instructions here https://istio.io/docs/setup/kubernetes/quick-start-gke-dm/
use the Google Deployment Mangager. I spent hours trying to install isito manually and it didn't work on GKE.

### Install bookinfo on local machine

### Download istio from internet into local machine for use

Install Istio from github onto local machine, it's needed for hacking the gateway config. 
It also has istioctl for use later. IMPORTANT: Google deployment manager installs the automatic
envoy injection so you don't need to wrapper microservice deployments with istioctl calls to
decorate the yaml files. 

```
cd ~/Brian/metrics
mkdir istio
cd istio/
curl -L https://git.io/getLatestIstio | sh -
cd istio-1.0.4/
export PATH=$PWD/bin:$PATH. 

// IMPORTANT: enter this path into .bash_profile too!
```


#  =====
# TODO: MOVE THIS TO MINIKUBE page - Initial Project install / setup
#

  
## Setup container registry
  Navigate to the project's Container Registry via google cloud console. Make sure it is enabled, or docker image pushes fail.
  
## Create GKE cluster (do this before adding istio)

```
// check gcloud settings are right (project, zone, etc.)
gcloud config list

// Create a new cluster named ‘royal-cluster’
// IMPORTANT: must be 4 standard nodes minimum or the istio install won't work!!
gcloud container clusters create royal-cluster \
    --machine-type=n1-standard-2 \
    --num-nodes=4 \
    --no-enable-legacy-authorization
    
// Get credentials to use new cluster
gcloud container clusters get-credentials royal-cluster --zone us-east4-c --project royal-2018-demo

// make sure kubectl pointing at the new cluster
kubectl config current-context

kubectl create clusterrolebinding root-cluster-admin-binding \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"


```

### Try out GKE cluster

Install a simple service, see it work.
```

// run a service, assumes you built the app and pushed it to google container registry first. listens on port 8082
kubectl run profile-svc --image gcr.io/royal-2018-demo/profile-svc:1.0 --port 8082

// expose service to external traffic
kubectl expose deployment profile-svc --type LoadBalancer --port 80 --target-port 8082

// check what external ip it used (e.g. 35.245.49.124 )
kubectl get services
 
http://35.245.49.124/health

// clean-up
kubectl delete deployment profile-svc
kubectl delete service profile-svc

```





## install istio


### Download istio from internet into local machine for use
```
cd ~/Brian/metrics
mkdir istio
cd istio/
curl -L https://git.io/getLatestIstio | sh -
cd istio-1.0.4/
export PATH=$PWD/bin:$PATH. 

// IMPORTANT: enter this path into .bash_profile too!
```

### Install istio on GKE cluster
Installs base istio without any apps on it. 

```
kubectl apply -f install/kubernetes/helm/istio/templates/crds.yaml

kubectl apply -f install/kubernetes/istio-demo.yaml

// Verify installation, check stuff in istio-system namespace
kubectl get svc -n istio-system
kubectl get pods -n istio-system
```


### Check it all out in google cloud console
Go to google cloud console click on cluster workload and services to make sure everything is Green!



XXX
- Setup istio bookinfo sample for GKE following _all_ the steps: https://istio.io/docs/setup/kubernetes/quick-start-gke-dm/
  - Make sure you do all the IAM stuff too or you will be sorry later!
  - For the 'Launch Deployment Manager' section, when in the Istio GKE Deployment manager webapp do:
    - Change the deployment name to 'royal-cluster'
    - Change the cluster name from 'istio-cluster' to 'royal-cluster'
    - Change the Zone to 'us-east4-c'
    - Leave everything else same in web form and click the 'deploy' button.
  - In the 'Bootstrap gcloud' section of instructions, do this instead:
    - $ gcloud config set project royal-2018-demo
    - $ gcloud container clusters list => Verify you see 'royal-cluster' at 'us-east4-c' location.
    - $ gcloud container clusters get-credentials royal-cluster --zone=us-east4-c  => get the credentials for cluster
  - Follow the 'Verify Istallation' steps to make sure it's configured properly. Also, do the following:
    - $ kubectl get pods  => should show bookinfo pods as 'Running' (not in other state)
 
