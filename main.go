package main

import (
	"github.com/koodinikkarit/pekka/pekka"
)

func main() {
	server := pekka.CreatePekkaServer(
		"jaska",
		"asdf321",
		"localhost",
		"3306",
		"pekka",
		"5054",
		"3111",
	)

	server.Start()
}
