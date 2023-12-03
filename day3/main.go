package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x       int
	y       int
	char    string
	checked bool
}

type numberWithSymbol struct {
	number int
	point  *point
}

var NumberSet = map[string]interface{}{
	"0": nil,
	"1": nil,
	"2": nil,
	"3": nil,
	"4": nil,
	"5": nil,
	"6": nil,
	"7": nil,
	"8": nil,
	"9": nil,
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]*point{}
	rowCounter := 0
	for scanner.Scan() {
		txt := scanner.Text()
		gridRow := []*point{}
		for column, rune := range txt {
			gridRow = append(gridRow, &point{
				x:       column,
				y:       rowCounter,
				char:    string(rune),
				checked: false,
			})
		}
		rowCounter++
		grid = append(grid, gridRow)
	}

	xMax := len(grid[0]) - 1
	// yMax := len(grid) -1
	validPartNums := []numberWithSymbol{}

	for _, row := range grid {
		for _, gridPoint := range row {
			if gridPoint.checked {
				continue
			}
			numberAsString := ""
			pointsToConsider := []*point{}
			for gridPoint.isNumber() {
				numberAsString += gridPoint.char
				pointsToConsider = append(pointsToConsider, gridPoint)
				gridPoint.checked = true
				if gridPoint.x == xMax {
					break
				}
				gridPoint = grid[gridPoint.y][gridPoint.x+1]
			}
			if numberAsString != "" {
				if symbolPoint, hasNearbySymbol := pointsHaveNearbySymbol(grid, pointsToConsider); hasNearbySymbol {
					validPartNumber, err := strconv.Atoi(numberAsString)
					if err != nil {
						panic(err)
					}
					validPartNums = append(validPartNums, numberWithSymbol{
						number: validPartNumber,
						point:  symbolPoint,
					})
				}
			}
		}
	}

	fmt.Println(validPartNums)
	sum := 0
	for _, validNum := range validPartNums {
		sum += validNum.number
	}

	fmt.Println(sum)

	// Part 2
	gearsMap := map[*point][]int{}
	for _, validNum := range validPartNums {
		if validNum.point.char == "*" {
			gearsMap[validNum.point] = append(gearsMap[validNum.point], validNum.number)
		}
	}

	part2Sum := 0
	for _, numbersAtPoint := range gearsMap {
		if len(numbersAtPoint) == 2 {
			product := numbersAtPoint[0] * numbersAtPoint[1]
			part2Sum += product
		}
	}
	fmt.Println(part2Sum)
}

type gridDelta struct {
	xDel int
	yDel int
}

var checkSurroundingList = []gridDelta{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}}

func pointsHaveNearbySymbol(grid [][]*point, points []*point) (*point, bool) {
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1

	for _, point := range points {

		for _, delta := range checkSurroundingList {
			newX := point.x + delta.xDel
			newY := point.y + delta.yDel
			if newX < 0 || newX > maxX || newY < 0 || newY > maxY {
				continue
			}
			if grid[newY][newX].char == "." || grid[newY][newX].isNumber() {
				continue
			} else {
				return grid[newY][newX], true
			}
		}
	}
	return &point{}, false
}

func (p *point) isNumber() bool {
	_, isNumber := NumberSet[p.char]
	return isNumber
}
