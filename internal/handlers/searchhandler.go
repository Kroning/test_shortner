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

// Main page/search http handler
// Shows links if there are something in POST 
func (p Page) mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
    return
  }

	// Mandatory templates
	p.Title = "Search"
  files := []string{
    tmplPath+"/base.tmpl",
    tmplPath+"/search_form.tmpl",
  }

	switch r.Method {
	case "GET":		
		 //http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			log.Println("ParseForm() err: ", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("alias")
		query := "SELECT id, alias, url, TO_CHAR(created_at, 'yyyy-mm-dd hh:mm:ss') FROM links WHERE alias LIKE $1;"
		rows, err := p.Db.Query(p.Ctx, query, "%"+name+"%")
		defer rows.Close()
		if err != nil {
			log.Println("Error ", err, " while executing query ", query)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		for rows.Next() {
	    var id,alias,url,created string
		  err = rows.Scan(&id, &alias, &url, &created)
			if err != nil {
				log.Println("Error ", err, " while Scan query ", query)
	      http.Error(w, "Internal Server Error", 500)
        return 
	    }
			slice := []string{id,alias,url,created}
			p.Result = append(p.Result,slice)
		}
		if len(p.Result) == 0 {
			p.Message = "No links found."
			files = append(files, tmplPath+"/message.tmpl")	
		} else {
			files = append(files, tmplPath+"/result.tmpl")
		}

		if err = rows.Err(); err != nil {
			log.Println("Error ", err, " while Scan query ", query)
			http.Error(w, "Internal Server Error", 500)
			return
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


