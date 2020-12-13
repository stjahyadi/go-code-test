# GoMaps2020

### Run jwt-client

Open cli and setup environment variable to store TOKEN_SECRET

```sh
set TOKEN_SECRET=supersecretphrase
```

In the same cli, go to [workspace]/jwt-client/ and run jwt-client to get JWT token from console

```sh
go run jwt_client.go
```

### Run server API

Open cli and setup environment variable to store TOKEN_SECRET

```sh
set TOKEN_SECRET=supersecretphrase
```

In the same cli, go to [workspace] and run main.go to start up the server
```sh
go run main.go
```

### Access API

Open cli to execute curl command
```sh
curl -X GET "http://localhost:8081/api/locations" -H  "accept: application/json" -H  "Token: <Access Token>"
```

### Swagger API documentation

Open cli and go to [workspace] and run go-swagger command to start up swagger docs

```sh
go-swagger serve -F=swagger swagger.yaml -p=8082
```

Open browser and open http://localhost:8082/docs

### Other notes

None.