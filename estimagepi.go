/*
 *Prompt: Using only basic arithmetic and a random number generator approximate
 *Pi.
 */
package main

import (
	"fmt"
	mr "math/rand"
	t "time"
)

func main() {
	var num int
	fmt.Print("Input how many points you would like to use: ")
	fmt.Scan(&num)
	fmt.Println(estimatePi(num))
}

func estimatePi(n int) float64 {
	var numPointsInCircle float64 = 0
	var numPointsInSquare float64 = float64(n)
	mr.Seed(t.Now().UnixNano())
	for i := 0; i < n; i++ {
		pointX := mr.Float32()
		pointY := mr.Float32()
		distance := (pointX * pointX) + (pointY * pointY)

		if distance < 1 {
			numPointsInCircle++
		}
	}
	return 4 * (numPointsInCircle / numPointsInSquare)
}
