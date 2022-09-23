package structs

import (
	"regexp"
	"strconv"
)

var elementFontTable = map[string]float64{
	"h1": float64(32),
	"h2": float64(28),
	"h3": float64(20),
	"p":  float64(14),
}

func ConvertSizeValue(sizeValue string) float64 {
	re := regexp.MustCompile("[0-9]+")
	valueString := re.FindString(sizeValue)
	value, err := strconv.ParseInt(valueString, 10, 0)

	if err != nil {
		return float64(14)
	}

	return float64(value)
}

func GetDefaultElementFontWeight(nodeName string) int {
	switch nodeName {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		return 600
	}
	return 400
}
