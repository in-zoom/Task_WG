package main

import (
	"Backend_task_4/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/cats", handlers.Getlist)
	http.ListenAndServe(":8080", router)
}
