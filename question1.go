package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	fmt.Print("Enter an int:")
	fmt.Scanln(&x)
	startTime := time.Now().UnixMilli()
	mySlice := make([]int, x+1, x+1)
	for i := 0; i <= x; i++ {
		mySlice[i] = i
	}
	endTime := time.Now().UnixMilli()
	fmt.Println("Took", (endTime - startTime), "milliseconds to initialize slice")
	//maps
	startTime = time.Now().UnixMilli()
	myMap := make(map[int]int)
	for i := 0; i <= x; i++ {
		myMap[i] = i
	}
	endTime = time.Now().UnixMilli()
	fmt.Println("Took", (endTime - startTime), "milliseconds to initialize map")

	//incrementing things
	//slice
	startTime = time.Now().UnixMilli()
	for i := 0; i <= x; i++ {
		mySlice[i] = mySlice[i] + 1
	}
	endTime = time.Now().UnixMilli()
	fmt.Println("Took", (endTime - startTime), "milliseconds to increment slice")

	//map
	startTime = time.Now().UnixMilli()
	for i := 0; i <= x; i++ {
		myMap[i] = myMap[i] + 1
	}
	endTime = time.Now().UnixMilli()
	fmt.Println("Took", (endTime - startTime), "milliseconds to increment map")
}
