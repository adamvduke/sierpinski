// Package web provides utilities for drawing a sierpinski triangle using html
// and javascript.
package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adamvduke/sierpinski/chaos"
	"github.com/adamvduke/sierpinski/config"
	"github.com/adamvduke/sierpinski/static"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	listenAddr = "localhost:8080"
)

// Serve starts an http server which provides visualizations of a sierpinski
// triangle.
func Serve() {
	mux := chi.NewMux()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	publicFS, err := static.PublicFS()
	if err != nil {
		panic(err)
	}
	ph, ok := http.StripPrefix("/public/", http.FileServerFS(publicFS)).(http.HandlerFunc)
	if ok {
		mux.Get("/public/*", ph)
	}
	mux.Get("/", writeHTML)
	mux.Get("/plot.js", writeJS)
	mux.Get("/data.json", writeJSON)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}

func writeHTML(w http.ResponseWriter, r *http.Request) {
	opts := getOptions(r)
	if err := WriteHTML(w, opts); err != nil {
		panic(err)
	}
}

func writeJS(w http.ResponseWriter, r *http.Request) {
	opts := getOptions(r)
	if err := WriteJS(w, opts); err != nil {
		panic(err)
	}
}

func writeJSON(w http.ResponseWriter, r *http.Request) {
	opts := getOptions(r)
	if err := WriteJSON(w, opts); err != nil {
		panic(err)
	}
}

func getOptions(req *http.Request) *RenderOpts {
	size := intParam(req, "size", config.DefaultBaseLength)
	count := intParam(req, "count", config.DefaultPointCount)
	radius := floatParam(req, "radius", config.DefaultPointRadius)
	red := intParam(req, "r", config.DefaultPointRed)
	green := intParam(req, "g", config.DefaultPointGreen)
	blue := intParam(req, "b", config.DefaultPointBlue)
	opts := &RenderOpts{
		Chaos:       chaos.New(size, count),
		PointRadius: radius,
		PointColor: &RGB{
			R: red,
			G: green,
			B: blue,
		},
	}
	return opts
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
