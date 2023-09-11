package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"main/helpers"
	"net/http"
)

type News struct {
}

func (new News) NewIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("News/IndexNew")...)
	if err != nil {
		fmt.Println(err)
	}
	view.ExecuteTemplate(w, "NewIndex", nil)
}

func (new News) SingleNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("News/SingleNew")...)
	if err != nil {
		fmt.Println(err)
	}
	view.ExecuteTemplate(w, "SingleNew", nil)
}
