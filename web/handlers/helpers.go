package handlers

import (
	"bytes"
	"log"
	"net/http"
	"text/template"

	"github.com/davecgh/go-spew/spew"
)

const (
	ErrorMessage = `Error: ***** {{ .Err }} ******
	
Request: {{ .R }}
	`
)

type ErrorPage struct {
	Err error
	R   string
}

func ErrorToBrowser(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Add("Content-Type", "text/plain")
	t := template.Must(template.New("ErrorMessage").Parse(ErrorMessage))
	rstring := spew.Sdump(r)
	data := ErrorPage{Err: err, R: rstring}
	var buf bytes.Buffer
	e := t.Execute(&buf, &data)
	if e != nil {
		log.Fatalf("Failed to create error page: %+v", e)
	}

	w.Write(buf.Bytes())
}
