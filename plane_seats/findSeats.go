package main

import (
	"fmt"
	"log"
)

func findSeats() {
outer:
	for _, group := range groups {
		var row int
		seatsAvailable := true

		start, finish := startDirection(group)

		for row = range planeSeats {

			for i := start; i < finish; i++ {
				if !planeSeats[row][i] {
					seatsAvailable = false
				}
			}

			if seatsAvailable {
				fmt.Print("Passengers can take seats:")
				for i := start; i < finish; i++ {
					fmt.Printf(" %d%c", row+1, rune('A'+i))
					planeSeats[row][i] = false
				}
				fmt.Println()
				printPlaneSeats(row, start, finish)

				continue outer
			}

			seatsAvailable = true
		}

		fmt.Println("Cannot fulfill passengers requirements.")
	}
}

func startDirection(group passangerGroup) (start, finish int) {
	switch {
	case group.side == "left" && group.position == "aisle":
		start = 3 - group.size
		finish = 3
	case group.side == "left" && group.position == "window":
		start = 0
		finish = group.size
	case group.side == "right" && group.position == "aisle":
		start = 3
		finish = 3 + group.size
	case group.side == "right" && group.position == "window":
		start = 6 - group.size
		finish = 6
	default:
		log.Fatalln("Erroneous input.")
	}

	return start, finish
}

func printPlaneSeats(row, start, finish int) {
	for i, planeRow := range planeSeats {
		for j, seat := range planeRow {
			if i == row && j >= start && j < finish {
				fmt.Print("X")
			} else if seat {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}

			if j == 2 {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}
