# go-service

Go simple service control based on kardianos service package.

## Install

```console
go get github.com/dainiauskas/go-service
```

## Use

```go
import "github.com/dainiauskas/go-service"
```

This gives you access to the `go-service` package.

### Example
```go
srv := service.New("Name", "Display", "Description")

// Set function to run
srv.SetCb(func() {
  fmt.Println("Hello service")
})
```
