# go-vapidkeys
[![Build Status](https://travis-ci.org/anaskhan96/go-vapidkeys.svg?branch=master)](https://travis-ci.org/anaskhan96/go-vapidkeys)

A small package in Go to generate VAPID public and private keys required for web push.

## Installation

```bash
go get github.com/anaskhan96/go-vapidkeys
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/anaskhan96/go-vapidkeys"
)

func main() {
	privateKey, publicKey, _ := vapidkeys.Generate()
	fmt.Println("Vapid Private Key:", privateKey)
	fmt.Println("Vapid Public Key:", publicKey)
}

```
