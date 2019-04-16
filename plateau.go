package main

type Plateau struct {
	x      int
	y      int
	rovers []Rover
}

func (plateau *Plateau) place(rover Rover) {

	if rover.x > plateau.x || rover.y > plateau.y || rover.x < 0 || rover.y < 0 {
		panic("Rover is outside the plateau")
	}

	for _, existing := range plateau.rovers {

		if &existing != &rover {

			if existing.x == rover.x && existing.y == rover.y {
				panic("Another rover placed in this cell")
			}
		}
	}

	plateau.rovers = append(plateau.rovers, rover)
}
