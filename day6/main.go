package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	TimesString := strings.Split(scanner.Text(), "Time:        ")

	Times := []int{}
	Distances := []int{}
	TimesSplit := strings.Split(TimesString[1], "     ")
	for _, time := range TimesSplit {
		num, err := strconv.Atoi(time)
		if err != nil {
			panic(err)
		}
		Times = append(Times, num)
	}

	scanner.Scan()
	DistancesString := strings.Split(scanner.Text(), "Distance:   ")
	DistancesSplit := strings.Split(DistancesString[1], "   ")
	for _, dist := range DistancesSplit {
		num, err := strconv.Atoi(dist)
		if err != nil {
			panic(err)
		}
		Distances = append(Distances, num)
	}

	fmt.Println(Times)
	fmt.Println(Distances)

	records := []int{}

	for i := 0; i < len(Times); i++ {
		recordBeaters := 0
		for j := 0; j <= Times[i]; j++ {
			dist := j * (Times[i] - j)
			if dist > Distances[i] {
				recordBeaters++
			}
		}
		records = append(records, recordBeaters)
	}

	product := 1
	for _, record := range records {
		product = product * record
	}

	fmt.Println(product)

	// Part 2
	realTime := ""
	for _, time := range TimesSplit {
		realTime += time
	}
	realDistance := ""
	for _, dist := range DistancesSplit {
		realDistance += dist
	}
	realTimeInt, err := strconv.Atoi(realTime)
	if err != nil {
		panic(err)
	}
	realDistInt, err := strconv.Atoi(realDistance)
	if err != nil {
		panic(err)
	}

	wins := 0
	for i := 0; i <= realTimeInt; i++ {
		dist := i * (realTimeInt - i)
		if dist > realDistInt {
			wins++
		}
	}

	fmt.Println(wins)
}
