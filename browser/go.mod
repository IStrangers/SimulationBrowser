module browser

go 1.19

require (
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220806181222-55e207c401ad
	renderer v1.0.0
)

require common v1.0.0 // indirect

replace (
	common => ../common
	renderer => ../renderer
)
