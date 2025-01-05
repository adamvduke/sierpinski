// Package web provides utilities for drawing a sierpinski triangle using html
// and javascript.
package web

import (
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/adamvduke/sierpinski/chaos"
	"github.com/adamvduke/sierpinski/config"
)

const (
	listenAddr = "localhost:8080"
)

// Serve starts an http server that serves a single route, "/", which provides
// an html visualization of a sierpinski triangle.
func Serve(open bool) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		size := intParam(r, "size", config.DefaultBaseLength)
		count := intParam(r, "count", config.DefaultPointCount)
		radius := floatParam(r, "radius", config.DefaultPointRadius)
		t := chaos.New(size, count)
		opts := &RenderOpts{
			PointRadius: radius,
			PointColor: &RGB{
				R: config.DefaultPointRed,
				G: config.DefaultPointGreen,
				B: config.DefaultPointBlue,
			},
		}
		if err := Write(t, w, opts); err != nil {
			panic(err)
		}
	})
	go maybeOpen(open)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func maybeOpen(open bool) error {
	time.Sleep(2 * time.Second)
	if open {
		if err := exec.Command("open", "http://"+listenAddr).Run(); err != nil {
			return err
		}
	}
	return nil
}

func intParam(r *http.Request, name string, defaultValue int) int {
	v, err := strconv.Atoi(r.URL.Query().Get(name))
	if err != nil {
		v = defaultValue
	}
	return v
}

func floatParam(r *http.Request, name string, defaultValue float64) float64 {
	v, err := strconv.ParseFloat(r.URL.Query().Get(name), 64)
	if err != nil {
		v = defaultValue
	}
	return v
}
