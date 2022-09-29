package assets

import (
	"log"
	"os"
)

func Logo() []byte {
	data, err := os.ReadFile("./assets/images/SimulationBrowser.ico")
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	return data
}
