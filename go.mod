module main

go 1.17

require controllers v1.0.0

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/rs/cors v1.8.2 // indirect
)

replace controllers v1.0.0 => ./api
