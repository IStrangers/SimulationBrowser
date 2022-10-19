module layout

go 1.19

require renderer v1.0.0

require (
	assets v1.0.0
	common v1.0.0 // indirect
	filesystem v1.0.0
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c // indirect
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
