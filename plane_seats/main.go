package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	findSeats()

}

var planeSeats [][6]bool // T - free, F - taken
var groups []passangerGroup

type passangerGroup struct {
	size     int
	side     string
	position string
}

func init() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	// Getting initial info aboout plane seats
	scanner.Scan()
	rows, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln()
	}

	for i := 0; i < rows; i++ {
		var planeRow [6]bool
		scanner.Scan()

		j := 0
		for _, r := range scanner.Text() {
			if r == '.' {
				planeRow[j] = true
				j++
			} else if r == '#' {
				planeRow[j] = false
				j++
			}
		}

		planeSeats = append(planeSeats, planeRow)
	}

	// Getting info about passanger groups
	scanner.Scan()
	groupsNum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln()
	}

	for i := 0; i < groupsNum; i++ {
		var group passangerGroup

		scanner.Scan()
		group.size, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}

		scanner.Scan()
		group.side = scanner.Text()

		scanner.Scan()
		group.position = scanner.Text()

		groups = append(groups, group)
	}
}

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
