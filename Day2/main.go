package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	x, y := parseInput("./input.txt")
	distanceX, distanceY := calculateDistance2D(x, y)

	aim := calculateAimSeries(y)
	finalX, finalDepth := calculatePositionWithAim(x, aim)

	fmt.Fprintf(os.Stdout, "X * Y : %d\n", distanceX*distanceY)
	fmt.Fprintf(os.Stdout, "Aimed X * Y : %d\n", finalX*finalDepth)

}

func calculateDistance2D(x, y []int) (int, int) {
	totalX := sumSlice(x)
	totalY := sumSlice(y)
	return totalX, totalY
}

func calculatePositionWithAim(x, aim []int) (int, int) {
	cumX, cumDepth := 0, 0
	size := len(x)

	for i := 0; i < size; i++ {
		cumX += x[i]
		//cumAim := aim[i-1] + aim[i]

		cumDepth += x[i] * aim[i]
	}
	return cumX, cumDepth
}

func calculateAimSeries(arr []int) []int {
	aimSeries := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		aimSeries[i] = aimSeries[i-1] + arr[i]
	}
	return aimSeries
}

func sumSlice(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func parseInput(path string) ([]int, []int) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\r\n")
	retX := make([]int, len(lines))
	retY := make([]int, len(lines))

	for i, line := range lines {
		sepLine := strings.Split(line, " ")

		switch sepLine[0] {
		case "forward":
			retX[i], _ = strconv.Atoi(sepLine[1])
		case "up":
			{
				retY[i], _ = strconv.Atoi(sepLine[1])
				retY[i] = retY[i] * -1
			}
		case "down":
			{
				retY[i], _ = strconv.Atoi(sepLine[1])
			}
		}

	}

	return retX, retY
}
