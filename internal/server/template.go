/*
template.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package server

import (
	"encoding/json"
	"errors"
	html_template "html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	text_template "text/template"

	"github.com/tdhite/kwite/pkg/funcs"
)

const (
	IsString = iota
	IsFile   = iota
)

const (
	TemplateKey = "template"
	ReadyKey    = "ready"
	AliveKey    = "alive"
	UrlKey      = "url"
	Kwite       = "kwite"
)

type Template struct {
	Template string
	Ready    string
	Alive    string
	Form     int
	name     string
	logger   *log.Logger
}

// Read the contents of a (configmap) file and return the string value
func readFile(path string) string {
	// open the file
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed reading ConfigMap entry at %s: %s\n", path, err)
	}
	return string(b)
}

// Creates a new text parsed Template
func (t *Template) newTextTemplate(ts string) (*text_template.Template, error) {
	var err error
	var tmpl *text_template.Template

	switch t.Form {
	case IsString:
		tmpl, err = text_template.New(Kwite).Funcs(funcs.TextTemplateFuncs()).Parse(ts)
	case IsFile:
		tmpl, err = text_template.New(Kwite).Funcs(funcs.TextTemplateFuncs()).ParseFiles(ts)
		t.name = filepath.Base(ts)
	default:
		err = errors.New("Unknown template form (neither of templateString or templateFile)")
	}

	if err != nil {
		t.logger.Println("Failed to parse Text template ", err)
		return nil, err
	}

	return tmpl, err
}

// Creates a new html parsed Template
func (t *Template) newHtmlTemplate(ts string) (*html_template.Template, error) {
	var err error
	var tmpl *html_template.Template

	switch t.Form {
	case IsString:
		tmpl, err = html_template.New(Kwite).Funcs(funcs.HtmlTemplateFuncs()).Parse(ts)
	case IsFile:
		tmpl, err = html_template.New(Kwite).Funcs(funcs.HtmlTemplateFuncs()).ParseFiles(ts)
		t.name = filepath.Base(ts)
	default:
		err = errors.New("Unknown template form (neither of templateString or templateFile)")
	}

	if err != nil {
		t.logger.Println("Failed to parse template ", err)
		return nil, err
	}

	return tmpl, nil
}

// Read http body as JSON and unmarshall into interface{}
func (t *Template) readJsonBody(r *http.Request, d *interface{}) error {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.logger.Println("Failed to read request Body: ", err)
		return err
	}
	err = json.Unmarshal([]byte(b), &d)
	if err != nil {
		t.logger.Printf("Failed to unmarshal Body JSON: %T\n%s\n%#v\n", err, err, err)
	}
	return nil
}

// Creates a new Template housing both text and html parsed templates
func NewTemplate(m map[string]string, form int, logger *log.Logger) (Template, error) {
	// All three templates are required
	if _, ok := m[TemplateKey]; !ok {
		return Template{}, errors.New("No template found.")
	}
	if _, ok := m[ReadyKey]; !ok {
		return Template{}, errors.New("No template found.")
	}
	if _, ok := m[AliveKey]; !ok {
		return Template{}, errors.New("No template found.")
	}

	t := Template{
		Form:     form,
		Template: m["template"],
		Ready:    m["ready"],
		Alive:    m["alive"],
		name:     Kwite,
		logger:   logger,
	}

	return t, nil
}

// Executes the parsed template as text; d is passed to the template executor.
// The result is written to to the io.Writer.
func (t *Template) ExecuteAsText(wr http.ResponseWriter, ts string, d interface{}) {
	if tmpl, err := t.newTextTemplate(ts); err != nil {
		t.logger.Println(err)
		wr.WriteHeader(http.StatusInternalServerError)
		return
	} else if err = tmpl.ExecuteTemplate(wr, t.name, d); err != nil {
		t.logger.Println(err)
		wr.WriteHeader(http.StatusInternalServerError)
	}
}

func (t *Template) ExecuteAsHtml(wr http.ResponseWriter, ts string, d interface{}) {
	if tmpl, err := t.newHtmlTemplate(ts); err != nil {
		t.logger.Println(err)
		wr.WriteHeader(http.StatusInternalServerError)
		return
	} else if err = tmpl.ExecuteTemplate(wr, t.name, d); err != nil {
		t.logger.Println(err)
		wr.WriteHeader(http.StatusInternalServerError)
	}
}

// Handles unknown http requests:
func (t *Template) HttpUnknownHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Kwite asked to handle URL but won't do that: " + r.URL.String())
	w.WriteHeader(http.StatusNotFound)
}

// Handles http request as follows:
//    If request Content-Type: application/json, json unmarshall Body as data
//    If request Content-Type: text/plain
func (t *Template) HttpHandler(w http.ResponseWriter, r *http.Request) {
	var d interface{}

	ct := r.Header.Get("Content-type")
	if ct == "application/json" {
		err := t.readJsonBody(r, &d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// select the right template string
	var ts string
	if strings.HasSuffix(r.RequestURI, "kwiteready") {
		ts = t.Ready
	} else if strings.HasSuffix(r.RequestURI, "kwitealive") {
		ts = t.Alive
	} else {
		ts = t.Template
	}

	accept := r.Header.Get("Accept")
	switch accept {
	case "application/json":
		{
			t.ExecuteAsText(w, ts, d)
		}
	case "text/plain":
		{
			t.ExecuteAsText(w, ts, d)
		}
	// all other media types go to HTML
	default:
		{
			t.ExecuteAsHtml(w, ts, d)
		}
	}
}
