package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxRed   int = 12
	maxBlue  int = 14
	maxGreen int = 13
)

type Game struct {
	red      int
	blue     int
	green    int
	maxRed   int
	maxBlue  int
	maxGreen int
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	gamesParsed := map[int]Game{}

	for scanner.Scan() {
		txt := scanner.Text()
		gameArray := strings.Split(txt, ":")
		if len(gameArray) == 0 {
			continue
		}
		gameString := gameArray[0]
		gameNumberArray := strings.Split(gameString, " ")
		gameNumber, err := strconv.Atoi(gameNumberArray[1])
		if err != nil {
			panic(err)
		}
		var newGame Game
		isValidGame := true
		turns := strings.Split(gameArray[1], ";")
		for _, turn := range turns {
			newGame.red, newGame.blue, newGame.green = 0, 0, 0
			draws := strings.Split(turn, ",")
			for _, draw := range draws {
				split := strings.Split(draw, " ")
				number, err := strconv.Atoi(split[1])
				if err != nil {
					panic(err)
				}
				switch split[2] {
				case "red":
					newGame.red += number
				case "blue":
					newGame.blue += number
				case "green":
					newGame.green += number
				default:
					panic(split[2])
				}
				if newGame.red > newGame.maxRed {
					newGame.maxRed = newGame.red
				}
				if newGame.blue > newGame.maxBlue {
					newGame.maxBlue = newGame.blue
				}
				if newGame.green > newGame.maxGreen {
					newGame.maxGreen = newGame.green
				}
			}
			// if newGame.red > maxRed || newGame.blue > maxBlue || newGame.green > maxGreen { // All games are valid for part 2
			// 	isValidGame = false
			// 	break
			// }

		}
		if isValidGame {
			gamesParsed[gameNumber] = newGame
		}
	}
	// sum := 0 // Part 1 Sum
	// for gameNumbers, _ := range gamesParsed {
	// 	sum += gameNumbers
	// }

	// Part 2 sum
	sum := 0
	for _, games := range gamesParsed {
		power := games.maxBlue * games.maxGreen * games.maxRed
		sum += power
	}

	fmt.Println(sum)
}
