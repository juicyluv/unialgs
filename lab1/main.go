package main

import (
	"fmt"
	"labs/pkg/tools"
	"math/rand"
	"time"
)

const (
	maxRandomValue = 100
)

type Array struct {
	arr []int
}

func (a *Array) Push(value int) {
	t := time.Now()

	a.arr = append(a.arr, value)

	tools.MeasureTime(t)
}

func (a *Array) Delete(value int) error {
	t := time.Now()

	index := a.Search(value)
	if index == -1 {
		return fmt.Errorf("there is no such element: %d", value)
	}

	a.arr = append(a.arr[:index], a.arr[index+1:]...)

	tools.MeasureTime(t)

	return nil
}

func (a *Array) Search(value int) int {
	t := time.Now()

	for i, v := range a.arr {
		if v == value {
			return i
		}
	}

	tools.MeasureTime(t)

	return -1
}

func (a *Array) BinarySearch(value int) int {
	size := len(a.arr)
	if size == 0 {
		return -1
	}

	return a.binarySearch(0, size, value)
}

func (a *Array) binarySearch(left, right, value int) int {
	if right < left {
		return -1
	}

	mid := left + (right-left)/2
	if a.arr[mid] == value {
		return mid
	}

	if a.arr[mid] > value {
		return a.binarySearch(left, mid-1, value)
	}

	return a.binarySearch(mid+1, right, value)
}

func (a *Array) Sort() {
	t := time.Now()

	for i := range a.arr {
		for j := range a.arr {
			if a.arr[i] < a.arr[j] {
				a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
			}
		}
	}

	tools.MeasureTime(t)
}

func (a *Array) String() string {
	return fmt.Sprintf("%v", a.arr)
}

func NewRandomArray(n int) *Array {
	t := time.Now()
	temp := make([]int, n)

	for i := range temp {
		temp[i] = rand.Intn(maxRandomValue)
	}

	tools.MeasureTime(t)

	return &Array{
		arr: temp,
	}
}

func main() {
	fmt.Print("Enter array length: ")
	n := tools.ReadInt()

	arr := NewRandomArray(n)
	fmt.Printf("Created array: %v\n", arr)

	fmt.Println("Inserting values 20, 30...")
	arr.Push(20)
	arr.Push(30)
	fmt.Printf("Array after inserting: %v\n", arr)

	fmt.Printf("Deleting value 20...\n")
	err := arr.Delete(20)

	if err != nil {
		panic("delete failed")
	}

	fmt.Printf("Array after deleting: %v\n", arr)

	fmt.Println("Sorting...")
	arr.Sort()
	fmt.Printf("Array after sort: %v\n", arr)

	fmt.Println("Searching index of 30...")
	linearIndex := arr.Search(30)
	fmt.Printf("Found 30. Index: %d\n", linearIndex)

	fmt.Println("Binary search of 30...")
	binaryIndex := arr.BinarySearch(30)
	fmt.Printf("Found 30. Index: %d\n", binaryIndex)

	fmt.Println("Finding negative number(should return -1)...")
	notFoundIndex := arr.BinarySearch(-123)

	if notFoundIndex == -1 {
		fmt.Println("Element is not found")
	} else {
		panic("failed to find negative number")
	}
}
