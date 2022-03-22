package handlers

import (
	"bytes"
	"net/http"
	"html/template"
)

func GetTemplate(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/template.html"))
	var buf bytes.Buffer
	err := template.Execute(&buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "text/html; charset=UTF-8")
	buf.WriteTo(w)
}