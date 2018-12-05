package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const a = -math.Pi

//const a = 2 * math.Pi / 3
//const a = 0
const b = math.Pi

func secCheb(a, b float64, n int) []float64 {
	var x []float64
	c1 := float64((a + b) / 2)
	c2 := float64((b - a) / 2)
	c3 := float64(2 * n)
	for k := 0; k < n; k++ {
		x = append(x, c1+c2*math.Cos(math.Pi*float64(2*k+1)/c3))
	}
	return x
}

func secEqualRem(a, b float64, n int) []float64 {
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

func yFromX(x []float64, f func(float64) float64) []float64 {
	y := make([]float64, len(x))
	for i, v := range x {
		y[i] = f(v)
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

func hand(x float64) float64 {
	return 1.77 + (x+1.)*(-0.5) + (x+1.)*(x-0.17)*(-0.13) + (x+1.)*(x-0.17)*(x-0.93)*5.6
}

func main() {
	fmt.Println(0, math.Sqrt(math.Pi))
	x0 := secEqualRem(0, math.Sqrt(math.Pi), 4)
	y0 := yFromX(x0, f)

	for i := 0; i < len(x0)/2; i++ {
		x0[i], x0[len(x0)-1-i] = x0[len(x0)-1-i], x0[i]
		y0[i], y0[len(y0)-1-i] = y0[len(y0)-1-i], y0[i]
	}

	//x0 = []float64{0.880, 0.881, 0.882, 0.883}
	//y0 = []float64{2.4109, 2.4133, 2.4157, 2.41181}

	fmt.Printf(" \n")

	fmt.Println(x0)
	fmt.Println(y0)

	yz := 1. / 3
	//yz = 2.4142
	intF := BuildInterpolationFunction(y0, x0)
	xz := intF(yz)

	fmt.Printf("Q(%.2f)  = %.2f\n", yz, xz)

	fmt.Printf("f(x) - f(x*)  = %3.5f\n", yz-f(xz))

	fmt.Print("\n")

	xR := []float64{1.69, 1.73, 1.77, 1.82, 1.85}
	yR := yFromX(xR, f)

	intR := BuildInterpolationFunction(yR, xR)
	xr := intR(yz)

	fmt.Printf("Q(%.2f)  = %.2f\n", yz, xr)

	fmt.Printf("f(x) - f(x*)  = %3.5f\n", yz-f(xr))

	fmt.Print("\n")

	x1 := secCheb(a, b, 15)
	y1 := yFromX(x1, f)
	f1 := BuildInterpolationFunction(x1, y1)

	x2 := secEqualRem(a, b, 15)
	y2 := yFromX(x2, f)
	f2 := BuildInterpolationFunction(x2, y2)

	fPol1 := plotter.NewFunction(f1)
	fPol1.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	fPol1.Width = vg.Inch / 20
	fPol1.Samples = 200

	fPol2 := plotter.NewFunction(f2)
	fPol2.Color = color.RGBA{R: 15, G: 93, B: 209, A: 200}
	fPol2.Width = vg.Inch / 20
	fPol2.Samples = 200

	fInt := plotter.NewFunction(intF)
	fInt.Color = color.RGBA{R: 255, G: 0, B: 0, A: 200}
	fInt.Width = vg.Inch / 20
	fInt.Samples = 200

	fIntArk := plotter.NewFunction(func(x float64) float64 {
		if x < -1 || x > 1 {
			return 0
		}
		return math.Sqrt(math.Acos(x))
	})
	fIntArk.Color = color.RGBA{R: 0, G: 0, B: 255, A: 200}
	fIntArk.Width = vg.Inch / 20
	fIntArk.Samples = 200

	fPolh := plotter.NewFunction(hand)
	fPolh.Color = color.RGBA{R: 65, G: 103, B: 187, A: 200}
	fPolh.Width = vg.Inch / 20
	fPolh.Samples = 200

	//fPolsub := plotter.NewFunction(sub(f, f1))
	//fPolsub.Color = color.RGBA{R: 190, G: 15, B: 209, A: 200}
	//fPolsub.Width = vg.Inch / 20
	//fPolsub.Samples = 200

	fPolOrig := plotter.NewFunction(f)
	fPolOrig.Color = color.RGBA{R: 200, G: 150, B: 25, A: 200}
	fPolOrig.Width = vg.Inch / 20
	fPolOrig.Samples = 200

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = a-0.5, b+0.5
	pl.Y.Min, pl.Y.Max = -2, 2
	pl.Add(fPol1)
	pl.Add(fPol2)
	pl.Add(fPolOrig)
	//pl.Add(fInt)
	//pl.Add(fIntArk)
	//pl.Add(fPolh)
	//pl.Add(fPolsub)

	pl.Title.Text = "Interpolation"
	pl.Title.Font.Size = vg.Inch
	pl.Legend.Font.Size = vg.Inch / 2
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 2
	pl.Legend.Add("f(x)", fPolOrig)
	pl.Legend.Add("CoefCheb", fPol1)
	pl.Legend.Add("CoefEqual", fPol2)
	//pl.Legend.Add("f(x)-Pn(x)", fPolsub)

	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "pol.png"); err != nil {
		panic(err.Error())

	}
}
