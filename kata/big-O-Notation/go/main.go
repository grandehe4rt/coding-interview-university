package main

import "fmt"

// O(N)
// O(N^2)
// O(log N)
// O(n log n)

var theArray []int
var arraySize int
var itemsInArray int = 0
var startTime float32
var endTime float32

func main() {
	fmt.Println(addItemToArray(5))
}

// O(1)
func addItemToArray(newItem int) []int {
	return append(theArray, newItem)
}
