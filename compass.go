package main

var compass = map[string]int{
	"N": 1,
	"E": 2,
	"S": 3,
	"W": 4,
}

func compassName(direction int) string {

	for k, v := range compass {

		if v == direction {
			return k
		}
	}
	return ""
}
