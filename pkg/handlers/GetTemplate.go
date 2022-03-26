package handlers

import (
	"bytes"
	"html/template"
	"net/http"
)

func (h handler) GetTemplate(w http.ResponseWriter, r *http.Request) {
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

func (h handler) GetTemplateRegister(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/register.html"))
	var buf bytes.Buffer
	err := template.Execute(&buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "text/html; charset=UTF-8")
	buf.WriteTo(w)
}