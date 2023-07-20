package util

import (
	"html/template"
	"net/http"
)

func RenderHandle(filename string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, err := template.ParseFiles(filename)
			if err != nil {
				panic(err)
			}
			t.Execute(w, nil)
		}
	}
}
