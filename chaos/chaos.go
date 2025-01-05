// Package chaos implements a sierpinski triangle via the "chaos game".
package chaos

import (
	"fmt"
	"math"
	"math/rand"
)

// Point represents a single point on a graph.
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// String provides a formatted string representation of a Point.
func (p *Point) String() string {
	return fmt.Sprintf("{x: %f, y: %f}", p.X, p.Y)
}

func (p *Point) midPoint(other *Point) *Point {
	x := (p.X + other.X) / 2
	y := (p.Y + other.Y) / 2
	return &Point{X: x, Y: y}
}

// Triangle implements a sierpinski triangle via the chaos game.
type Triangle struct {
	SideLength int      `json:"length"`
	PointCount int      `json:"point_count"`
	Points     []*Point `json:"points"`

	xMid   float64
	height float64
	a      *Point
	b      *Point
	c      *Point
}

// Opts provides a data carrier for specifying the size of a Triangle and the
// number of Points to create.
type Opts struct {
	Size, PointCount int
}

// New creates a Triangle with the given length and number of points.
func New(sideLength, pointCount int) *Triangle {
	xMid := float64(sideLength) / 2.0
	fSideLength := float64(sideLength)
	height := math.Sqrt(fSideLength*fSideLength - xMid*xMid)

	t := &Triangle{
		a:          &Point{X: 0.0, Y: 0.0},
		b:          &Point{X: xMid, Y: height},
		c:          &Point{X: fSideLength, Y: 0.0},
		Points:     []*Point{},
		SideLength: sideLength,
		PointCount: pointCount,
		height:     height,
		xMid:       xMid,
	}
	t.GeneratePoints()
	return t
}

// GeneratePoints populates the Triangle's set of graphable Points.
func (s *Triangle) GeneratePoints() []*Point {
	graphPoints := make([]*Point, s.PointCount)
	current := s.randomPoint()
	vertex := s.randomVertex()

	for i := 0; i < s.PointCount; i++ {
		current = current.midPoint(vertex)
		vertex = s.randomVertex()
		graphPoints[i] = current
	}
	s.Points = graphPoints
	return graphPoints
}

func (s *Triangle) randomPoint() *Point {
	r := rand.Float64()
	x := r * s.xMid
	slope := s.height / s.xMid
	yMax := x * slope
	y := rand.Float64() * yMax

	// flip half the points over the x center line
	if rand.Float64() > 0.5 {
		toCenter := s.xMid - x
		x = x + (2 * toCenter)
	}
	return &Point{X: x, Y: y}
}

func (s *Triangle) randomVertex() *Point {
	idx := rand.Intn(3)
	return []*Point{s.a, s.b, s.c}[idx]
}
