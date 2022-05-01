package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type IndexPage struct {
	DefaultSearchTerm string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := IndexPage{DefaultSearchTerm: "Дозор"}
	t := template.Must(template.ParseFiles("./templates/index.html"))
	err := t.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
}
