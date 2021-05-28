package handlers

import (
	"html/template"
	"net/http"
)

func schedulesPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("internal/templates/html/schedules.html", "internal/templates/html/header.html", "internal/templates/html/footer.html")
	tmpl.ExecuteTemplate(w, "schedules", nil)
}
