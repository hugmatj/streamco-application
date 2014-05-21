package main

import(
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/binding"

  "fmt"
)

type Show struct {
  Image         string  `json:"image"`
  Slug          string  `json:"slug"`
  Title         string  `json:"title"`
  Drm           bool    `json:"drm"`
  EpisodeCount  int     `json:"episodeCount"`
}

type DrmParams struct {
  Payload       []Show    `json:"payload"`
  Skip          int       `json:"skip"`
  Take          int       `json:"take"`
  TotalRecords  int       `json:"totalRecords"`
}

func ok(r render.Render, shows []Show) {
  r.JSON(200, map[string] []Show { "response": shows })
}

func err(r render.Render, message string) {
  r.JSON(400, map[string] string { "error": message })
}

func errorHandler(errors binding.Errors, r render.Render) {
  if errors.Len() > 0 {
    e := errors[0]
    message := fmt.Sprintf("%s: %s", e.Classification, e.Message)

    err(r, message)
  }
}

func main() {
  m := martini.Classic()

  m.Use(render.Renderer(render.Options{
    Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
    IndentJSON: true, // Output human readable JSON
  }))

  m.Get("/", func() string {
    return "hello world"
  });

  m.Post("/drm", binding.Json(DrmParams{}), errorHandler, func(params DrmParams, r render.Render) {
    valid := make([]Show, 0)

    for _, show := range params.Payload {
      if show.Drm && show.EpisodeCount > 0 {
        valid = append(valid, show)
      }
    }

    ok(r, valid)
  })

  m.Run()
}
