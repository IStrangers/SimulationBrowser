package assets

import (
	"log"
	"os"
)

func Logo() []byte {
	data, err := os.ReadFile("./assets/images/logo.png")
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	return data
}

func Previous() []byte {
	data, err := os.ReadFile("./assets/images/arrowLeft.png")
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	return data
}

func Next() []byte {
	data, err := os.ReadFile("./assets/images/arrowRight.png")
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	return data
}

func Reload() []byte {
	data, err := os.ReadFile("./assets/images/reload.png")
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	return data
}

func Tools() []byte {
	data, err := os.ReadFile("./assets/images/tools.png")
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	return data
}
