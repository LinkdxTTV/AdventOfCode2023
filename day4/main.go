package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	CardNumber     int
	WinningNumbers map[int]interface{}
	HaveNumbers    []int
	PointValue     int
	Matches        int
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	AllCards := []Card{}
	NumberOfEachScratchcard := map[int]int{}

	for scanner.Scan() {
		newCard := Card{
			WinningNumbers: map[int]interface{}{},
		}
		text := scanner.Text()
		cardWithNumbers := strings.Split(text, ":")
		cardNumberAsString := strings.Split(cardWithNumbers[0], "Card")
		cardNumberInt, err := strconv.Atoi(strings.TrimSpace(cardNumberAsString[1]))
		if err != nil {
			panic(err)
		}
		newCard.CardNumber = cardNumberInt

		WinningAndHadNumbers := strings.Split(cardWithNumbers[1], "|")
		WinningNumbers := strings.Split(WinningAndHadNumbers[0], " ")
		for _, num := range WinningNumbers {
			if num == " " || num == "" {
				continue
			}
			numInt, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			newCard.WinningNumbers[numInt] = nil
		}

		HadNumbers := strings.Split(WinningAndHadNumbers[1], " ")
		for _, num := range HadNumbers {
			if num == " " || num == "" {
				continue
			}
			numInt, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			newCard.HaveNumbers = append(newCard.HaveNumbers, numInt)
		}

		matches := 0
		for _, num := range newCard.HaveNumbers {
			if _, ok := newCard.WinningNumbers[num]; ok {
				matches++
			}
		}
		newCard.Matches = matches
		newCard.PointValue = int(math.Pow(2, float64(matches-1)))
		sum += newCard.PointValue
		AllCards = append(AllCards, newCard)
		NumberOfEachScratchcard[newCard.CardNumber] = 1

	}
	fmt.Println(sum)

	// Part 2
	for _, card := range AllCards {
		for i := card.CardNumber + 1; i <= card.CardNumber+card.Matches; i++ {
			NumberOfEachScratchcard[i] += NumberOfEachScratchcard[card.CardNumber]
		}
	}

	sum = 0
	for _, cards := range NumberOfEachScratchcard {
		sum += cards
	}
	fmt.Println(sum)
}
