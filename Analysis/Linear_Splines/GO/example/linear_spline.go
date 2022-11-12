package example

import (
	"errors"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math"
	"math/rand"
	"strconv"
)

var argTable = [...]float64{0, 3, 5, 7.5, 8, 11, 14, 18, 20, 22}
var valTable = [...]float64{80, 90, 105, 140, 131, 162, 197, 251, 280, 310}

func getValLinear(arg float64) (float64, error) {
	if arg < 0 {
		return 0, errors.New("arg too small")
	}
	if arg > argTable[len(argTable)-1] {
		return 0, errors.New("arg too big")
	}

	var index int
	for index = 0; index < len(argTable); index++ {
		if EqualTolerance(arg, argTable[index], 0.1) {
			return valTable[index], nil
		}
		if arg > argTable[index] && arg < argTable[index+1] {
			break
		}
	}

	x1, x2 := argTable[index], argTable[index+1]
	y1, y2 := valTable[index], valTable[index+1]

	a := (y2 - y1) / (x2 - x1)
	b := y1 - (a * x1)

	return a*arg + b, nil
}

func EqualTolerance(a, b, tolerance float64) bool {
	tol := tolerance
	if diff := math.Abs(a - b); diff < tol {
		return true
	}
	return false
}

func PlotValLinear(path string, start float64, end float64, step float64) error {
	if start < float64(0) {
		return errors.New("Argument Error: start is not valid")
	}
	if start >= end {
		return errors.New("Argument Error: end is not greater than start")
	}
	if step <= float64(0) {

		return errors.New("Argument Error: step is not valid " + fmt.Sprintf("%f", step))
	}

	count := 0
	for i := start; true; i += step {
		count++
		if EqualTolerance(i, end, 0.1) {
			break
		}
		if i > end {
			return errors.New("Argument Error: step is not valid" +
				fmt.Sprintf(" %f %d %f", step, count, i))
		}
	}

	fmt.Println("start to end with step needs: " + strconv.Itoa(count) + " steps")

	p := plot.New()
	p.Title.Text = "Fuel Per Hour"
	p.X.Label.Text = "Speed"
	p.Y.Label.Text = "Fuel"

	pts := make(plotter.XYs, count)
	var currX float64 = 0
	var currY float64 = 0
	var err error
	for i, _ := range pts {
		currX = (float64(i) * step) + start
		pts[i].X = currX
		currY, err = getValLinear(currX)
		pts[i].Y = currY
		if err != nil {
			fmt.Println("Error getValLinear")
			fmt.Println(err)
			continue
		}
	}

	refPts := make(plotter.XYs, len(argTable))
	for i, _ := range argTable {
		refPts[i].X = argTable[i]
		refPts[i].Y = valTable[i]
	}

	err = nil
	/*
		err = plotutil.AddLinePoints(p,
			"öppis", pts,
			"Second", randomPoints(15))
	*/
	err = plotutil.AddLinePoints(p,
		"Linear Splines", pts,
		"Ref Points", refPts)

	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, path); err != nil {
		panic(err)
	}

	return nil
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
