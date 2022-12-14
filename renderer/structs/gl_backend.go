package structs

import (
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	"log"
	"strings"
)

type GLBackend struct {
	program uint32

	vao uint32
	vbo uint32

	Texture uint32
	quad    []float32
}

var quad = []float32{
	1.0, 1.0, 1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,

	-1.0, 1.0, 1.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
}

func createNewGLProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	gl.Init()
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func CreateGLBackend() *GLBackend {
	program, err := createNewGLProgram(vertexShader, fragmentShader)
	if err != nil {
		log.Fatal(err)
	}

	gl.UseProgram(program)
	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)
	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	backend := &GLBackend{
		program: program,
	}

	gl.GenVertexArrays(1, &backend.vao)
	gl.BindVertexArray(backend.vao)

	gl.GenBuffers(1, &backend.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, backend.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(quad)*4, gl.Ptr(quad), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(backend.program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(backend.program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	return backend
}
