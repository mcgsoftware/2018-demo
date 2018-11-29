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

## View grafana metrics
Setup a tunnel to view grafana with, then point browser to: http://localhost:3000 
```
$ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 3000:3000 &
```

Send some traffic to productpage for metrics purposes:
```
$ for i in {1..100}; do curl -o /dev/null -s -w "%{http_code}\n" http://${GATEWAY_URL}/productpage; done
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
Set kubectl to point at a specific project and GKE cluster:
```
$ gcloud config set project [PROJECT_NAME]
$ gcloud container clusters get-credentials [CLUSTER_NAME]
```

# Initial Project install / setup

- Create a project in google cloud: 'royal-2018-demo'
  - Also setup your local shell gcloud to point at this new project: $ gcloud config set project royal-2018-demo
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
   
    



