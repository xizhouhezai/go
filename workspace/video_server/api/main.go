package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9990", r)
}
