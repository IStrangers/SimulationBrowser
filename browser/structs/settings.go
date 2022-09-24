package structs

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Settings struct {
	HomePage string

	WindowWidth int
	WindowHeight int

	HiDPI bool
}

var defaultSettings = Settings{
	HomePage: WebBrowserName + "://HomePage",

	WindowWidth:  600,
	WindowHeight: 600,

	HiDPI: true,
}

func LoadSettings(path string) *Settings {
	settingsPath := flag.String("settings", path, "This flag sets the location for the browser settings file.")
	flag.Parse()
	settingsData,err := os.ReadFile(*settingsPath)

	if err != nil {
		fmt.Println("Unable to read settings file", settingsPath)
		fmt.Println("Loading default settings.")
		return &defaultSettings
	}

	err = json.Unmarshal(settingsData, &defaultSettings)
	if err != nil {
		fmt.Println("Error loading settings from file;", err)
	}

	return &defaultSettings
}