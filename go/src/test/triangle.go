package main

import "math"

func calcTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}
