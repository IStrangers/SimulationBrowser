module renderer

go 1.19

require (
	common v1.0.0
	layout v1.0.0
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69
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
