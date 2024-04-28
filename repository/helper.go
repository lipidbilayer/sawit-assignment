package repository

import (
	"math"
	"sort"
	"strconv"
)

func CalculateStats(objects []EstateObject) (int, int, int) {
	if len(objects) == 0 {
		return 0, 0, 0
	}

	min := objects[0].Height
	max := objects[0].Height
	for _, tree := range objects {
		if tree.Height < min {
			min = tree.Height
		}
		if tree.Height > max {
			max = tree.Height
		}
	}

	// Calculate median height
	median := CalculateMedian(objects)

	return min, max, median
}

func CalculateMedian(objects []EstateObject) int {
	heights := make([]int, len(objects))
	for i, tree := range objects {
		heights[i] = tree.Height
	}

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})

	n := len(heights)
	if n%2 == 0 {
		return (heights[n/2-1] + heights[n/2]) / 2
	}
	return heights[n/2]
}

func SumDroneTravelDistance(estate Estate, objectMap map[string]EstateObject) int {
	currentHeight := 1
	distance := 1

	start := 1
	increment := 1
	for y := 1; y <= estate.Width; y++ {
		for x := start; x > 0 && x <= estate.Length; {
			coordinate := strconv.Itoa(x) + "," + strconv.Itoa(y)
			if val, ok := objectMap[coordinate]; ok {
				distance += int(math.Abs(float64(currentHeight - (val.Height + 1))))
				currentHeight = val.Height + 1
			} else {
				distance += int(math.Abs(float64(1 - currentHeight)))
				currentHeight = 1
			}
			x += increment
			if x > 0 && x < estate.Length {
				distance += 10
			}
		}
		start = estate.Length
		if y%2 == 0 {
			start = 1
		}
		if y <= estate.Width {
			distance += 10
		}
		increment = increment * -1
	}

	distance += 1
	return distance
}
