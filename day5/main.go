package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type SourceDestinationMap struct {
	Mappings []func(int) (int, error)
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	entireText := ""
	for scanner.Scan() {
		entireText += scanner.Text()
		entireText += "\n"
	}

	Maps := strings.Split(entireText, "\n\n")

	seedsString := strings.Split(Maps[0], ": ")
	eachSeed := strings.Split(seedsString[1], " ")
	seeds := []int{}
	for _, str := range eachSeed {
		seeds = append(seeds, atoi(str))
	}
	allMaps := []SourceDestinationMap{}

	// PART 2 PARSING
	allSeedsPart2 := []Range{}
	allMappingsPart2 := [][]Mapping{}
	for i := 0; i < len(seeds); i += 2 {
		allSeedsPart2 = append(allSeedsPart2, Range{
			Start: seeds[i],
			End:   seeds[i] + seeds[i+1] - 1,
		})
	}

	//

	for i, maps := range Maps {
		if i == 0 { // skip seeds
			continue
		}
		SourceDestinationMap, mappings := parseMap(maps)
		allMaps = append(allMaps, SourceDestinationMap)
		allMappingsPart2 = append(allMappingsPart2, mappings)
	}

	minResult := math.MaxInt

	for _, seed := range seeds { // Part 1
		seed := seed
		for _, sdm := range allMaps {
			for _, f := range sdm.Mappings {
				num, err := f(seed)
				if err == nil {
					seed = num
					break
				}
			}
		}

		if seed < minResult {
			minResult = seed
		}
	}

	fmt.Println(minResult)

	// Part 2
	reuse := allSeedsPart2
	fmt.Println(reuse)
	for _, mappings := range allMappingsPart2 {
		reuse = part2SortArraysIntoArrays(reuse, mappings)
	}

	fmt.Println(reuse)
	minResult = math.MaxInt
	for _, Range := range reuse {
		if Range.Start < minResult {
			minResult = Range.Start
		}
	}
	fmt.Println(minResult)
}

func CreateMapping(Destination, Source, Range int) func(int) (int, error) {
	return func(num int) (int, error) {
		if num < Source || num >= Source+Range {
			return 0, fmt.Errorf("not this one")
		}

		return num - Source + Destination, nil
	}
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func parseMap(input string) (SourceDestinationMap, []Mapping) {
	mappings := []Mapping{}
	output := SourceDestinationMap{
		Mappings: []func(int) (int, error){},
	}
	ranges := strings.Split(input, "\n")
	for i, line := range ranges {
		if i == 0 || line == "" {
			continue
		}
		aRange := strings.Split(line, " ")
		output.Mappings = append(output.Mappings, CreateMapping(atoi(aRange[0]), atoi(aRange[1]), atoi(aRange[2])))
		mappings = append(mappings, Mapping{
			Start:       atoi(aRange[1]),
			End:         atoi(aRange[1]) + atoi(aRange[2]) - 1,
			Destination: atoi(aRange[0]),
		})
	}
	output.Mappings = append(output.Mappings, func(num int) (int, error) {
		return num, nil
	})
	return output, mappings
}

type Range struct {
	Start int
	End   int
}

type Mapping struct {
	Start       int
	End         int
	Destination int
}

func part2SortArraysIntoArrays(inputs []Range, Mappings []Mapping) []Range {
	output := []Range{}
	i := 0
	for {
		if i == len(inputs) {
			break
		}
		input := inputs[i]
		for _, mapping := range Mappings {
			if input.Start >= mapping.Start && input.End <= mapping.End { // Totally inside
				output = append(output, Range{
					Start: mapping.Destination - mapping.Start + input.Start,
					End:   mapping.Destination - mapping.Start + input.End,
				})
				input = Range{0, -1}
				break
			} else if input.Start <= mapping.End && input.Start >= mapping.Start && input.End > mapping.End {
				output = append(output, Range{
					Start: mapping.Destination - mapping.Start + input.Start,
					End:   mapping.Destination - mapping.Start + mapping.End, // Mapping End > x > input end leftover
				})
				input.Start = mapping.End + 1
			} else if input.End >= mapping.Start && input.Start < mapping.Start && input.End <= mapping.End {
				output = append(output, Range{
					Start: mapping.Destination - mapping.Start + mapping.Start,
					End:   mapping.Destination - mapping.Start + input.End, // input start -> mapping start leftover
				})
				input.End = mapping.Start - 1
			} else if input.Start < mapping.Start && input.End > mapping.End { // We encapsulate the mapping
				output = append(output, Range{
					Start: mapping.Destination,
					End:   mapping.Destination + mapping.End - mapping.Start,
				})
				input.End = mapping.Start - 1
				inputs = append(inputs, Range{
					Start: mapping.End + 1,
					End:   input.End,
				})
			}
		}
		if input.End-input.Start >= 0 {
			output = append(output, input) // Leftovers after mappings
		}
		i++
	}
	return output
}
