package main

import (
	"github.com/julienschmidt/httprouter"
	"Backend_task_6/middleware"
	"github.com/joho/godotenv"
	"Backend_task_6/handlers"
	"net/http"
    "log"
)

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}

func main() {
	initialize()
	router := httprouter.New()
	router.GET("/ping", middleware.LimitRequestsMiddleware(middleware.CatsServiceHandler()))
	router.GET("/cats", middleware.LimitRequestsMiddleware(middleware.CatsListHandler()))
	router.POST("/cat", handlers.AddCat)
	http.ListenAndServe(":8080", router)
}
