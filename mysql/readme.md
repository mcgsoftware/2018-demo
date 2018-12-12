# MySQL Notes

## Configure Istio for MySQL use
The database connection to mysql won't work unless there is a 'serviceentry' created in Istio. The egress will fail.
To install the service entry, do the following steps:

1. Get the ip address for the mysql host and update the yaml file with it. 
```
   $ host sql9.freemysqlhosting.net

   // Dont forget to edit yaml file with this ip address!
```
  
2. Apply the mysql-network yaml file
```
   $ kubectl apply -f mysql-network.yaml
   
   // verify it worked
   $ kubectl get serviceentry mysql-external -o yaml
```

# Using mysql command line
Using mysql shell.
```
// login and enter the passwd
$ mysql -u sql9267914 -p -h sql9.freemysqlhosting.net

mysql> use sql9267914;
mysql> show tables;
...
```
# Using MySQL in GCP

```
// login to mysql from cloud shell
gcloud sql connect royal-mysql --user=root  


```

# Database Setup
Apply the schema.sql and data.sql files to the mysql database. 

