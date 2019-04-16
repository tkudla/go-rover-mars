package main

import "fmt"

type Rover struct {
	x         int
	y         int
	direction int
}

func (rover *Rover) left() {

	if rover.direction == compass["N"] {
		rover.direction = compass["W"]
	} else {
		rover.direction--
	}
}

func (rover *Rover) right() {

	if rover.direction == compass["W"] {
		rover.direction = compass["N"]
	} else {
		rover.direction++
	}
}

func (rover *Rover) move() {

	switch rover.direction {

	case compass["N"]:
		rover.y++
		break
	case compass["E"]:
		rover.x++
		break
	case compass["S"]:
		rover.y--
		break
	case compass["W"]:
		rover.x--
		break
	}
}

func (rover *Rover) command(cmd string) {

	for i := 0; i < len(cmd); i++ {
		char := string(cmd[i])

		switch char {
		case "M":
			rover.move()
			break
		case "L":
			rover.left()
			break
		case "R":
			rover.right()
			break
		}
	}
}

func (rover *Rover) state() string {

	return fmt.Sprintf("%d %d %d", rover.x, rover.y, compassName(rover.direction))
}
