package controllers

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, tmpl_err := template.ParseFiles("views/base.html")
	if tmpl_err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
}
