package main

import (
	"fmt"
	"math"
)

type floatSlice []float64

func (slice floatSlice) pos(value float64) float64 {
	for _, v := range slice {
		if v == value {
			return 1.0
		}
	}
	return -1.0
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
			for y := 0.0; y < a; y++ {
				if len(result) > b {
					break
				}

				modster := math.Mod(math.Pow(x, y), a)
				if result.pos(modster) != -1.0 {
					break
				}
				// length of result is > 54 its not a primitive root
				// if modster is already in result than it's not a primitive root
				result = append(result, math.Mod(math.Pow(x, y), a))
			}

			if len(result) == b {
				fmt.Printf("Primitive Root: %v \n", x)
			}

		}

	}

}
