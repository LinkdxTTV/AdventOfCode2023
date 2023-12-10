package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

var nodeMap = map[string]*Node{}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	directions := scanner.Text()

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}

		split := strings.Split(txt, " = ")
		paths := strings.Split(split[1], ", ")
		left := strings.TrimPrefix(paths[0], "(")
		right := strings.TrimSuffix(paths[1], ")")
		node := &Node{
			left:  left,
			right: right,
		}
		nodeMap[split[0]] = node
	}

	// fmt.Println(directions)
	// fmt.Println(nodeMap)

	currentNode := "AAA"
	steps := 0
	i := 0
	var nextNode string
	for {
		if currentNode == "ZZZ" {
			break
		}
		if i >= len(directions) {
			i = 0
		}
		path := string(directions[i])
		switch path {
		case "L":
			nextNode = nodeMap[currentNode].left
		case "R":
			nextNode = nodeMap[currentNode].right
		default:
			panic("uh")
		}
		currentNode = nextNode
		steps++
		i++
	}
	fmt.Println(steps)

	// Part 2
	starters := []string{}
	for k := range nodeMap {
		if string(k[2]) == "A" {
			starters = append(starters, k)
		}
	}

	fmt.Println(starters)
	currentNodes := starters
	i = 0
	steps = 0
	endInZSteps := [][]int{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{}} // 6 runners
	for {
		if steps%10000 == 0 {
			for i, steps := range endInZSteps {
				if len(steps) < 1 {
					continue
				}
				fmt.Println(i, steps[0])
			}
		}
		if DoAllNodesEndInZ(currentNodes) {
			break
		}
		nextNodes := []string{}
		if i >= len(directions) {
			i = 0
		}
		path := string(directions[i])

		for i, node := range currentNodes {
			var nextNode string
			switch path {
			case "L":
				nextNode = nodeMap[node].left
				nextNodes = append(nextNodes, nextNode)
			case "R":
				nextNode = nodeMap[node].right
				nextNodes = append(nextNodes, nextNode)
			default:
				panic("lol")
			}
			if string(nextNode[2]) == "Z" {
				endInZSteps[i] = append(endInZSteps[i], steps)
			}
		}
		currentNodes = nextNodes
		// Do cycle detection
		for i, starter := range endInZSteps {
			if len(starter) < 3 {
				continue
			}
			len := len(starter)
			if starter[len-1]-starter[len-2] == starter[len-2]-starter[len-3] {
				fmt.Println("cycle detected for ", i, starter[len-1]-starter[len-2])
			}
		}
		i++
		steps++

		// Find the lowest common multiple of all the cycle numbers
		/*
			cycle detected for  0 15871
			cycle detected for  1 19099
			cycle detected for  2 18023
			cycle detected for  3 21251
			cycle detected for  4 12643
			cycle detected for  5 16409
		*/
		// 17099847107071
	}

	fmt.Println(steps)
}

func DoAllNodesEndInZ(nodes []string) bool {
	for _, node := range nodes {
		if string(node[2]) != "Z" {
			return false
		}
	}
	return true
}
