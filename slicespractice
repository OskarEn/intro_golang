/*This was an exercise on slices with the GoBridge group
Thanks to Peter for walking me through everything so well
You put an input and will display the output in blueshades
*/

package main

import "golang.org/x/tour/pic"


//the inputs in integer form, the outpt matrix as uint8
func Pic(dx, dy int) [][]uint8 { 
	frame := make([][]uint8, dy)//creatig the frame with length dy
	for y := 0; y < dy; y++ {
		frame[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			frame[y][x] = uint8()//any equation within the parenthesis
		}

	}
	return frame
}

func main() {
	pic.Show(Pic)
}
