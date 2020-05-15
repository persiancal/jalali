package main

import (
	"fmt"
	"math"
)

const leapPart = 0.24219879

func main() {
	for i := 0; i <= 282000; i++ {
		d := float64(i) * leapPart
		//		fmt.Println(float64(i) * leapPart)
		if d-math.Floor(d) < 0.00001 {
			fmt.Println(i, "--", i%2820, "=>", d)
		}
	}
}
