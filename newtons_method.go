//Oskar Enmalm
//04.01.17

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 { //need a limit
	//x needs to dynamic, hence we don't declare it
	var z float64 = 2
	var n float64 = 1
	var limit float64 = 10

	for ; n < limit; n++ {
		z = (z - ((z*z) - x) / (2 * z))
		println(z)
		
	}
		return z
}

func main() {
	fmt.Println(Sqrt(81))
}
