package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const a = -1 * math.Pi
const b = math.Pi

func secCheb(n int) []float64 {
	var x []float64
	c1 := float64((a + b) / 2)
	c2 := float64((b - a) / 2)
	c3 := float64(2 * n)
	for k := 0; k < n; k++ {
		x = append(x, c1+c2*math.Cos(math.Pi*float64(2*k+1)/c3))
	}
	return x
}

func secEqualRem(n int) []float64 {
	var x []float64
	h := float64(b-a) / float64(n-1)
	for k := 0; k < n; k++ {
		x = append(x, a+float64(k)*h)
	}
	return x
}

func f(x float64) float64 {
	x = math.Pow(x, 2)
	return math.Cos(x)
}

func y(x []float64) [][]float64 {
	var y [][]float64
	n := len(x)
	for i := 0; i < n; i++ {
		y = append(y, []float64{})
		for j := 0; j < n; j++ {
			if j == 0 {
				y[i] = append(y[i], f(x[i]))
			} else {
				y[i] = append(y[i], 0)
			}
		}
	}
	return y
}

func prints(x []float64) {
	for i := 0; i < len(x); i++ {
		fmt.Printf("% 4.4f\n", x[i])
	}
}

func sub(a func(x float64) float64, b func(y float64) float64) func(float64) float64 {
	return func(value float64) float64 {
		return a(value) - b(value)
	}
}

func main() {

	x1 := secEqualRem(15)
	y1 := y(x1)
	f1 := polWithValue(y1, x1)
	dividedDiffTable(x1, y1)
	printPol(y1, x1)
	//PrintTableDiv(y1)

	fmt.Print("\n")

	xb := secEqualRem(7)
	yb := y(xb)
	backwardDiffTable(xb, yb)
	//PrintBackwardTableDiv(yb)
	fmt.Printf("f(x*) = %3.5f\n", polByBackWithValue(yb, xb)(0.))
	//fmt.Printf("%3.3f\n")

	fmt.Printf("f(x)  = %3.5f\n", math.Cos(math.Pow(0, 2)))

	fmt.Printf("f(x) - f(x*)  = %3.5f\n", math.Cos(math.Pow(0, 2))-polByBackWithValue(yb, xb)(0.))

	fmt.Print("\n")

	x2 := secEqualRem(15)
	y2 := y(x2)
	f2 := polWithValue(y2, x2)
	dividedDiffTable(x2, y2)
	printPol(y2, x2)
	//PrintTableDiv(y2)

	fPol1 := plotter.NewFunction(f1)
	fPol1.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	fPol1.Width = vg.Inch / 20
	fPol1.Samples = 200

	fPol2 := plotter.NewFunction(f2)
	fPol2.Color = color.RGBA{R: 15, G: 93, B: 209, A: 200}
	fPol2.Width = vg.Inch / 20
	fPol2.Samples = 200

	fPolsub := plotter.NewFunction(sub(func(x float64) float64 {
		return math.Cos(x * x)
	}, f1))
	fPolsub.Color = color.RGBA{R: 190, G: 15, B: 209, A: 200}
	fPolsub.Width = vg.Inch / 20
	fPolsub.Samples = 200

	PolW := plotter.NewFunction(func(x float64) float64 {
		res := 1.
		for i := 0; i < len(x1); i++ {
			res *= x - x1[i]
		}
		return res
	})
	PolW.Color = color.RGBA{R: 15, G: 209, B: 93, A: 200}
	PolW.Width = vg.Inch / 20
	PolW.Samples = 200

	fPolOrig := plotter.NewFunction(func(x float64) float64 {
		return math.Cos(x * x)
	})
	fPolOrig.Color = color.RGBA{R: 200, G: 150, B: 25, A: 200}
	fPolOrig.Width = vg.Inch / 20
	fPolOrig.Samples = 200

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = a-0.5, b+0.5
	pl.Y.Min, pl.Y.Max = -2, 2
	pl.Add(fPol1)
	pl.Add(fPol2)
	pl.Add(fPolOrig)
	pl.Add(fPolsub)
	pl.Add(PolW)

	pl.Title.Text = "Interpolation"
	pl.Title.Font.Size = vg.Inch
	pl.Legend.Font.Size = vg.Inch / 2
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 2
	pl.Legend.Add("f(x)", fPolOrig)
	pl.Legend.Add("CoefCheb", fPol1)
	pl.Legend.Add("CoefEqual", fPol2)
	pl.Legend.Add("f(x)-Pn(x)", fPolsub)
	pl.Legend.Add("W(n)", PolW)

	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "pol.png"); err != nil {
		panic(err.Error())
	}

}
