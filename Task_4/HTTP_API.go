package main

import (
	"Backend_task_4/init"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type cats struct {
	name           string
	color          string
	tailLength     float32
	whiskersLength float32
}

func main() {
	http.HandleFunc("/cats", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	db := iinit.Init()
	rows, err := db.Query("SELECT * FROM cats")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	listСats := make([]*cats, 0)
	for rows.Next() {
		cat := new(cats)
		err := rows.Scan(&cat.name, &cat.color, &cat.tailLength, &cat.whiskersLength)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		listСats = append(listСats, cat)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, cat := range listСats {
		fmt.Fprintf(w, "%s, %s, %.2f, %.2f\n", cat.name, cat.color, cat.tailLength, cat.whiskersLength)
	}

}
