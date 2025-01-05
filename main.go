// Sierpinski is a small program that provides visualization(s) of a sierpinski
// triangle.
package main

import (
	"flag"
	"log"
	"os/exec"

	"github.com/adamvduke/sierpinski/chaos"
	"github.com/adamvduke/sierpinski/config"
	"github.com/adamvduke/sierpinski/web"
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
		web.Serve(*open)
	}

	opts := &chaos.Opts{Size: *size, PointCount: *count}
	renderOpts := &web.RenderOpts{
		PointRadius: *radius,
		PointColor: &web.RGB{
			R: *r,
			G: *g,
			B: *b,
		},
	}

	// writes the rendered file to disk and returns the path to the file
	path, err := web.WriteFile(opts, renderOpts)
	if err != nil {
		log.Fatal(err)
	}

	if *open {
		if err := exec.Command("open", path).Run(); err != nil {
			log.Fatal(err)
		}
	}
}
