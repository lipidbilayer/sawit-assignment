package repository

import (
	"testing"
)

func TestCalculateStats(t *testing.T) {
	// Test case 1: Empty input
	objects := []EstateObject{}
	min, max, median := CalculateStats(objects)
	if min != 0 || max != 0 || median != 0 {
		t.Errorf("CalculateStats failed for empty input, expected (0, 0, 0) but got (%d, %d, %d)", min, max, median)
	}

	// Test case 2: Non-empty input
	objects = []EstateObject{{Height: 5}, {Height: 10}, {Height: 15}}
	min, max, median = CalculateStats(objects)
	if min != 5 || max != 15 || median != 10 {
		t.Errorf("CalculateStats failed for non-empty input, expected (5, 15, 10) but got (%d, %d, %d)", min, max, median)
	}
}

func TestCalculateMedian(t *testing.T) {
	// Test case 1: Odd number of elements
	objects := []EstateObject{{Height: 5}, {Height: 10}, {Height: 15}}
	median := CalculateMedian(objects)
	if median != 10 {
		t.Errorf("CalculateMedian failed for odd number of elements, expected 10 but got %d", median)
	}

	// Test case 2: Even number of elements
	objects = []EstateObject{{Height: 5}, {Height: 10}, {Height: 15}, {Height: 20}}
	median = CalculateMedian(objects)
	if median != 12 {
		t.Errorf("CalculateMedian failed for even number of elements, expected 12 but got %d", median)
	}
}

func TestSumDroneTravelDistance(t *testing.T) {
	estate := Estate{Width: 1, Length: 5}
	objects := map[string]EstateObject{
		"2,1": {XLocation: 2, YLocation: 1, Height: 5},
		"3,1": {XLocation: 2, YLocation: 1, Height: 3},
		"4,1": {XLocation: 2, YLocation: 1, Height: 4},
	}

	expectedDistance := 54
	distance := SumDroneTravelDistance(estate, objects)
	if distance != expectedDistance {
		t.Errorf("Expected distance %d, but got %d", expectedDistance, distance)
	}
}
