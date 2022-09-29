package structs

import "fmt"

type NoSuchElementError string

func (e NoSuchElementError) Error() string {
	return fmt.Sprintf("no such element: %q", string(e))
}
