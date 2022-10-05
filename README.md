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

```yaml
service:
  ports:
    http: 8080
    grpc: 8082
    debug: 8084

env:
  log_level: DEBUG
```

## Run

```shell
$ APP_ENV=dev go run example/main.go
```
