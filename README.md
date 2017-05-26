# go-vapidkeys

A small package in Go to generate VAPID public and private keys required for web push.

## Installation

```bash
go get github.com/anaskhan96/go-vapidkeys
```

## Usage

```go
package main

import (
	"github.com/anaskhan96/go-vapidkeys"
)

func main() {
	privKey, pubKey, _ := vapidkeys.Generate()
	println(privKey)
	println(pubKey)
}

```
