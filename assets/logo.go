package assets

import (
	"log"
	"os"
)

func Logo() []byte {
	data, err := os.ReadFile("./assets/images/Aix.ico")
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	return data
}
