package handlers

import (
	"Backend/task_5/validation"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type cat struct {
	Name           string `json:"name"`
	Color          string `json:"color"`
	TailLength     int    `json:"tail_length"`
	WhiskersLength int    `json:"whiskers_length"`
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
	resultAttribute, errr := validation.ValidateAttribute(attribute)
	if errr != nil {
		w.WriteHeader(400)
		return
	}
	resultOrder, errr := validation.ValidateOrder(order)
	if errr != nil {
		w.WriteHeader(400)
		return
	}
	resultOffset, errrr := validation.ValidateOffset(offset)
	if errrr != nil {
		w.WriteHeader(400)
		return
	}
	resultLimit, errrrr := validation.ValidateLimit(limit)
	if errrrr != nil {
		w.WriteHeader(400)
		return
	}
	resultCatslist, err := сatslist(resultAttribute, resultOrder, resultOffset, resultLimit)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(resultCatslist); err != nil {
		w.WriteHeader(500)
	}
}

func Addcat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var addedCat cat
	err := json.NewDecoder(r.Body).Decode(&addedCat)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	resultNameCat, errr := validation.ValidateName(addedCat.Name)
	if errr != nil {
		w.WriteHeader(500)
		return
	}

	resultColorCat, errr := validation.ValidateColor(addedCat.Color)
	if errr != nil {
		w.WriteHeader(500)
		return
	}

	resultTailLengthCat, errr := validation.ValidTailLength(addedCat.TailLength)
	if errr != nil {
		w.WriteHeader(500)
		return
	}

	resultWhiskersLengthCat, err := validation.ValidWhiskersLength(addedCat.WhiskersLength)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	errrr := addNewCat(resultNameCat, resultColorCat, resultTailLengthCat, resultWhiskersLengthCat)
	if errrr != nil {
		w.WriteHeader(400)
		return
	}

}
