package main

import(
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/binding"
)

type Show struct {
  Image         map[string] string  `json:"image"`
  Slug          string              `json:"slug"`
  Title         string              `json:"title"`
  Drm           bool                `json:"drm"`
  EpisodeCount  int                 `json:"episodeCount"`
}

type DrmResponse struct {
  Image   string  `json:"image"`
  Slug    string  `json:"slug"`
  Title   string  `json:"title"`
}

func newDrmResponseFrom(show * Show) DrmResponse {
  return DrmResponse {
    show.Image["showImage"],
    show.Slug,
    show.Title,
  }
}

type DrmParams struct {
  Payload       []Show    `json:"payload"`
  Skip          int       `json:"skip"`
  Take          int       `json:"take"`
  TotalRecords  int       `json:"totalRecords"`
}

func ok(r render.Render, shows []DrmResponse) {
  r.JSON(200, map[string] []DrmResponse { "response": shows })
}

func err(r render.Render, message string) {
  r.JSON(400, map[string] string { "error": message })
}

func errorHandler(errors binding.Errors, r render.Render) {
  if errors.Len() > 0 {
    err(r, "Could not decode request")
  }
}

func drmHandler(params DrmParams, r render.Render) {
  valid := make([]DrmResponse, 0)

  for _, show := range params.Payload {
    if show.Drm && show.EpisodeCount > 0 {
      valid = append(valid, newDrmResponseFrom(& show))
    }
  }

  ok(r, valid)
}

func NewMartiniServer() * martini.ClassicMartini {
  m := martini.Classic()

  m.Use(render.Renderer(render.Options{
    Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
    IndentJSON: true, // Output human readable JSON
  }))

  m.Get("/hello", func() string { return "Hello world" })
  m.Post("/", binding.Json(DrmParams{}), errorHandler, drmHandler)
  return m
}

func main() {
  m := NewMartiniServer()
  m.Run()
}
