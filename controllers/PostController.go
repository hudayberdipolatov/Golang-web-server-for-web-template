package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"main/helpers"
	"net/http"
)

type Post struct {
}

func (post Post) PostIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("Posts/IndexPost")...)
	if err != nil {
		fmt.Println(err)
	}
	view.ExecuteTemplate(w, "PostIndex", nil)
}

func (post Post) SinglePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("Posts/SinglePost")...)
	if err != nil {
		fmt.Println(err)
	}
	view.ExecuteTemplate(w, "SinglePost", nil)
}
