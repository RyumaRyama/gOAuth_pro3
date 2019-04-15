package main

import (
  "net/http"
  "github.com/labstack/echo"
  "html/template"
  "io"
)
type Template struct {
    templates *template.Template
}
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}
type ServiceInfo struct {
  Title string
}

var serviceInfo = ServiceInfo {
  "We are goAuth_pro3",
}


func main(){

  t := &Template{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }

  e := echo.New()

  e.Renderer = t

  e.Static("/public", "./public/")
  e.Static("/public/css", "./public/css")

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "こんにちは")
  })

  e.GET("/sample", func(c echo.Context) error {

    data := struct {
      ServiceInfo
      Content_a string
      Content_b string
      Content_c string
      Content_d string
    } {
      ServiceInfo: serviceInfo,
      Content_a: "goを使っててて",
      Content_b: "oauth認証をしよう。",
      Content_c: "アイウエオ",
      Content_d: "かきくけこ",
    }

    return c.Render(http.StatusOK, "sample", data)
  })

  e.GET("/login", func(c echo.Context) error {
    return c.Render(http.StatusOK, "login", "momo")
  })

  e.Logger.Fatal(e.Start(":1124"))
  }
