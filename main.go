package main

import (
	"embed"
	"github.com/google/uuid"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"time"
)

//go:embed html
var html embed.FS

func main() {
	staticFiles, err := fs.Sub(html, "html/static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/now", now)
	mux.HandleFunc("/uuid", UUID)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(staticFiles))))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(html, "html/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(w, "index", struct {
		Name string
		Now  string
		UUID string
	}{
		Name: "Person",
		Now:  time.Now().String(),
		UUID: uuid.NewString(),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func now(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(html, "html/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(w, "now", struct{ Now string }{
		Now: time.Now().String(),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UUID(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(html, "html/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(w, "uuid", struct{ UUID string }{
		UUID: uuid.NewString(),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
