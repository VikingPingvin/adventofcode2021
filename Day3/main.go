package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := parseInput("./input.txt")
	if content == nil {
		os.Exit(1)
	}

	gamma := calculateGammaRateBinary(content)
	epsilon := calculateEpsilonRateBinary(gamma)

	fmt.Fprintf(os.Stdout, "Day3 Task1: %d", calcTask1(gamma, epsilon))
}

func calcTask1(gamma, eps string) int {
	return binaryToDecimal(gamma) * binaryToDecimal(eps)
}

func calculateGammaRateBinary(input []string) string {
	retStr := make([]string, len(input[0]))

	for i := 0; i < len(input[0]); i++ {
		ones, zeroes := 0, 0
		for _, binaryLine := range input {
			bin, _ := strconv.Atoi(string(binaryLine[i]))
			if bin == 1 {
				ones++
			} else {
				zeroes++
			}
		}
		if ones > zeroes {
			retStr[i] = "1"
		} else {
			retStr[i] = "0"
		}
	}

	return strings.Join(retStr, "")
}

func calculateEpsilonRateBinary(input string) string {
	var retStr = ""
	for i := 0; i < len(input); i++ {
		current, _ := strconv.Atoi(string(input[i]))

		if current == 1 {
			retStr = retStr + "0"
		} else {
			retStr = retStr + "1"
		}
	}

	return retStr
}

func binaryToDecimal(b string) int {
	num, _ := strconv.ParseInt(b, 2, 32)
	return int(num)
}

func parseInput(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\r\n")
	retArr := make([]string, len(lines))
	for i, line := range lines {
		retArr[i] = line
	}

	return retArr
}
