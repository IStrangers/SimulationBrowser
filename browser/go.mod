module browser

go 1.19

require (
	assets v1.0.0
	layout v0.0.0-00010101000000-000000000000
	renderer v1.0.0
	ui v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220806181222-55e207c401ad // indirect
	network v0.0.0-00010101000000-000000000000 // indirect
)

require (
	common v1.0.0 // indirect
	filesystem v1.0.0
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69 // indirect
	profiler v1.0.0
)

replace (
	assets => ../assets
	browser => ../browser
	common => ../common
	filesystem => ../filesystem
	layout => ../layout
	network => ../network
	profiler => ../profiler
	renderer => ../renderer
	ui => ../ui
)
