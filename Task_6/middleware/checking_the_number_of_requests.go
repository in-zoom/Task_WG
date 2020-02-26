package middleware

import (
	"github.com/julienschmidt/httprouter"
	"Backend_task_6/handlers"
	"golang.org/x/time/rate"
	"net/http"
)

var limit = rate.NewLimiter(10, 10)

func LimitRequestsMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if !limit.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next(w, r, ps)
	}
}

func CatsServiceHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.IndexCatsService(w, r, ps)
	}
}

func CatsListHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.Getlist(w, r, ps)

	}
}

