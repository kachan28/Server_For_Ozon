package handlers

import (
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("internal/templates/html/index.html", "internal/templates/html/header.html", "internal/templates/html/footer.html")
	tmpl.ExecuteTemplate(w, "index", true)
}
