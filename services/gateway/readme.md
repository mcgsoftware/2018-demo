# Gateway

We have to modify the Istio Bookinfo demo gateway to expose our services to 
Istio ingress gateway. Note the new istio version has some changes. 

We need to download the bookinfo demo from istio's github. 

Then we get the gateway config from: 
bookinfo_home/networking/bookinfo-gateway.yaml

We make a copy and call it:
bookinfo-gateway-mcg.yaml

# Gateway issues
 had to modify the bookinfo gateway to get traffic from outside to flow through the
Ingress gateway. I added a few lines of yaml at the end of
/istio-master/samples/bookinfo/networking/bookinfo-gateway.yaml

The ugly part is I have to map a URL to a fixed service endpoint, I must use 'prefix'
url mapping to pass through to multiple URI in same service!  Note the URL has to match
the code. For Example:

    // setup router
	router := mux.NewRouter()
	router.HandleFunc("/gcpservice", HealthHandler)
	router.HandleFunc("/gcpservice/health", HealthHandler)
	router.HandleFunc("/gcpservice/caller", callServiceHandler)
	router.HandleFunc("/health", HealthHandler)
	router.HandleFunc("/caller", callServiceHandler)

and you call it like this:

http://192.168.64.19:31380/gcpservice         => executes HealthHandler
http://192.168.64.19:31380/gcpservice/health  => executes HealthHandler
http://192.168.64.19:31380/gcpservice/caller  => executes CallServiceHandler


```
$ cp bookinfo-gateway.yaml bookinfo-gateway-mcg.yaml
$ vi bookinfo-gateway-mcg.yaml
$ kubectl apply -f bookinfo-gateway-mcg.yaml    ==> apply my hack
```

// Run this to revert to original gateway
kubectl apply -f bookinfo-gateway.yaml

```
-- file: bookinfo-gateway-mcg.yaml ---
--- key section below ---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: bookinfo
spec:
  hosts:
  - "*"
  gateways:
  - bookinfo-gateway
  http:
  - match:
    - uri:
        exact: /productpage
    - uri:
        exact: /login
    - uri:
        exact: /logout
    - uri:
        prefix: /api/v1/products
    route:
    - destination:
        host: productpage
        port:
          number: 9080
  - match:
    - uri:
        prefix: /gcpservice
    route:
    - destination:
        host: hellogcp

```



