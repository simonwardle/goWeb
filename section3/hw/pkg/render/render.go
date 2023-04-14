package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/simonwardle/goWeb/pkg/config"
)

var app *config.AppConfig

//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get template cache from app config
	tmplCache := app.TemplateCache

	// get requested template from cache
	t, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("error: could not get template from template cache.")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.tmpl.html from ./templates
	tmpls, err := filepath.Glob("./templates/*.tmpl.html")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.tmpl.html
	for _, tmpl := range tmpls {
		//get the name of the file and create a template set
		name := filepath.Base(tmpl)
		tmplset, err := template.New(name).ParseFiles(tmpl)
		if err != nil {
			return myCache, err
		}
		//get any layout files
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		//if we have any in the directory add then all to the template set
		if len(layouts) > 0 {
			tmplset, err = tmplset.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		//populate myCache with the complete template set
		//e.g myCache[home.tmpl.html] will have the home.tmpl.html & base.layout.tmpl
		myCache[name] = tmplset
	}
	return myCache, nil
}
