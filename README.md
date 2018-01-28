# Go env

Small lib for configuring your go code according to current environment `RUN_ENV`.

## Installation
```go
go get github.com/thedanielforum/run-env
```

## How to use
To get the current env you should call `env.InitEnv()` when your application is starting.
It is important to call `env.InitEnv()` or `env.SetEnv("the env you want")` before calling any og the "IsEnv" functions.

### Example
```go
package main

import (
    "fmt"
    "github.com/thedanielforum/env"
)

func init() {
    env.InitEnv()
}

func main() {
    fmt.Println(env.GetEnv().String())
}
```
