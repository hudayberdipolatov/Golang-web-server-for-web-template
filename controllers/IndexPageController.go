package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"main/helpers"
	"net/http"
)

type IndexData struct {
}

func (index IndexData) IndexPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("IndexPage")...)
	if err != nil {
		fmt.Println(err)
	}
	view.ExecuteTemplate(w, "IndexPage", nil)
}
