package structs

type Selection struct {
	StartCursor *Cursor
	EndCursor   *Cursor
	Content     string
}
