# Google Cloud Notes
These are notes on using Google Cloud for the demo. 

# Learning kubernetes

First, create a google project, download gcloud tools and create a cluster. Their quickstart documentation is the easiest way to do this: https://cloud.google.com/kubernetes-engine/docs/quickstart. Play around with a temporary cluster and deploy something to it. 

# Demo web tools

## View the webapp
```
$ export GATEWAY_URL=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
$ echo $GATEWAY_URL
```
Now point browser at: http://$GATEWAY_URL/productpage

## Istio Service Graph

You can view istio's service graph via: http://localhost:8088/dotviz

```
$ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=servicegraph -o jsonpath='{.items[0].metadata.name}') 8088:8088 &
```


## Prometheus

Show Prometheus via port-forward: http://localhost:9090/graph 
Try http_requests_total as sample metric. See some prometheus metrics samples here: https://istio.io/docs/tasks/telemetry/querying-metrics/#about-the-prometheus-add-on
```
$ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=prometheus -o jsonpath='{.items[0].metadata.name}') 9090:9090 &
```

## Grafana
Setup a port-forward to view grafana with, then point browser to: http://localhost:3000 
```
$ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 3000:3000 &
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

## Special kubectl for GKE environment only
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
```

# Initial Project install / setup
 
- Create a project in google cloud: 'royal-2018-demo'
  - Also setup your local shell gcloud to point at this new project: $ gcloud config set project royal-2018-demo
- Navigate to the project's Container Registry via google cloud console. Make sure it is enabled, or docker image pushes fail.
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
