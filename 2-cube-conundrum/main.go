package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
)

// Configure supported cubes
const redCubeName = "red"
const blueCubeName = "blue"
const greenCubeName = "green"

// Configure max cubes per game
const maxRedCubes = 12
const maxGreenCubes = 13
const maxBlueCubes = 14

// map cube names to max
var cubesTypes = map[string]int{
	redCubeName:   maxRedCubes,
	blueCubeName:  maxBlueCubes,
	greenCubeName: maxGreenCubes,
}

var gameNumRegex = regexp.MustCompile(`[0-9]*:`)

func partOne(ch chan<- string, wg *sync.WaitGroup) {
	// Clean up
	defer wg.Done()

	// Open the input file for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Sum the game numbers
	var sum int

	// We'll use a scanner so we can read line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// Get the game record
		gameLine := scanner.Text()

		// Get the game number (including trailing semicolon)
		gameNum := gameNumRegex.FindString(gameLine)

		// Strip the semicolon
		gameNum = strings.Trim(gameNum, ":")

		// Trim everything before the colon (':')
		gameData := strings.Split(gameLine, ":")[1]

		// Split into rounds
		rounds := strings.Split(gameData, ";")

		// Process rounds
		isValid := true

		for _, round := range rounds {
			// Remove commas
			round = strings.ReplaceAll(round, ",", "")

			// Split on spaces to give us tokens that represent the game data
			tokens := strings.Split(round, " ")

		CUBES:
			for cubeColour, cubeMax := range cubesTypes {
				// Get the index of the colour name
				index := slices.Index(tokens, cubeColour)

				// Handle cube not present in round
				if index < 0 {
					continue CUBES
				}

				// Get the number of cubes of that colour (the token before the matching name colour)
				number, err := strconv.Atoi(tokens[index-1])
				if err != nil {
					// Failed to convert string to int
					log.Fatal(err)
					continue CUBES
				}

				// Debug
				// fmt.Printf("Game %v Round %v CubeColour %v CubeMax %v Index %v Value %v\n", gameNum, i, cubeColour, cubeMax, index, number)

				// Set the flag if the game is not valid
				if number > cubeMax {
					isValid = false
				}
			}
		}

		// If the game is valid, add the game number to the sum
		if isValid {
			number, err := strconv.Atoi(gameNum)
			if err != nil {
				// Failed to convert string to int
				log.Fatal(err)
				continue
			}

			sum += number
		}
	}

	// Solution
	ch <- fmt.Sprintf("Part 1: %v", sum)
}

func partTwo(ch chan<- string, wg *sync.WaitGroup) {
	// Clean up
	defer wg.Done()

	// Open the input file for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Sum the game numbers
	var sum int

	// We'll use a scanner so we can read line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// Get the game record
		gameLine := scanner.Text()

		// Get the game number (including trailing semicolon)
		gameNum := gameNumRegex.FindString(gameLine)

		// Strip the semicolon
		gameNum = strings.Trim(gameNum, ":")

		// Trim everything before the colon (':')
		gameData := strings.Split(gameLine, ":")[1]

		// Split into rounds
		rounds := strings.Split(gameData, ";")

		// Store the minimum possible cubes for the game to be valid
		minCubeCount := map[string]int{
			redCubeName:   0,
			blueCubeName:  0,
			greenCubeName: 0,
		}

		for _, round := range rounds {
			// Remove commas
			round = strings.ReplaceAll(round, ",", "")

			// Split on spaces to give us tokens that represent the game data
			tokens := strings.Split(round, " ")

		CUBES:
			for cubeColour := range cubesTypes {
				// Get the index of the colour name
				index := slices.Index(tokens, cubeColour)

				// Handle cube not present in round
				if index < 0 {
					continue CUBES
				}

				// Get the number of cubes of that colour (the token before the matching name colour)
				number, err := strconv.Atoi(tokens[index-1])
				if err != nil {
					// Failed to convert string to int
					log.Fatal(err)
					continue CUBES
				}

				// Update the max cube count
				if minCubeCount[cubeColour] < number {
					minCubeCount[cubeColour] = number
				}
			}
		}

		// Calculate powers and add it to the sum

		power := 1
		for _, v := range minCubeCount {
			power *= v
		}

		// Debug
		// fmt.Printf("Game %v MinRedCubes %v MinBlueCubes %v MinGreenCubes %v Power %v\n", gameNum, minCubeCount[redCubeName], minCubeCount[blueCubeName], minCubeCount[greenCubeName], power)

		sum += power

	}

	// Solution
	ch <- fmt.Sprintf("Part 2: %v", sum)
}

func main() {
	// Using a WaitGroup to ensure all subroutines complete
	var wg sync.WaitGroup

	// Using channels to get results from goroutines
	ch := make(chan string, 2)

	// Run goroutines
	wg.Add(2)
	go partOne(ch, &wg)
	go partTwo(ch, &wg)

	// Wait for completion
	wg.Wait()
	close(ch)

	for result := range ch {
		// Print results
		fmt.Println(result)
	}
}
