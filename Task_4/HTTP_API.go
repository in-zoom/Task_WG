package main

import (
	"Backend/task_4/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/cats", handlers.Getlist)
	router.POST("/cat", handlers.Addcat)
    http.ListenAndServe(":8080", router)
}
