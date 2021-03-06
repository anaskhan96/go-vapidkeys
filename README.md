# go-vapidkeys
[![Build Status](https://travis-ci.org/anaskhan96/go-vapidkeys.svg?branch=master)](https://travis-ci.org/anaskhan96/go-vapidkeys)
[![GoDoc](https://godoc.org/github.com/anaskhan96/go-vapidkeys?status.svg)](https://godoc.org/github.com/anaskhan96/go-vapidkeys)
[![Go Report Card](https://goreportcard.com/badge/github.com/anaskhan96/go-vapidkeys)](https://goreportcard.com/report/github.com/anaskhan96/go-vapidkeys)

A small package in Go to generate VAPID public and private keys required for sending web push notifications (HTTP/2 Server Push).

## Installation

```bash
go get github.com/anaskhan96/go-vapidkeys
```

Run `go test` in the package's directory to run tests.

## Usage

```go
package main

import (
	"fmt"
	"github.com/anaskhan96/go-vapidkeys"
	"log"
)

func main() {
	privateKey, publicKey, err := vapidkeys.Generate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Vapid Private Key:", privateKey)
	fmt.Println("Vapid Public Key:", publicKey)
}

```
