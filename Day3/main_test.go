package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputPath = "./input_test.txt"

func TestGammaRateBinary(t *testing.T) {
	testInput := parseInput(testInputPath)

	expectedGamma := "10110"

	assert.Equal(t, expectedGamma, calculateGammaRateBinary(testInput))
}

func TestGammaToEpsilon(t *testing.T) {

	input := "10110"
	expected := "01001"

	assert.Equal(t, expected, calculateEpsilonRateBinary(input))

}

func TestParseInput(t *testing.T) {
	expectedValue := []string{"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"00111",
		"01111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	content := parseInput(testInputPath)

	if !assert.ElementsMatch(t, expectedValue, content) {
		t.Errorf("got %v, want %v", expectedValue, content)
	}
}

func TestBinToDecimal(t *testing.T) {
	input := "10110"
	expected := 22

	assert.Equal(t, expected, binaryToDecimal(input))
}
