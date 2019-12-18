package main

import (
	"Backend_task_4/login"
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net/http"
)

type cats struct {
	Name           string  `json:"name"`
	Color          string  `json:"color"`
	TailLength     float32 `json:"tail_length"`
	WhiskersLength float32 `json:"whiskers_length"`
}

func main() {
	router := httprouter.New()
	router.GET("/cats", getList)
	http.ListenAndServe(":8080", router)
}

func getList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var str string
	if len(r.URL.RawQuery) > 0 {
		str = r.URL.Query().Get("cats")
		if str == "" {
			w.WriteHeader(400)
			return
		}
	}
	catslist, err := catsList(str)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(catslist); err != nil {
		w.WriteHeader(500)
	}
}
func catsList(str string) ([]cats, error) {
	db := login.Init()
	var rows *sql.Rows
	var err error
	rows, err = db.Query("SELECT * FROM cats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]cats, 0)
	var cat cats
	for rows.Next() {
		if err = rows.Scan(&cat.Name, &cat.Color, &cat.TailLength, &cat.WhiskersLength); err != nil {
			return nil, err
		}
		list = append(list, cat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
