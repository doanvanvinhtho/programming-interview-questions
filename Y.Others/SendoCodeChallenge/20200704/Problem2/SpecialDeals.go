package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(inputFilename, outputFilename string) {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("Open input file error: %v", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Printf("Create output file error: %v", err)
		return
	}
	defer outputFile.Close()

	var numberOfTestcase int
	fmt.Fscanf(inputFile, "%d", &numberOfTestcase)

	for t := 0; t < numberOfTestcase; t++ {
		var n, k int
		fmt.Fscanf(inputFile, "%d%d", &n, &k)

		if n <= 0 {
			fmt.Fprintf(outputFile, "-1\n")
			continue
		}

		var a []int
		a = make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Fscanf(inputFile, "%d", &a[i])
		}

		sort.Ints(a)

		sum := 0
		j := 1
		for i := n - 1; i >= 0; i-- {
			if j == k {
				j = 1
			} else {
				sum += a[i]
				j++
			}
		}

		fmt.Fprintf(outputFile, "%d\n", sum)
	}
}

func main() {
	solve("Input.txt", "Output.txt")
	// solve("Input100.txt", "Output100.txt")
	// solve("Input500.txt", "Output500.txt")
	// solve("Input999.txt", "Output999.txt")
	// solve("Input1000.txt", "Output1000.txt")
}
