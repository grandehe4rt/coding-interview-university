package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O(N^2)
// O(log N)
// O(n log n)

var theArray []int = []int{1, 2, 3, 4, 5}
var arraySize int
var itemsInArray int = 0
var startTime int64
var endTime int64

func main() {
	// fmt.Println(linearSearchForValue(5))
	linearSearchForValue(100)
}

func generateRandomArray() {
	for i := 0; i < arraySize; i++ {
		arraySlice[i] = rand.Intn(1000) + 10
	}
}

// O(1)
func addItemToArray(newItem int) []int {
	// Literalmente adiciona um elemento ao Array (Pika)
	return append(theArray, newItem)
}

// O(N)
func linearSearchForValue(value int) {
	var valueInArray bool = false
	var indexsWithValues string = ""

	startTime = time.Now().UnixMilli()

	for i := 0; i < arraySize; i++ {
		if theArray[i] == value {
			valueInArray = true
			indexsWithValues += fmt.Sprint(i) + " "
		}
	}

	endTime = time.Now().UnixMilli()
	fmt.Println("Value found: ", valueInArray)
	fmt.Println("Linear Search Took ", fmt.Sprint(endTime-startTime))

}
