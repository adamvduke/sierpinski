package web

import (
	"embed"
	"html/template"
	"io"
	"os"

	"github.com/adamvduke/sierpinski/chaos"
)

const (
	templateName  = "visualize.html.tmpl"
	templatesGlob = "templates/*.html.tmpl"
	visualizeHTML = "visualize.html"
)

//go:embed templates
var templates embed.FS

// RGB represents a color used to draw a sierpinski triangle.
type RGB struct {
	R, G, B int
}

// RenderOpts represents the options available when rendering a sierpinski
// triangle.
type RenderOpts struct {
	PointRadius float64
	PointColor  *RGB
}

// WriteFile uses the given parameters to create a chaos.Triangle write an html
// visualization to a file.
func WriteFile(opts *chaos.Opts, renderOpts *RenderOpts) (string, error) {
	t := chaos.New(opts.Size, opts.PointCount)
	if err := writeFile(t, visualizeHTML, renderOpts); err != nil {
		return "", err
	}
	return visualizeHTML, nil
}

// writeFile uses the given Triangle and RenderOpts to render an html template
// and write it to the given file.
func writeFile(t *chaos.Triangle, outputFile string, opts *RenderOpts) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	return Write(t, outFile, opts)
}

// Write uses the given Triangle and RenderOpts to render and html template
// and write it to the given io.Writer.
func Write(t *chaos.Triangle, writer io.Writer, opts *RenderOpts) error {
	points := make([]template.JS, len(t.Points))
	for idx, p := range t.Points {
		points[idx] = template.JS(p.String())
	}
	tmplData := struct {
		Points         []template.JS
		GraphDimension int
		PointRadius    float64
		R, G, B        int
	}{
		Points:         points,
		GraphDimension: t.SideLength,
		PointRadius:    opts.PointRadius,
		R:              opts.PointColor.R,
		B:              opts.PointColor.B,
		G:              opts.PointColor.G,
	}
	tmpl, err := template.ParseFS(templates, templatesGlob)
	if err != nil {
		return err
	}
	if err := tmpl.ExecuteTemplate(writer, templateName, tmplData); err != nil {
		return err
	}
	return nil
}
