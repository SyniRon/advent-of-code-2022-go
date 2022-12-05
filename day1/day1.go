package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// pull file of calorie loads
	input, _ := os.ReadFile("./day1/input.txt")
	// convert file from byte to string type
	inputString := string(input)
	// split string by double newline into slices
	splitString := strings.Split(inputString, "\n\n")
	loadList := []int{}
	// iterate over every slice
	for _, carriedCalories := range splitString {
		// split each slice into slices for each calories
		elf := strings.Split(carriedCalories, "\n")
		elfLoad := 0
		// for each smaller slice convert to int and add to total
		for _, calories := range elf {
			caloriesStr := calories
			caloriesInt, _ := strconv.Atoi(caloriesStr)
			elfLoad += caloriesInt
		}
		// add these totals to our list of loads
		loadList = append(loadList, elfLoad)
	}
	// sort the list from largets to smallest
	sort.Slice(loadList, func(i, j int) bool { return loadList[i] > loadList[j] })
	// print top 3 loads
	fmt.Printf("The largest load is %v\n", loadList[0])
	fmt.Printf("The second largest load is %v\n", loadList[1])
	fmt.Printf("The third largest load is %v\n", loadList[2])
	// add top 3 loads
	totalLargeLoad := loadList[0] + loadList[1] + loadList[2]
	// print total of top 3 loads
	fmt.Printf("The total of the 3 largest loads is %v\n", totalLargeLoad)
}
