package main

import(
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/binding"

)

type Show struct {
  Image         interface{}  `json:"image"`
  Slug          interface{}  `json:"slug"`
  Title         interface{}  `json:"title"`
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

func main() {
  m := martini.Classic()

  m.Use(render.Renderer(render.Options{
    Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
    IndentJSON: true, // Output human readable JSON
  }))

  m.Post("/drm", binding.Bind(DrmParams{}), func(params DrmParams, r render.Render) {
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
