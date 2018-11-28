package main

import "fmt"

func fact(n int) float64 {
	res := 1.
	for j := n; j > 1; j-- {
		res *= float64(j)
	}
	return res
}

func backwardDiffTable(x []float64, y [][]float64) {
	n := len(x)
	for i := 1; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			y[j][i] = y[j][i-1] - y[j-1][i-1]
		}
	}
}

func PrintBackwardTableDiv(y [][]float64) {
	n := len(y)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("% 4.2f\t", y[i][j])
		}
		fmt.Printf("\n")
	}
}

func polByBackWithValue(y [][]float64, x []float64) func(float64) float64 {
	n := len(x)
	return func(value float64) float64 {
		value = -x[n-1]
		res := y[n-1][0]
		n := len(y)
		for i := 1; i < n; i++ {
			t := y[n-1][i]
			for j := 0; j < i; j++ {
				t *= value + float64(j)
			}
			res += t / fact(i)
		}
		return res
	}
}
