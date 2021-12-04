package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	inputPath = "./input.txt"
)

//var numIncreased, numWindowIncreased int = 0, 0

func main() {
	numbers := parseInputFile(inputPath)

	numIncreased := calcIncreased(numbers)
	numWindowIncreased := calcSlidingWindowIncreased(numbers, 3)

	fmt.Fprintf(os.Stdout, "Simple increases: %d \nWindow increases: %d", numIncreased, numWindowIncreased)

}

func calcSlidingWindowIncreased(nums []int, windowSize int) int {
	numWindowIncreased := 0
	for i, num := range nums {
		if i+3 >= len(nums) {
			return numWindowIncreased
		}
		currentWindow := num + nums[i+1] + nums[i+2]
		nextWindow := nums[i+1] + nums[i+2] + nums[i+3]

		if nextWindow > currentWindow {
			numWindowIncreased++
		}
	}
	return numWindowIncreased
}

func calcIncreased(nums []int) int {
	numIncreased := 0
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num > nums[i-1] {
			numIncreased++
		}
	}
	return numIncreased
}

func parseInputFile(path string) []int {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error Opening Input file")
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\r\n")

	intArr := make([]int, len(lines))

	for i, strNum := range lines {
		intArr[i], _ = strconv.Atoi(strNum)
	}

	return intArr
}
