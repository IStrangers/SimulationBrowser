module tests

go 1.19

require (
	browser v1.0.0
	filesystem v1.0.0
	renderer v1.0.0
	common v1.0.0
	network v1.0.0
)

require (
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220806181222-55e207c401ad // indirect
)

replace (
	browser => ./browser
	common => ./common
	filesystem => ./filesystem
	network => ./network
	renderer => ./renderer
)
