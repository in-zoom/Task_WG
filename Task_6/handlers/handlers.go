package handlers

import (
	"Backend_task_6/login"
	"Backend_task_6/validation"
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth/limiter"
	"github.com/julienschmidt/httprouter"
)

type cat struct {
	Name           string `json:"name"`
	Color          string `json:"color"`
	TailLength     int    `json:"tail_length"`
	WhiskersLength int    `json:"whiskers_length"`
}

type errMessage struct {
	Message string `json:"message"`
}

type Limiter struct{
	x int
	max int64
   r  *limiter.Limiter

}

func IndexCatsService(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	fmt.Fprintf(w, "Cats Service. Version 0.1")
}

func Getlist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var attribute, order, offset, limit string
	url := r.URL.Query()

	if len(r.URL.RawQuery) > 0 {
		attribute = url.Get("attribute")
		order = url.Get("order")
		offset = url.Get("offset")
		limit = url.Get("limit")

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resultAttribute, err := validation.ValidateAttribute(attribute)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultOrder, err := validation.ValidateOrder(order)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultOffset, err := validation.ValidateOffset(offset)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultLimit, err := validation.ValidateLimit(limit)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultCatslist, err := сatslist(resultAttribute, resultOrder, resultOffset, resultLimit)
	if err != nil {
		responseError(w, 500, err)
		return
	}
	if err = json.NewEncoder(w).Encode(resultCatslist); err != nil {
		responseError(w, 500, err)
	}
}

func AddCat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var addedCat cat
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(r.Body).Decode(&addedCat)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultNameCat, err := validation.ValidateName(addedCat.Name, login.Init())
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultColorCat, err := validation.ValidateColor(addedCat.Color)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultTailLengthCat, err := validation.ValidTailLength(addedCat.TailLength)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultWhiskersLengthCat, err := validation.ValidWhiskersLength(addedCat.WhiskersLength)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	errr := addNewCat(resultNameCat, resultColorCat, resultTailLengthCat, resultWhiskersLengthCat)
	if errr != nil {
		responseError(w, 500, err)
		return
	}

}

func responseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	errMessage := errMessage{err.Error()}
	json.NewEncoder(w).Encode(errMessage)
}
