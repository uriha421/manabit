package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
  "text/template"
  "path/filepath"
  "sync"
)

type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func() {
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  if err := t.templ.Execute(w, nil); err != nil {
    log.Fatal(err)
  }
}

func main() {
	r := mux.NewRouter()
	r.Handle("/uploader", &templateHandler{filename: "uploader.html"})
	r.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
