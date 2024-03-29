package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
)

func lineAndPoints() {
	p := plot.New()

	p.Title.Text = "Hello Price"
	p.X.Label.Text = "Quantity Demand"
	p.Y.Label.Text = "Price"

	points := plotter.XYs{
		{2.0, 60000.0},
		{4.0, 40000.0},
		{6.0, 30000.0},
		{8.0, 25000.0},
		{10.0, 23000.0},
	}

	plotutil.AddLinePoints(p, points)

	p.Save(4*vg.Inch, 4*vg.Inch, "price.png")
}

func histograms() {
	p := plot.New()
	p.Title.Text = "Histogram"

	bins := plotter.XYs{
		{10, 10},
		{20, 20},
		{30, 50},
		{40, 20},
		{50, 10},
	}

	h, _ := plotter.NewHistogram(bins, 5)
	p.Add(h)

	p.Save(4*vg.Inch, 4*vg.Inch, "histogram.png")
}

type Squares struct {
	plotter.XYs
}

func (s *Squares) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)

	c.SetColor(color.RGBA{R: 196, B: 128, A: 255})

	r := vg.Length(10.0)
	for _, p := range s.XYs {
		p1 := vg.Point{trX(p.X) - r, trY(p.Y) - r}
		p2 := vg.Point{trX(p.X) - r, trY(p.Y) + r}
		p3 := vg.Point{trX(p.X) + r, trY(p.Y) + r}
		p4 := vg.Point{trX(p.X) + r, trY(p.Y) - r}

		var p vg.Path
		p.Move(p1)
		p.Line(p2)
		p.Line(p3)
		p.Line(p4)
		p.Line(p1)
		p.Close()
		c.Fill(p)
	}
}

func main() {
	points := plotter.XYs{
		{2, 2},
		{4, 4},
		{6, 6},
		{8, 8},
		{10, 10},
	}

	s := Squares{points}

	p := plot.New()
	p.Title.Text = "Squares"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.X.Min = 0
	p.X.Max = 20
	p.Y.Min = 0
	p.Y.Max = 20

	p.Add(&s)
	p.Save(4*vg.Inch, 4*vg.Inch, "squares.png")
}

func (s *Squares) DataRange() (float64, float64, float64, float64) {
	return plotter.XYRange(s.XYs)
}

func (s *Squares) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	boxes := make([]plot.GlyphBox, len(s.XYs))

	r := vg.Length(10.0)
	for i, p := range s.XYs {
		boxes[i].X = plt.X.Norm(p.X)
		boxes[i].Y = plt.Y.Norm(p.Y)
		boxes[i].Rectangle = vg.Rectangle{
			Min: vg.Point{X: -r, Y: -r},
			Max: vg.Point{X: +r, Y: +r},
		}
	}
	return boxes
}