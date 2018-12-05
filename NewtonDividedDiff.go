package main

import (
	"fmt"
	"strings"
)

func BuildInterpolationFunction(x, y []float64) func(float64) float64 {
	var divTable [][]float64
	n := len(x)

	for i := 0; i < n; i++ {
		divTable = append(divTable, []float64{})
		for j := 0; j < n; j++ {
			if j == 0 {
				divTable[i] = append(divTable[i], y[i])
			} else {
				divTable[i] = append(divTable[i], 0)
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			divTable[j][i] = (divTable[j+1][i-1] - divTable[j][i-1]) / (x[i+j] - x[j])
		}
	}

	return func(value float64) float64 {
		res := divTable[0][0]
		n := len(y)
		for i := 1; i < n; i++ {
			t := divTable[0][i]
			for j := 0; j < i; j++ {
				t *= value - x[j]
			}
			res += t
		}
		return res
	}

}

func printPol(y [][]float64, x []float64) {
	var coefficient []float64
	for i := 0; i < len(x); i++ {
		coefficient = append(coefficient, y[0][i])
	}
	var parts []string
	for i, a := range coefficient {
		var part string
		if i == 0 {
			part = fmt.Sprintf("%.2f", a)
		} else {
			part += fmt.Sprintf("%.2f", a)
			for j := 0; j < i; j++ {
				if x[j] < 0 {
					part += fmt.Sprintf("*(x+%.3f)", (-1)*x[j])
				} else {
					part += fmt.Sprintf("*(x-%.3f)", x[j])
				}
			}
		}
		parts = append(parts, part)
	}
	fmt.Println(strings.Join(parts, "+"))
}
