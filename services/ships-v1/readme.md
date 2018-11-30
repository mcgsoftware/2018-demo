# Java Spring Boot version of Ship service

# Build

## Build Docker Images

- open terminal in IDE
```

mvn clean

// build java artifacts first
mvn package


// Set the version we want to deploy
export VER="1.0"

// build docker image on local docker repo
docker build -t gcr.io/royal-2018-demo/ship-v1:$VER .

docker push gcr.io/royal-2018-demo/ship-v1:$VER

```

## Deploy on GKE

1. Edit /kubernetes/ships-deploy.yaml file, changing image version to $VER

2. Run istioctl to inject an envoy proxy into the pod (for minikube only)

3. Verify deployment.

```
// For minikube, need to do this to inject the envoy proxy during deployment
istioctl kube-inject -f ./kubernetes/ships-deploy.yaml | kubectl apply -f -

// For Google GKE, it gets installed with automatic envoy proxy injector
kubectl apply -f ./kubernetes/ships-deploy.yaml

kubectl get pods

// Test it out, port 8070 is port from profile service (and code)
kubectl port-forward <pod_name> 8070:8070

// Point browser at: http://127.0.0.1:8070/royal/api/ships/AL

```