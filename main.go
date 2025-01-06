// Sierpinski is a small program that provides visualization(s) of a sierpinski
// triangle.
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/adamvduke/sierpinski/chaos"
	"github.com/adamvduke/sierpinski/config"
	"github.com/adamvduke/sierpinski/web"
)

const (
	visualizeHTML = "visualize.html"
)

func main() {
	serve := flag.Bool("serve", false, "whether the program should start a server or just render the template on disk")
	open := flag.Bool("open", false, "whether the program should open the visualization when ready")
	size := flag.Int("size", config.DefaultBaseLength, "length of a single side of the triangle")
	count := flag.Int("count", config.DefaultPointCount, "number of points to generate")
	radius := flag.Float64("radius", config.DefaultPointRadius, "size of points when graphed")
	r := flag.Int("r", config.DefaultPointRed, "r value for creating an RGB value")
	g := flag.Int("g", config.DefaultPointGreen, "g value for creating an RGB value")
	b := flag.Int("b", config.DefaultPointBlue, "b value for creating an RGB value")

	flag.Parse()

	if *serve {
		// blocks forever
		web.Serve()
	}

	renderOpts := &web.RenderOpts{
		Chaos:       chaos.New(*size, *count),
		PointRadius: *radius,
		PointColor: &web.RGB{
			R: *r,
			G: *g,
			B: *b,
		},
	}

	// writes the rendered file to disk
	if err := writeFile(renderOpts); err != nil {
		log.Fatal(err)
	}

	if *open {
		if err := exec.Command("open", visualizeHTML).Run(); err != nil {
			log.Fatal(err)
		}
	}
}

// writeFile uses the given parameters to create a chaos.Triangle write an html
// visualization to a file.
func writeFile(opts *web.RenderOpts) error {
	outFile, err := os.Create(visualizeHTML)
	if err != nil {
		return err
	}
	if err := web.WriteHTML(outFile, opts); err != nil {
		return err
	}
	return nil
}
