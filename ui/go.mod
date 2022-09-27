module ui

go 1.19

replace (
	browser => ../browser
	common => ../common
	filesystem => ../filesystem
	network => ../network
	profiler => ../profiler
	renderer => ../renderer
	ui => ../ui
)

require (
	browser v0.0.0-00010101000000-000000000000
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220806181222-55e207c401ad
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c
	renderer v0.0.0-00010101000000-000000000000
)

require (
	common v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69 // indirect
)
