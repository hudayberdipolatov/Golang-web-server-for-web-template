package routes

import (
	"github.com/julienschmidt/httprouter"
	controller "main/controllers"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	r.GET("/", controller.IndexData{}.IndexPage)
	// news route begin
	r.GET("/news", controller.News{}.NewIndex)
	r.GET("/news/single-new", controller.News{}.SingleNew)
	// news route end
	// post route begin
	r.GET("/posts", controller.Post{}.PostIndex)
	r.GET("/posts/single-post", controller.Post{}.SinglePost)
	// post route end
	//  server file
	r.ServeFiles("/assets/*filepath", http.Dir("assets"))
	return r
}
