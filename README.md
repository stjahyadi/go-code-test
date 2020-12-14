# Go Code Test

### Run jwt-client

Open new cli and setup environment variable to store TOKEN_SECRET

```sh
set TOKEN_SECRET=supersecretphrase
```

In the same cli, go to [workspace]/go-code-test/jwt-client/ and run jwt-client to get JWT token from console

```sh
go run jwt_client.go
```

### Run server API

Open new cli and setup environment variable to store TOKEN_SECRET

```sh
set TOKEN_SECRET=supersecretphrase
```

In the same cli, go to [workspace]/go-code-test/ and run *go run* to start up the server
```sh
go run main.go
```

### Access API

Open cli to execute curl command
```sh
curl -X GET "http://localhost:8081/api/locations" -H  "accept: application/json" -H  "Token: <Access Token>"
```

### Swagger API documentation

Open cli and go to [workspace]/go-code-test/ and run *go-swagger* command to start up swagger docs

```sh
go-swagger serve -F=swagger swagger.yaml -p=8082
```

Open browser and open http://localhost:8082/docs

### Unit Testing

Open cli and go to [workspace]/go-code-test/ and run *go test* command

```sh
go test
```

### [Additional] Build and Run server API as docker 

Open cli and go to [workspace]/go-code-test/ and execute docker command as below

```sh
docker build -t gomaps-app .
```

Once the image is successfully built, run docker 

```sh
docker run -e TOKEN_SECRET=supersecret --publish 8080:8081 --detach --name gomaps-app gomaps-app:latest
```

Access it via [ip_address]:8080
