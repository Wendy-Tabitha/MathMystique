package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"math-skills/mathskills"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	data := string(args)
	value := []float64{}
	g := strings.Split(data, "\n")
	for _, str := range g {

		str = strings.TrimSpace(str)

		if str == "" {
			continue
		}

		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error converting the string to float")
			os.Exit(1)
		}

		value = append(value, num)
	}
	if len(value) == 0 {
		return
	}
	//sorts the data in the data.txt file
	sort.Float64s(value)

	average := mathskills.Average(value)
	median := mathskills.Median(value)
	variance := mathskills.Variance(value)
	deviation := mathskills.Deviation(value)

	fmt.Printf("Average: %.0f\n", math.Round(average))
	fmt.Printf("Median: %.0f\n", math.Round(median))
	fmt.Printf("Variance: %.0f\n", math.Round(variance))
	fmt.Printf("Deviation: %.0f\n", math.Round(deviation))
}
