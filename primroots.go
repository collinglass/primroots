package main

import (
	"fmt"
	"math"
)

type floatSlice []float64

func (slice floatSlice) pos(value float64) int {
	for k, v := range slice {
		if v == value {
			return k
		}
	}
	return -1
}

func gcd(a, b float64) float64 {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func main() {
	a := 81.0
	b := 54

	for x := 2.0; x < a; x++ {
		if gcd(x, a) == 1.0 {
			result := make(floatSlice, 1, 81)
			result[0] = 1
			r := 1.0

			for y := 1.0; y < a; y++ {
				r = math.Mod(r*x, a)
				// length of result is > 54 its not a primitive root
				if len(result) > b {
					break
				}

				// if (x**y mod a) is already in result than it's not a primitive root
				if result.pos(r) != -1 {
					break
				}
				result = append(result, r)

				if len(result) == b {
					fmt.Printf("Primitive Root: %v \n", x)
				}
			}
		}

	}

}
