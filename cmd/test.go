package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// )

// type Estate struct {
// 	Width  int
// 	Length int
// }

// type EstateObject struct {
// 	XLocation int
// 	YLocation int
// 	Height    int
// }

// func main() {
// 	estate := Estate{Width: 1, Length: 5}
// 	objects := []EstateObject{
// 		{XLocation: 2, YLocation: 1, Height: 10},
// 		{XLocation: 3, YLocation: 1, Height: 20},
// 		{XLocation: 4, YLocation: 1, Height: 10},
// 		// {XLocation: 2, YLocation: 1, Height: 5},
// 		// {XLocation: 3, YLocation: 1, Height: 3},
// 		// {XLocation: 4, YLocation: 1, Height: 4},
// 	}

// 	objectsAsMap := ConvertObjectsToMap(objects)
// 	distance := TraverseThroughField(estate, objectsAsMap)
// 	fmt.Println(distance)
// }

// func TraverseThroughField(estate Estate, objectMap map[string]EstateObject) int {
// 	currentHeight := 1
// 	distance := IncrementDistance(0, 1, "start fly")

// 	start := 1
// 	increment := 1
// 	for y := 1; y <= estate.Width; y++ {
// 		for x := start; x > 0 && x <= estate.Length; {
// 			coordinate := strconv.Itoa(x) + "," + strconv.Itoa(y)
// 			// fmt.Print(coordinate + " ")
// 			if val, ok := objectMap[coordinate]; ok {
// 				distance = IncrementDistance(distance, int(math.Abs(float64(currentHeight-(val.Height+1)))), "found tree")
// 				currentHeight = val.Height + 1
// 			} else {
// 				distance = IncrementDistance(distance, int(math.Abs(float64(1-currentHeight))), "ground level")
// 				currentHeight = 1
// 			}
// 			x += increment
// 			if x > 0 && x < estate.Length {
// 				distance = IncrementDistance(distance, 10, "move x")
// 			}
// 		}
// 		start = estate.Length
// 		if y%2 == 0 {
// 			start = 1
// 		}
// 		if y <= estate.Width {
// 			distance = IncrementDistance(distance, 10, "move y axis")
// 		}
// 		increment = increment * -1
// 	}

// 	distance = IncrementDistance(distance, 1, "landing")
// 	return distance
// }

// func IncrementDistance(currentDistance int, addDistance int, log string) int {
// 	fmt.Printf("%d+%d log:%s\n", currentDistance, addDistance, log)
// 	currentDistance += addDistance
// 	return currentDistance
// }

// func ConvertObjectsToMap(objects []EstateObject) (objectMap map[string]EstateObject) {
// 	if objectMap == nil {
// 		objectMap = map[string]EstateObject{}
// 	}
// 	for _, value := range objects {
// 		coordinate := strconv.Itoa(value.XLocation) + "," + strconv.Itoa(value.YLocation)
// 		objectMap[coordinate] = value
// 	}
// 	return objectMap
// }
