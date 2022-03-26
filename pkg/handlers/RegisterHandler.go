package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func (h handler) RegisteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**RegisterHandler running**")
	tpl, _ = template.ParseGlob("template/*html")
	tpl.ExecuteTemplate(w, "register.html", nil)
}
