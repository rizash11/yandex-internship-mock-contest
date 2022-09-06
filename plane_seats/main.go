package main

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
