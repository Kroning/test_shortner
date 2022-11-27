/*
Makes handlers for routing. Returns pages and responses for API calls
*/
package handlers

import (
	"fmt"
	"net/http"
	"log"
	"html/template"

	_ "github.com/jackc/pgx/v5/pgxpool"
)

/* 
const tmplPath = "views" 

type Page struct {
    Title	string
		Db		*pgxpool.Pool
}

func NewPage(t string, db *pgxpool.Pool) Page {
	return Page{
		Title : t,
		Db :	db,
	} 
}
*/

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
		return
  }
}

func (p Page) saveHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is save!")
}

