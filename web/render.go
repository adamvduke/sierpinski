package web

import (
	"embed"
	"encoding/json"
	"html/template"
	"io"

	"github.com/adamvduke/sierpinski/chaos"
)

const (
	templateVisualizeHTML = "visualize.html.tmpl"
	templatePlotJS        = "plot.js.tmpl"
	templatesGlob         = "templates/*.tmpl"
)

//go:embed templates
var templates embed.FS

// RGB represents a color used to draw a sierpinski triangle.
type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// RenderOpts represents the options available when rendering a sierpinski
// triangle.
type RenderOpts struct {
	Chaos       *chaos.Triangle `json:"chaos"`
	PointRadius float64         `json:"point_radius"`
	PointColor  *RGB            `json:"point_color"`
}

// WriteHTML uses the given Triangle and RenderOpts to render an html template
// and write it to the given io.Writer.
func WriteHTML(writer io.Writer, opts *RenderOpts) error {
	tmpl, err := template.ParseFS(templates, templatesGlob)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(writer, templateVisualizeHTML, opts)
}

// WriteJS uses the given Triangle and RenderOpts to render a javascript template
// and write it to the given io.Writer.
func WriteJS(writer io.Writer, opts *RenderOpts) error {
	tmpl, err := template.ParseFS(templates, templatesGlob)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(writer, templatePlotJS, opts)
}

// WriteJSON uses the given Triangle and RenderOpts to render a JSON representation
// of the Triangle and write it to the given io.Writer.
func WriteJSON(writer io.Writer, opts *RenderOpts) error {
	jsonBytes, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	_, err = writer.Write(jsonBytes)
	return err
}
