# go-vapidkeys

This is a quite small and trivial (that's putting it *nicely*) package, so I'll give a brief backstory:

> Wanting to implement web push from the server side in Go, I found two working libraries [push-encryption-go](https://github.com/GoogleChrome/push-encryption-go) and [webpush-go](https://github.com/SherClockHolmes/webpush-go), but none of them provided the luxury of generating VAPID keys, like [this](https://github.com/web-push-libs/web-push) npm package.
> To implement this functionality, I had to look up stuff like elliptic curves, the ECDH algorithm, and Base 64 Encoding with its types, including padding and non-padding.
> To avoid all of this frustration for the next gopher to come along, I decided to write this small but helpful package.

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
