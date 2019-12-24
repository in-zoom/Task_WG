package handlers

import (
    "encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"   
)


type Cats struct {
	Name           string  `json:"name"`
	Color          string  `json:"color"`
	TailLength     float32 `json:"tail_length"`
	WhiskersLength float32 `json:"whiskers_length"`
}

func Getlist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var param1 string
	var param2 string
	var param3 string
	var param4 string

	if len(r.URL.RawQuery) > 0 {
		param1 = r.URL.Query().Get("attribute")
		param2 = r.URL.Query().Get("order")
		param3 = r.URL.Query().Get("offset")
		param4 = r.URL.Query().Get("limit")
	}
	resultCatslist, err := Catslist(param1, param2, param3, param4)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(resultCatslist); err != nil {
		w.WriteHeader(500)
	}
}
