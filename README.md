# appcore

## Import cursor library to the service

```shell
$ go get github.com/gladilindv/appcore
```

```go
import "github.com/gladilindv/appcore"
```

## Example
see `example/main.go`


## Configuration
Minimal configuration file in folder `.cfg/k8s`
```
service:
  ports:
    http: 8080
    grpc: 8082
    debug: 8084

env:
  log_level: DEBUG
```

## run 
```shell
ENV_APP=dev go run example/main.go
```
