package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func partOne(ch chan<- string, wg *sync.WaitGroup) {
	// Clean up
	defer wg.Done()

	// Open the input file for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// We'll use a scanner so we can read line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {

	}

	// Solution
	ch <- "Part 1: TODO"
}

func partTwo(ch chan<- string, wg *sync.WaitGroup) {
	// Clean up
	defer wg.Done()

	// Open the input file for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// We'll use a scanner so we can read line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {

	}

	// Solution
	ch <- "Part 2: TODO"
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
