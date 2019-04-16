package main

import (
	"fmt"
)

func main() {

	rover1 := &Rover{
		x:         1,
		y:         2,
		direction: compass["N"],
	}

	rover1.command("LMLMLMLMM")

	rover2 := &Rover{
		x:         3,
		y:         3,
		direction: compass["E"],
	}

	rover2.command("MMRMMRMRRM")

	fmt.Println(rover1.state())
	fmt.Println(rover2.state())
}
