package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var digitsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func partOne(ch chan<- string, wg *sync.WaitGroup) {
	// Clean up
	defer wg.Done()

	// Regex to find digits in the input strings
	numRegex := regexp.MustCompile(`[0-9]`)

	// Open the input file for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Open a second file for output
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Sum the total
	var sum int

	// We'll use a scanner so we can read line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// Get digits
		digitArray := numRegex.FindAllString(scanner.Text(), -1)

		// Store first and last in a string
		result := string(digitArray[0]) + string(digitArray[len(digitArray)-1])

		// Debug
		// fmt.Println(scanner.Text(), ":", digitArray, ":", result)

		// Write to output
		outputFile.WriteString(result)

		// Add to sum
		number, err := strconv.Atoi(result)
		if err != nil {
			// Failed to convert string to int
			log.Fatal(err)
			continue
		}

		// Conversion was successful, add to sum
		sum += number
	}

	// Close the files
	inputFile.Close()
	outputFile.Close()

	// Solution
	ch <- fmt.Sprintf("Part 1 solution is: %v", sum)
}

func partTwo(ch chan<- string, wg *sync.WaitGroup) {
	// Clean up
	defer wg.Done()

	// Regex to find digits in the input strings
	numRegex := regexp.MustCompile(`[0-9]`)

	// Open the input file for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Open a second file for output
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Sum the total
	var sum int

	// We'll use a scanner so we can read line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// Replace string representations of digits
		line := replaceStringsWithNumbers(scanner.Text())

		// Get digits
		digitArray := numRegex.FindAllString(line, -1)

		// Store first and last in a string
		result := string(digitArray[0]) + string(digitArray[len(digitArray)-1])

		// Debug
		// fmt.Println(scanner.Text(), ":", line, ":", digitArray, ":", result)

		// Write to output
		outputFile.WriteString(result)

		// Add to sum
		number, err := strconv.Atoi(result)
		if err != nil {
			// Failed to convert string to int
			log.Fatal(err)
			continue
		}

		// Conversion was successful, add to sum
		sum += number
	}

	// Close the files
	inputFile.Close()
	outputFile.Close()

	// Solution
	ch <- fmt.Sprintf("Part 2 solution is: %v", sum)
}

func replaceStringsWithNumbers(input string) string {
	// Iterate over the digit <> string map
	for k, v := range digitsMap {
		// Find the location of all matching strings
		// indexes is an [][]int, where:
		//  - The outer array represents the set of matches
		//  - The inner array is two ints representing the start and end index of a match
		// https://pkg.go.dev/regexp#Regexp.FindAllStringIndex
		indexes := regexp.MustCompile(k).FindAllStringIndex(input, -1)

		// For each match ...
		for _, r := range indexes {
			// ... replace the second character of the match with the corresponding digit
			// e.g., eight becomes e8ght
			// This allows all cases of multiple overlapping matches to work
			// e.g., oneight > o1e8ght > 18
			input = input[:r[0]+1] + fmt.Sprint(v) + input[r[0]+2:]
		}

	}
	return input
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
