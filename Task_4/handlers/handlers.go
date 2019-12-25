package handlers

import (
	"Backend_task_4/validation"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type cat struct {
	Name           string  `json:"name"`
	Color          string  `json:"color"`
	TailLength     float32 `json:"tail_length"`
	WhiskersLength float32 `json:"whiskers_length"`
}

func Getlist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var attribute, order string
	//var order string
	//var param3 string
	//var param4 string

	if len(r.URL.RawQuery) > 0 {
		attribute = r.URL.Query().Get("attribute")
		order = r.URL.Query().Get("order")

	}
	resultparam, errr := validation.WhiteParam(attribute, order)
	if errr != nil {
		w.WriteHeader(400)
		return
	}
	resultCatslist, err := сatslist(resultparam)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(resultCatslist); err != nil {
		w.WriteHeader(500)
	}
}
