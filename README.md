# Go env

Small lib for configuring your go code according to current environment `RUN_ENV`.

[![CircleCI](https://circleci.com/gh/thedanielforum/env.svg?style=svg)](https://circleci.com/gh/thedanielforum/env) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/0c235721a7c34876be2b7c3b77444173)](https://www.codacy.com/app/thedanielforum/env?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=thedanielforum/env&amp;utm_campaign=Badge_Grade)

## Installation
```go
go get github.com/thedanielforum/env
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
