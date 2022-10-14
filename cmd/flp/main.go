package main

import (
	"flag"

	"github.com/Drylozu/FormaLaPalabra/pkg/server"
)

var public = flag.String("public", "./public/", "Specifies the public/ directory")
var secret = flag.String("secret", "abcdefghijklmnñopqrstuvwxyzABCDEFGHIJKLMNÑOPQRSTUVWXYZ", "Specifies the public/ directory")
var address = flag.String("address", "127.0.0.1:3000", "Specifies the address to listen")

func main() {
	flag.Parse()

	server.CreateAndListen(*address, &server.Options{
		Public:    *public,
		SecretKey: *secret,
	})
}
