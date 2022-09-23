package common

import "strings"

func IsBase64Image(content string) bool {
	return strings.HasPrefix(content, "data:image/png;base64")
}
