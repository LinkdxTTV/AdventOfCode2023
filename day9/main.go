package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Generation struct {
	numbers []int
	subgen  *Generation
	parent  *Generation
}

func (g *Generation) String() string {
	output := fmt.Sprint(g.numbers)
	if g.subgen != nil {
		output += g.subgen.String()
	}
	return output
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	Originals := []*Generation{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		newLineOfInts := []int{}
		lineSplit := strings.Split(scanner.Text(), " ")
		for _, numString := range lineSplit {
			newLineOfInts = append(newLineOfInts, atoi(numString))
		}
		Originals = append(Originals, &Generation{
			numbers: newLineOfInts,
			subgen:  nil,
			parent:  nil,
		})
	}

	// fmt.Println(Originals)

	for _, original := range Originals {
		// Iterate downwards
		workingGeneration := original
		for !IsIntArrayAllZeros(workingGeneration.numbers) {
			workingGeneration.subgen = &Generation{
				numbers: GenerateSubgeneration(workingGeneration.numbers),
			}
			workingGeneration.subgen.parent = workingGeneration
			workingGeneration = workingGeneration.subgen
		}
		// Work our way back up
		// Place the bottom zero
		workingGeneration.numbers = append(workingGeneration.numbers, 0)

		// Part 2 place the left zero
		newLeft := []int{0}
		newLeft = append(newLeft, workingGeneration.numbers...)
		workingGeneration.numbers = newLeft

		workingGeneration = workingGeneration.parent
		for {
			// part 1
			workingGeneration.numbers = append(workingGeneration.numbers, workingGeneration.numbers[len(workingGeneration.numbers)-1]+workingGeneration.subgen.numbers[len(workingGeneration.subgen.numbers)-1])

			// part 2
			newLeft := []int{workingGeneration.numbers[0] - workingGeneration.subgen.numbers[0]}
			newLeft = append(newLeft, workingGeneration.numbers...)
			workingGeneration.numbers = newLeft

			if workingGeneration.parent == nil {
				break
			}
			workingGeneration = workingGeneration.parent
		}

	}
	// fmt.Println(Originals)
	// Part 1
	sum := 0
	for _, Original := range Originals {
		sum += Original.numbers[len(Original.numbers)-1]
	}
	fmt.Println(sum)
	// Part 2
	sum = 0
	for _, Original := range Originals {
		sum += Original.numbers[0]
	}
	fmt.Println(sum)
}

func GenerateSubgeneration(input []int) []int {
	output := []int{}
	for i := 1; i < len(input); i++ {
		output = append(output, input[i]-input[i-1])
	}
	return output
}

func IsIntArrayAllZeros(input []int) bool {
	for _, num := range input {
		if num != 0 {
			return false
		}
	}
	return true
}

func atoi(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return num
}
