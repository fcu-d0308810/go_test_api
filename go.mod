module main

go 1.17

require (
	github.com/gorilla/mux v1.8.0
	github.com/rs/cors v1.8.2

	controllers v1.0.0
)

replace (
	controllers v1.0.0 => ./api
)
