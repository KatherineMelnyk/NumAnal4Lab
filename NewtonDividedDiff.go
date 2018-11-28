package main

import (
	"fmt"
	"strings"
)

func dividedDiffTable(x []float64, y [][]float64) {
	n := len(x)
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			y[j][i] = (y[j][i-1] - y[j+1][i-1]) / (x[j] - x[i+j])
		}
	}
}

func PrintTableDiv(y [][]float64) {
	n := len(y)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("% 3.3f\t", y[i][j])
		}
		fmt.Printf("\n")
	}
}

func polWithValue(y [][]float64, x []float64) func(float64) float64 {
	return func(value float64) float64 {
		res := y[0][0]
		n := len(y)
		for i := 1; i < n; i++ {
			t := y[0][i]
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
