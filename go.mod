module tests

go 1.19

require (
	browser v1.0.0
	common v1.0.0
	renderer v1.0.0
	network v1.0.0
	filesystem v1.0.0
)

replace (
	browser => ./browser
	common => ./common
	renderer => ./renderer
	network => ./network
	filesystem => ./filesystem
)