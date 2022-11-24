package adminhandler

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

const tmplPath = "views" // "../../../views"

type Page struct {
    Title	string
		db		string
}

func NewPage(t string, db string) Page {
	return Page{
		Title : t,
		db :	db,
	} 
}

func (p Page) MainInitHandlers() {
	http.HandleFunc("/", p.mainHandler)
  http.HandleFunc("/save", p.saveHandler)
}

func (p Page) mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
    return
  }
	p.Title = "Search"
	files := []string{
		tmplPath+"/base.tmpl",
    tmplPath+"/search_form.tmpl",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil {
		log.Print(err.Error())
    http.Error(w, "Internal Server Error", 500)
    return
  }

  err = ts.ExecuteTemplate(w, "base", p)
  if err != nil {
		log.Print(err.Error())
    http.Error(w, "Internal Server Error", 500)
  }
  //fmt.Fprintf(w, "This is main!")
}

func (p Page) saveHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is save!")
}

