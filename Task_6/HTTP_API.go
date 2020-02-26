package main

import (
	
	"github.com/julienschmidt/httprouter"
	"Backend_task_6/middleware"
	"Backend_task_6/handlers"
	"net/http"
)

func main() {
	router := httprouter.New()
	//router.GET("/ping", handlers.IndexCatsService)
	router.GET("/ping", middleware.LimitRequestsMiddleware(middleware.CatsServiceHandler()))
	router.GET("/cats", middleware.LimitRequestsMiddleware(middleware.CatsListHandler()))
	router.POST("/cat", handlers.AddCat)
	http.ListenAndServe(":8080", router)
}
