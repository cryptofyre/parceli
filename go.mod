module cryptofyre.org/parceli

go 1.17

require (
	github.com/akamensky/argparse v1.3.1
	brandonplank.org/parcelilib v0.0.0
)

replace (
	brandonplank.org/parcelilib => ./library
)