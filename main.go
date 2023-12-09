package main

import (
	"html/template"
	"net/http"
)

func main() {
	templates := loadTemplates()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})

	http.ListenAndServe(":8080", nil)
}

func loadTemplates() *template.Template {
	result := template.New("template")
	basePath := "./template/"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
