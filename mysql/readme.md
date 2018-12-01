# MySQL Notes

## Configure Istio for MySQL use
The database connection to mysql won't work unless there is a 'serviceentry' created in Istio. The egress will fail.
To install the service entry, do the following steps:

1. Get the ip address for the mysql host
```
   $ host sql9.freemysqlhosting.net
```
  
2. Apply the mysql-network yaml file
```
   $ apply -f mysql-network.yaml
   
   // verify it worked
   kubectl get serviceentry mysql-external -o yaml
```


# Using MySQL in GCP

```
// login to mysql from cloud shell
gcloud sql connect royal-mysql --user=root  


```

## Todo
- add shipinfo table
  - refactor reservations table to use fkey instead of enum
  - refactor service to pull ship name via join of tables

