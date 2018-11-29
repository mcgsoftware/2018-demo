# Google Cloud Notes
These are notes on using Google Cloud for the demo. 

# Learning kubernetes

First, create a google project, download gcloud tools and create a cluster. Their quickstart documentation is the easiest way to do this: https://cloud.google.com/kubernetes-engine/docs/quickstart. Play around with a temporary cluster and deploy something to it. 

# GKE kubectl commands

## General kubectl for GKE
Display current kubectl's current cluster settings:
```
$ kubectl config current-context
```
Set kubectl to point at a specific project and GKE cluster:
```
$ gcloud config set project [PROJECT_NAME]
$ gcloud container clusters get-credentials [CLUSTER_NAME]
```
## Istio project kubectl commands
Show all the app deployments in default namespace:
```
$ kubectl get deployments -n default 
$ kubectl get deployments            => same as above, default namespace implied.
```

Show all of of istio's deployments, they are in istio-system namespace:
```
$ kubectl get deployments,ing -n istio-system
```
Show istio's gateway:
```
$ kubectl get svc istio-ingressgateway -n istio-system
```

# Project setup

- Create a project in google cloud: 'royal-2018-demo'
  - Also setup your local shell gcloud to point at this new project: $ gcloud config set project royal-2018-demo
- Setup istio bookinfo sample for GKE following _all_ the steps: https://istio.io/docs/setup/kubernetes/quick-start-gke-dm/
  - Make sure you do all the IAM stuff too or you will be sorry later!
  - For the 'Launch Deployment Manager' section, when in the Istio GKE Deployment manager webapp do:
    - Change the deployment name to 'royal-cluster-1'
    - Change the cluster name from 'istio-cluster' to 'royal-cluster'
    - Change the Zone to 'us-east4-c'
    - Leave everything else same in web form and click the 'deploy' button.
  - In the 'Bootstrap gcloud' section of instructions, do this instead:
    - $ gcloud config set project royal-2018-demo
    - $ gcloud container clusters list => Verify you see 'royal-cluster' at 'us-east4-c' location.
    - $ gcloud container clusters get-credentials royal-cluster --zone=us-east4-c  => get the credentials for cluster
  - Follow the 'Verify Istallation' steps to make sure it's configured properly.
   
    



