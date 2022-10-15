package main

import (
	"flag"
	"fmt"

	"github.com/Drylozu/FormaLaPalabra/pkg/server"
)

var help = flag.Bool("help", false, "Shows the help message")
var public = flag.String("public", "./public/", "Specifies the public/ directory")
var secret = flag.String("secret", "abcdefghijklmnñopqrstuvwxyzABCDEFGHIJKLMNÑOPQRSTUVWXYZ", "Specifies the secret key to auth")
var address = flag.String("address", "127.0.0.1:3000", "Specifies the address to listen")

func main() {
	flag.Parse()

	if *help {
		fmt.Println("Formá la palabra\n\nOptions:")
		flag.PrintDefaults()
		return
	}

	server.CreateAndListen(*address, &server.Options{
		Public:    *public,
		SecretKey: *secret,
	})
}
