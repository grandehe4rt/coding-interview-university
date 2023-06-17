package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O(N^2)
// O(log N)
// O(n log n)

// var theArray []int
// var arraySize int
// var itemsInArray int
var startTime int64
var endTime int64

func main() {
	b := &BigONotation{}
	result := b.addItemToArray(5)
	fmt.Println(result)
}

type BigONotation struct {
	theArray     []int
	arraySize    int
	itemsInArray int
}

func NewBigONotation(size int) *BigONotation {
	return &BigONotation{
		arraySize:    size,
		theArray:     make([]int, size),
		itemsInArray: 0,
	}
}

func (b *BigONotation) generateRandomArray() {
	for i := 0; i < b.arraySize; i++ {
		b.theArray[i] = rand.Intn(1000) + 10
	}
}

// O(1)
func (b *BigONotation) addItemToArray(newItem int) []int {
	// Literalmente adiciona um elemento ao Array (Pika)
	return append(b.theArray, newItem)
}

// O(N)
func (b *BigONotation) linearSearchForValue(value int) {
	var valueInArray bool = false
	var indexsWithValues string = ""

	startTime = time.Now().UnixMilli()

	for i := 0; i < b.arraySize; i++ {
		if b.theArray[i] == value {
			valueInArray = true
			indexsWithValues += fmt.Sprint(i) + " "
		}
	}

	fmt.Println("Value found: ", valueInArray)
	endTime = time.Now().UnixMilli()
	fmt.Println("Linear Search Took ", fmt.Sprint(endTime-startTime))

}
