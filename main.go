package main

import (
	"fmt"
	"net/http"

	Router "controllers"
)

func main() {
	fmt.Println("Start")
	r := Router.NewRouter()
	http.ListenAndServe(":5000", r)
}
