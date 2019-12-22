package main

import (
	"Backend_task_4/login"
	"database/sql"
	"encoding/json"
	"fmt"
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
	//router.GET("/cats/:attribute", getList)
	http.ListenAndServe(":8080", router)
}

func getList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var str1 string
	var str2 string
	if len(r.URL.RawQuery) > 0 {
		str1 = r.URL.Query().Get("attribute")
		str2 = r.URL.Query().Get("order")
		if str1 == "" || str2 == "" {
			w.WriteHeader(400)
			return
		}
	}
	catslist, err := catsList(str1, str2)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(catslist); err != nil {
		w.WriteHeader(500)
	}
}
func catsList(str1 string, str2 string) ([]cats, error) {
	db := login.Init()
	var rows *sql.Rows
	var err error
	if str1 != "" || str2 != "" {
		param := str1 + " " + str2
		rows, err = db.Query(fmt.Sprintf("SELECT * FROM cats ORDER BY %s", param))
	}
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
