module renderer

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
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69
)
