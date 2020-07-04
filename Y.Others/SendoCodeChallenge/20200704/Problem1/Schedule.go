package main

import (
	"fmt"
	"os"
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
		var numberOfShop int
		fmt.Fscanf(inputFile, "%d", &numberOfShop)

		if numberOfShop <= 0 {
			fmt.Fprintf(outputFile, "-1 -1 -1\n")
			continue
		}

		maxTime := 23 // 0h..23h

		var hash map[int]*[25]int
		hash = make(map[int]*[25]int)

		resultDate := -1
		resultTimeStart := -1
		resultMax := -1

		for shop := 0; shop < numberOfShop; shop++ {
			var date int
			var timeStart, timeEnd int8
			fmt.Fscanf(inputFile, "%d%d%d", &date, &timeStart, &timeEnd)

			x, ok := hash[date]
			if !ok {
				var xNew [25]int
				for time := 0; time < len(xNew); time++ {
					xNew[time] = 0
				}
				hash[date] = &xNew
				x = &xNew
			}

			x[timeStart]++
			x[timeEnd]--
		}

		for date, x := range hash {
			sum := 0
			for time := 0; time <= maxTime; time++ {
				sum += x[time]
				if (resultMax < sum) ||
					(resultMax == sum && date < resultDate) ||
					(resultMax == sum && date == resultDate && time < resultTimeStart) {
					resultDate = date
					resultMax = sum
					resultTimeStart = time
				}
			}
		}

		fmt.Fprintf(outputFile, "%d %d %d\n", resultMax, resultDate, resultTimeStart)
	}
}

func main() {
	solve("Input.txt", "Output.txt")
	// solve("Input1mil.txt", "Output1mil.txt")
	// solve("Input2mil.txt", "Output2mil.txt")
	// solve("Input3mil.txt", "Output3mil.txt")
	// solve("Input5mil.txt", "Output5mil.txt")
}
