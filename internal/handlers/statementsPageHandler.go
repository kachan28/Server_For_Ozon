package handlers

import (
	"html/template"
	"net/http"
)

func statementsPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("internal/templates/html/statements.html", "internal/templates/html/header.html", "internal/templates/html/footer.html")
	tmpl.ExecuteTemplate(w, "statements", nil)
}
