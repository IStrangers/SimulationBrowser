module browser

go 1.19

require (
	renderer v0.0.0-00010101000000-000000000000
	ui v0.0.0-00010101000000-000000000000
)

require (
	assets v1.0.0 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220806181222-55e207c401ad // indirect
)

require (
	common v1.0.0 // indirect
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69 // indirect
	profiler v1.0.0
)

replace (
	assets => ../assets
	browser => ../browser
	common => ../common
	filesystem => ../filesystem
	network => ../network
	profiler => ../profiler
	renderer => ../renderer
	ui => ../ui
)
