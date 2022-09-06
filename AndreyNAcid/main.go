package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	check(err)

	b := bufio.NewScanner(f)
	b.Split(bufio.ScanWords)

	b.Scan()

	m := make(map[int]int)
	for b.Scan() {
		i, err := strconv.Atoi(b.Text())
		check(err)
		m[i]++
	}

	if len(m) == 0 {
		log.Fatalln(-1)
	}

	firstReservoir := 0
	firstBool := true
	prevReservoir := 0
	secondReservoir := 0

	for k := range m {
		if firstBool {
			firstReservoir = k
			firstBool = false
			continue
		}
		secondReservoir = k
		if prevReservoir > secondReservoir {
			log.Fatalln(-1)
		}

	}

	fmt.Println(secondReservoir - firstReservoir)
}

func check(err error) {
	if err != nil {
		log.Fatalln(-1)
	}
}
