package main

import "os"
import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"
import "github.com/martini-contrib/binding"
import "gopkg.in/mgo.v2"

type Post struct {
	Title   string `form:"title" json:"title"`
	Summary string `form:"summary" json:"summary"`
	Content string `form:"content" json:"content"`
}

func main() {

	session, _ := mgo.Dial(os.Getenv("DB_PORT_27017_TCP_ADDR"))
	db := session.DB("blog")
	defer session.Close()

	app := martini.Classic()
	app.Use(render.Renderer())

	app.Get("/", func(r render.Render) {
		var posts []Post
		db.C("posts").Find(nil).All(&posts)
		r.HTML(200, "blog", posts)
	})

	app.Get("/json", func(r render.Render) {
		var posts []Post
		db.C("posts").Find(nil).All(&posts)
		r.JSON(200, posts)
	})

	app.Post("/new", binding.Bind(Post{}), func(r render.Render, post Post) {
		db.C("posts").Insert(post)
	})

	app.Run()

}
