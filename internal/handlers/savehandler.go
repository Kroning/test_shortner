/*
Handler for creating new short links
*/
package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	store "github.com/Kroning/test_shortner/internal/storage"
)

// Creates a new link
func (p Page) saveHandler(w http.ResponseWriter, r *http.Request) {
	p.Title = "New"
	files := []string{
		tmplPath + "/base.tmpl",
		tmplPath + "/save_form.tmpl",
	}

	switch r.Method {
	case "GET":
	case "POST":
		// Parsing FORM
		if err := r.ParseForm(); err != nil {
			log.Println("ParseForm() err: ", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		alias := r.FormValue("alias")
		url := r.FormValue("url")

		// Looking for duplicate in database
		_, err := p.Storage.CheckLinkExistance(p.Ctx, alias)
		if err == nil {
			// No errors - link exists!
			p.Message = "Link already exists."
			files = append(files, tmplPath+"/message.tmpl")
		} else {
			if err != store.LinkNotExists {
				// Real Error
				http.Error(w, "Internal Server Error", 500)
				return
			}
			err := p.Storage.InsertLink(p.Ctx, alias, url)
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			p.Message = "Link " + alias + " : " + url + " created."
			files = append(files, tmplPath+"/message.tmpl")
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		http.Error(w, "Forbidden", 403)
		return
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
