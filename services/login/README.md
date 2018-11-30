# Login Service
This is a Javalin com.rccl.service.

## Usage

POST http://<external-ip>/royal/api/login with HTML form username=sri and password=brian

It returns json object with message of success or failure.


GET http://<external-ip>/royal/api/login/health  returns health uptime

## Build Instructions

```
cd project_home

mvn clean

mvn package

// try it out locally (needs to use 'shaded' jar file or it won't find javalin.io dependency)
java -cp target/login-1.0-SNAPSHOT-shaded.jar com.rccl.service.HelloWorld

// Export to local docker repo
mvn dockerfile:build


docker push gcr.io/royal-2018-demo/login:1.0

```

## Run in docker locally

```
mvn package

docker run -p 7000:7000 gcr.io/royal-2018-demo/login:1.0

```

## Deploy to GKE and Istio
Note: GKE Istio has automatic envoy proxy injection, so we don't need to use istioctl to decorate our deployment
yaml file. 

## Deploy on GKE

1. Edit /kubernetes/ships-deploy.yaml file, changing image version to $VER

2. Run istioctl to inject an envoy proxy into the pod (for minikube only)

3. Verify deployment.

```
// For minikube, need to do this to inject the envoy proxy during deployment
istioctl kube-inject -f ./kubernetes/login-deploy.yaml | kubectl apply -f -

// First, remove old service
kubectl delete deployments login-v1
kubectl delete services login


// For Google GKE, it gets installed with automatic envoy proxy injector
kubectl apply -f ./kubernetes/login-deploy.yaml

kubectl get pods

// Test it out, port 8070 is port from profile service (and code)
kubectl port-forward <pod_name> 7000:7000

// Point browser at: http://127.0.0.1:7000/royal/api/login

```

### Running in Istio
The gateway has to be adjusted to include login service or you can't see it from port 80 outside the cluster.
(see gateway for info on this).

Once working, you should be able to see it using the gateway's external ip:

http://<external_ip>/royal/api/login
