package main

import (
	"flag"
	"fmt"

	"github.com/Drylozu/FormaLaPalabra/pkg/server"
)

var help = flag.Bool("help", false, "Muestra el mensaje de ayuda")
var public = flag.String("public", "./public/", "Especifica la carpeta public/")
var secret = flag.String("secret", "abcdefghijklmnñopqrstuvwxyzABCDEFGHIJKLMNÑOPQRSTUVWXYZ", "Especifica la contraseña del servidor")
var address = flag.String("address", "0.0.0.0:3000", "Especifica la dirección de escucha")

func main() {
	flag.Parse()

	if *help {
		fmt.Println("Formá la palabra\n\nOpciones:")
		flag.PrintDefaults()
		return
	}

	server.CreateAndListen(*address, &server.Options{
		Public:    *public,
		SecretKey: *secret,
	})
}
