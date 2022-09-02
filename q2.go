package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type SortInterface []int

func (a SortInterface) Len() int {
	return len(a)
}

func (a SortInterface) Less(x int, y int) bool {
	return a[x] < a[y]
}

func (a SortInterface) Swap(x int, y int) {
	tmp := a[x]
	a[x] = a[y]
	a[y] = tmp
}

func main() {
	var x int
	fmt.Print("Enter an int: ")
	fmt.Scanln(&x)
	rand.Seed(time.Now().UnixMilli())
	var mySlice SortInterface = make([]int, x+1, x+1)
	for i := 0; i < x; i++ {
		mySlice[i] = rand.Int()
	}

	startTime := time.Now().UnixMilli()
	sort.Sort(mySlice)
	endTime := time.Now().UnixMilli()
	fmt.Println("Took", (endTime - startTime), "milliseconds to Sort the slice")
	startTime = time.Now().UnixMilli()
	sort.Stable(mySlice)
	endTime = time.Now().UnixMilli()
	fmt.Println("Took", (endTime - startTime), "milliseconds to Stable the slice")
}
