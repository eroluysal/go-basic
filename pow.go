package main

import "fmt"
import "math"

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return lim
	}
}

func main() {
	fmt.Println(pow(2, 4, 10), pow(3, 5, 8))
}
