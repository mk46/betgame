# Bet game project

## To start mongodb from docker use below command:

```
docker run -d -p 27017:27017 --name  -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo
```

27017 port will be map localhost, which will help to run golang app.
