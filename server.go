package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  // render html templates from templates directory
  m.Use(render.Renderer(render.Options{
    Layout: "index",
  }))

  m.Get("/", func(r render.Render) string {
    r.HTML(200, "index", "jeremy")
    return "string"
  })

  m.Get("/:postid", func( params martini.Params ) string {
    return "<h1>Post "+ params["postid"] +"!</h1>"
  })

  m.Run()
}