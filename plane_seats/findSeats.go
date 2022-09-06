package main

import "log"

func findSeats() {
	for _, group := range groups {
		var rowSeats [6]bool
		var row int
		for row, rowSeats = range planeSeats {
			var start int
			var backwards bool
			switch {
			case group.side == "left" && group.position == "aisle":
				start = 2
				backwards = true
			case group.side == "left" && group.position == "window":
				start = 0
			case group.side == "right" && group.position == "aisle":
				start = 3
			case group.side == "right" && group.position == "window":
				start = 5
				backwards = true
			default:
				log.Fatalln("Erroneous input.")
			}

			if !backwards {

			} else {

			}
		}

	}
}
