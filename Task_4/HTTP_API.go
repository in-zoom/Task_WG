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
	//router.GET("/cats", getList)
	router.GET("/cats", ggetList)
	http.ListenAndServe(":8080", router)
}

func ggetList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var param1 string
	var param2 string
	var param3 string
	var param4 string

	if len(r.URL.RawQuery) > 0 {
		param1 = r.URL.Query().Get("attribute")
		param2 = r.URL.Query().Get("order")
		param3 = r.URL.Query().Get("offset")
		param4 = r.URL.Query().Get("limit")
		/*if param1 == "" || param2 == "" || param3 == "" || param4 == "" {
		m3 == "" || param4 == "" {
					w.WriteHeader(400)
					return
				}*/
	}
	catslist, err := catsList(param1, param2, param3, param4)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(catslist); err != nil {
		w.WriteHeader(500)
	}
}
func catsList(param1 string, param2 string, param3 string, param4 string) ([]cats, error) {
	db := login.Init()
	var rows *sql.Rows
	var err error
	if param1 == "" && param2 == "" && param3 == "" && param4 == "" {
		rows, err = db.Query("SELECT * FROM cats")
	} else if param1 != "" && param2 != "" && param3 == "" && param4 == "" {
		param := param1 + " " + param2
		rows, err = db.Query(fmt.Sprintf("SELECT * FROM cats ORDER BY %s", param))
	} else if param1 == "" && param2 == "" && param3 != "" && param4 != "" {
		rows, err = db.Query(fmt.Sprintf("SELECT * FROM cats offset %s limit %s", param3, param4))
	} else if param1 != "" || param2 != "" && param3 != "" || param4 != "" {
		param := param1 + " " + param2
		rows, err = db.Query(fmt.Sprintf("SELECT * FROM cats ORDER BY %s offset %s limit %s", param, param3, param4))
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
