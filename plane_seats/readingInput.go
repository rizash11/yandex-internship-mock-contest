package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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
