package main

import (
	"fmt"
	"labs/pkg/display"
	"labs/pkg/mysort"
	"labs/pkg/tools"
	"math/rand"
	"os"
	"time"
)

const maxRandomValue = 100000

var arr []int

func main() {
	fmt.Println("Welcome.")
	menu()
}

func NewRandomArray(cnt int) []int {
	arr := make([]int, cnt)

	for i := range arr {
		arr[i] = rand.Intn(maxRandomValue)
	}

	return arr
}

func menu() {
	fmt.Println("Menu:")
	fmt.Println("1. Create a new randomized array.")
	fmt.Println("2. Sort an array.")
	fmt.Println("3. Quit")
	fmt.Print("Your choice: ")

	choice := tools.ReadInt()

	switch choice {
	case 1:
		fmt.Print("Enter array length: ")
		n := tools.ReadInt()

		fmt.Println("Randomizing a new array.")
		now := time.Now()
		arr = NewRandomArray(n)
		tools.MeasureTime(now)
		sortMenu()
	case 2:
		sortMenu()
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Incorrect input.")
		menu()
	}
}

func sortMenu() {
	if arr == nil {
		now := time.Now()
		arr = NewRandomArray(10000)
		tools.MeasureTime(now)
		fmt.Println("Created new randomized array with length 1000.")
	}

	fmt.Println("Choose the sort algorithm:")
	fmt.Println("1. Insertion sort")
	fmt.Println("2. Selection sort")
	fmt.Println("3. Merge sort")
	fmt.Println("4. Quick sort")
	fmt.Println("5. Quit")
	fmt.Print("Your choice: ")

	sortChoice := tools.ReadInt()

	switch sortChoice {
	case 1:
		fmt.Println("---Insertion sort---")
		temp := copyArr()
		now := time.Now()
		temp = mysort.InsertionSort(temp)
		tools.MeasureTime(now)
		display.PrintArray(temp)
	case 2:
		fmt.Println("---Selection sort---")
		temp := copyArr()
		now := time.Now()
		temp = mysort.SelectionSort(temp)
		tools.MeasureTime(now)
		display.PrintArray(temp)
	case 3:
		fmt.Println("---Merge sort---")
		temp := copyArr()
		now := time.Now()
		temp = mysort.MergeSort(temp)
		tools.MeasureTime(now)
		display.PrintArray(temp)
	case 4:
		fmt.Println("---Quick sort---")
		temp := copyArr()
		now := time.Now()
		temp = mysort.QuickSort(temp)
		tools.MeasureTime(now)
		display.PrintArray(temp)
	case 5:
		menu()
	default:
		fmt.Println("Incorrect input.")
	}
}

func copyArr() []int {
	temp := make([]int, len(arr))

	copy(temp, arr)
	return temp
}
