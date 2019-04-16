package main

import (
	"fmt"
)

func main() {

	plateau := Plateau{
		x: 5,
		y: 5,
	}

	rover1 := Rover{
		x:         1,
		y:         2,
		direction: compass["N"],
	}

	rover1.command("LMLMLMLMM")

	plateau.place(rover1)

	rover2 := Rover{
		x:         3,
		y:         3,
		direction: compass["E"],
	}

	rover2.command("MMRMMRMRRM")

	plateau.place(rover2)

	fmt.Println(rover1.state())
	fmt.Println(rover2.state())
}
