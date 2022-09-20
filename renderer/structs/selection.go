package structs

type Selection struct {
	StartCursor *Cursor
	EndCursor   *Cursor
	HTML        string
}
