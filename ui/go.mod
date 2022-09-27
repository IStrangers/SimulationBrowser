module ui

go 1.19

require (
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220806181222-55e207c401ad
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c
	renderer v1.0.0
)

require (
	assets v1.0.0
	common v1.0.0 // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69 // indirect
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
