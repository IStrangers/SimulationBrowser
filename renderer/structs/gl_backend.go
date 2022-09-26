package structs

type GLBackend struct {
	program uint32

	vao uint32
	vbo uint32

	texture uint32
	quad    []float32
}
