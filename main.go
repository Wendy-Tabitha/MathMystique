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
	sort.Float64s(value)

	average := mathskills.Average(value)

	fmt.Printf("Average: %g\n", math.Round(average))
}
