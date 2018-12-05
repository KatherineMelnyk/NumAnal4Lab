package main

//import (
//	"image/color"
//	"math"
//)
//
//	x1 := secEqualRem(15)
//	y1 := y(x1)
//	f1 := polWithValue(y1, x1)
//	dividedDiffTable(x1, y1)
//	printPol(y1, x1)
//
//	x2 := secEqualRem(15)
//	y2 := y(x2)
//	f2 := polWithValue(y2, x2)
//	dividedDiffTable(x2, y2)
//	printPol(y2, x2)
//
//	fPol1 := plotter.NewFunction(f1)
//	fPol1.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
//	fPol1.Width = vg.Inch / 20
//	fPol1.Samples = 200
//
//	fPol2 := plotter.NewFunction(f2)
//	fPol2.Color = color.RGBA{R: 15, G: 93, B: 209, A: 200}
//	fPol2.Width = vg.Inch / 20
//	fPol2.Samples = 200
//
//	fPolsub := plotter.NewFunction(sub(func(x float64) float64 {
//		return math.Cos(x * x)
//	}, f1))
//	fPolsub.Color = color.RGBA{R: 190, G: 15, B: 209, A: 200}
//	fPolsub.Width = vg.Inch / 20
//	fPolsub.Samples = 200
//
//	fPolOrig := plotter.NewFunction(func(x float64) float64 {
//		return math.Cos(x * x)
//	})
//	fPolOrig.Color = color.RGBA{R: 200, G: 150, B: 25, A: 200}
//	fPolOrig.Width = vg.Inch / 20
//	fPolOrig.Samples = 200
//
//	pl, _ := plot.New()
//	pl.X.Min, pl.X.Max = a-0.5, b+0.5
//	pl.Y.Min, pl.Y.Max = -2, 2
//	pl.Add(fPol1)
//	pl.Add(fPol2)
//	pl.Add(fPolOrig)
//	pl.Add(fPolsub)
//	pl.Add(PolW)
//
//	pl.Title.Text = "Interpolation"
//	pl.Title.Font.Size = vg.Inch
//	pl.Legend.Font.Size = vg.Inch / 2
//	pl.Legend.XOffs = -vg.Inch
//	pl.Legend.YOffs = vg.Inch / 2
//	pl.Legend.Add("f(x)", fPolOrig)
//	pl.Legend.Add("CoefCheb", fPol1)
//	pl.Legend.Add("CoefEqual", fPol2)
//	pl.Legend.Add("f(x)-Pn(x)", fPolsub)
//	pl.Legend.Add("W(n)", PolW)
//
//	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "pol.png"); err != nil {
//		panic(err.Error())
