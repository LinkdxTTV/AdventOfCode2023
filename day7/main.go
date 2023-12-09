package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Score Score
	Bid   int
}

type Score struct {
	Type       int
	Tiebreaker []int
}

var stringToIntValuePart1 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var stringToIntValuePart2 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allHands := []Hand{}
	i := 0
	for scanner.Scan() {
		i++
		if scanner.Text() == "" {
			continue
		}
		split := strings.Split(scanner.Text(), " ")
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		allHands = append(allHands, Hand{
			Cards: split[0],
			Score: part2Shim(split[0]), // Replace with GetScoreFromCards for part1
			Bid:   bid,
		})
	}

	sort.Slice(allHands, func(i, j int) bool {
		if allHands[i].Score.Type == allHands[j].Score.Type {
			for k := 0; k < 5; k++ {
				if allHands[i].Score.Tiebreaker[k] < allHands[j].Score.Tiebreaker[k] {
					return true
				} else if allHands[i].Score.Tiebreaker[k] > allHands[j].Score.Tiebreaker[k] {
					return false
				}
			}
		}
		return allHands[i].Score.Type < allHands[j].Score.Type
	})

	// fmt.Println(allHands)

	totalWinnings := 0
	for index, hand := range allHands {
		totalWinnings += (index + 1) * hand.Bid
	}
	fmt.Println(totalWinnings)
}

// Five of a kind -> 7
// ...
// ...
// High card -> 1

func GetScoreFromCards(cards string) Score {
	score := Score{}
	cardMap := map[string]int{}
	for _, cardRune := range cards {
		cardMap[string(cardRune)]++
		score.Tiebreaker = append(score.Tiebreaker, stringToIntValuePart2[string(cardRune)]) // Change part for each part here
	}

	// Analyze card map
	sortable := []int{}
	for _, num := range cardMap {
		sortable = append(sortable, num)
	}

	sort.Slice(sortable, func(i, j int) bool {
		return sortable[i] < sortable[j]
	})

	if len(sortable) == 1 {
		score.Type = 7
		return score
	}
	if sortable[len(sortable)-1] == 4 {
		score.Type = 6
		return score
	}
	if sortable[len(sortable)-1] == 3 && sortable[len(sortable)-2] == 2 {
		score.Type = 5
		return score
	}
	if sortable[len(sortable)-1] == 3 {
		score.Type = 4
		return score
	}
	if sortable[len(sortable)-1] == 2 && sortable[len(sortable)-2] == 2 {
		score.Type = 3
		return score
	}
	if sortable[len(sortable)-1] == 2 {
		score.Type = 2
		return score
	}
	score.Type = 1
	return score
}

type CardWithJ struct {
	cards      string
	jLocations []int
}

func part2Shim(cards string) Score {
	// Shortcut
	if cards == "JJJJJ" {
		return Score{
			Type:       7,
			Tiebreaker: []int{1, 1, 1, 1, 1},
		}
	}
	possibleCards := []CardWithJ{{cards, []int{}}}
	i := 0
	for {
		if i == len(possibleCards) {
			break
		}
		for location, cardRune := range possibleCards[i].cards {
			if string(cardRune) == "J" {
				for k, _ := range stringToIntValuePart2 {
					if k == "J" {
						continue
					}
					newCards := possibleCards[i].cards[:location] + k + possibleCards[i].cards[location+1:]
					newPossible := CardWithJ{newCards, possibleCards[i].jLocations}
					newPossible.jLocations = append(newPossible.jLocations, location)
					possibleCards = append(possibleCards, newPossible)
				}
			}
		}
		i++
	}

	allScores := []Score{}
	for _, cards := range possibleCards {
		if strings.Contains(cards.cards, "J") {
			continue
		}
		score := GetScoreFromCards(cards.cards)
		for _, location := range cards.jLocations {
			score.Tiebreaker[location] = 1
		}
		allScores = append(allScores, score)
	}

	sort.Slice(allScores, func(i, j int) bool {
		if allScores[i].Type == allScores[j].Type {
			for k := 0; k < 5; k++ {
				if allScores[i].Tiebreaker[k] < allScores[j].Tiebreaker[k] {
					return true
				} else if allScores[i].Tiebreaker[k] > allScores[j].Tiebreaker[k] {
					return false
				}
			}
		}
		return allScores[i].Type < allScores[j].Type
	})

	return allScores[len(allScores)-1]
}
