module tests

go 1.19

require (
	browser v1.0.0
	filesystem v1.0.0
	renderer v1.0.0
)

require (
	assets v1.0.0 // indirect
	common v1.0.0 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220806181222-55e207c401ad // indirect
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69 // indirect
	layout v0.0.0-00010101000000-000000000000 // indirect
	network v1.0.0 // indirect
	profiler v1.0.0 // indirect
	ui v1.0.0 // indirect
)

replace (
	assets => ./assets
	browser => ./browser
	common => ./common
	filesystem => ./filesystem
	layout => ./layout
	network => ./network
	profiler => ./profiler
	renderer => ./renderer
	ui => ./ui
)
