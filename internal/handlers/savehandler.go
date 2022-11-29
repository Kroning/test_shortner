/*
Handler for creating new short links
*/
package handlers

import (
  "fmt"
  "net/http"
  "log"
  "html/template"

  "github.com/jackc/pgx/v5"
)

// Creates a new link
func (p Page) saveHandler(w http.ResponseWriter, r *http.Request) {
  p.Title = "New"
  files := []string{
    tmplPath+"/base.tmpl",
    tmplPath+"/save_form.tmpl",
  }

  switch r.Method {
  case "GET":
  case "POST":
    // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
    if err := r.ParseForm(); err != nil {
      log.Println("ParseForm() err: ", err)
      http.Error(w, "Internal Server Error", 500)
      return
    }

    alias := r.FormValue("alias")
		url := r.FormValue("url")
		
    query := "SELECT id FROM links WHERE alias = $1 and deleted_at IS NULL;"
    row := p.Db.QueryRow(p.Ctx, query, alias)
    var id string
    err := row.Scan(&id)
    if err == nil {
			// No errors - link exists!
			p.Message = "Link already exists."
      files = append(files, tmplPath+"/message.tmpl")	
    } else {
      if err != pgx.ErrNoRows {
				// Real Error
        log.Println("Error ", err, " while Scan query ", query)
        http.Error(w, "Internal Server Error", 500)
        return
      }
			// pgx.ErrNoRows = we can create link
			query := "INSERT INTO links VALUES(default,$1,$2,NOW(),NULL);"
			_, err := p.Db.Exec(p.Ctx, query, alias, url)
			if err != nil {
				log.Println("Link innsertion error: ", err)
				http.Error(w, "Internal Server Error", 500)
				return
			}
			p.Message = "Link "+alias+" : "+url+" created."
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

