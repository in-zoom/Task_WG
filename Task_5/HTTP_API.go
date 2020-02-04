package main

import (
	"Backend/task_5/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/ping", handlers.IndexCatsService)
	router.GET("/cats", handlers.Getlist)
	router.POST("/cat", handlers.AddCat)
    http.ListenAndServe(":8080", router)
}
