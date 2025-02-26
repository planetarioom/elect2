package main

import (
	"fmt"
	"strconv"
)

func sortAscending(orig []int) []int {
	copy := append([]int{}, orig...)

	n := len(copy)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if copy[j] > copy[j+1] {
				copy[j], copy[j+1] = copy[j+1], copy[j]
			}
		}
	}

	return copy
}

func sortDescending(orig []int) []int {
	copy := append([]int{}, orig...)

	n := len(copy)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if copy[j] < copy[j+1] {
				copy[j], copy[j+1] = copy[j+1], copy[j]
			}
		}
	}
	
	return copy
}

func sortEvenOdd(orig []int) ([]int, []int) {
	var evenSlice, oddSlice []int

	for _, num := range orig {
		if num % 2 == 0{
			evenSlice = append(evenSlice, num)
		} else {
			oddSlice = append(oddSlice, num)
		}
	}

	return evenSlice, oddSlice
}

func mean(orig []int) float64 {
    if len(orig) == 0 {
        fmt.Println("Empty slice.")
        return 0
    }

    sum := 0

    for _, num := range orig {
        sum += num
    }
    return float64(sum) / float64(len(orig))
}

func median(orig []int) float64 {
    n := len(orig)
    if n == 0 {
        fmt.Println("Empty string.")
        return 0
    }

    if n % 2 == 1 {
        return float64(orig[n/2])
    } 
    return float64(orig[(n/2)-1]+orig[n/2]) / 2.0
}

func mode(orig []int) []int {
    frequency := make(map[int]int)
    maxFrequency := 0

    for _, num := range orig {
        frequency[num]++
        if frequency[num] > maxFrequency {
            maxFrequency = frequency[num]
        }
    }

    var modes []int 
    for num, count := range frequency {
        if count == maxFrequency {
            modes = append(modes, num)
        }
    }

    return modes
}

func main() {
	var size int

	for {
		var input string
		fmt.Print("Enter an integer for the array size: ")
		fmt.Scanln(&input)

		if num, err := strconv.Atoi(input); err == nil {
			size = num
			break
		}
		fmt.Println("Please enter an integer.")
	}

	arr := make([]int, size)

	for i := range arr {
		var num int

		fmt.Printf("Enter number for index %v: ", i)
		fmt.Scanln(&num)

		arr[i] = num
	}

    arrAsc := sortAscending(arr)
    fmt.Println("")
	fmt.Printf("Ascending order: %v\n", arrAsc)
	fmt.Printf("Descending order: %v\n\n", sortDescending(arr))

	even, odd := sortEvenOdd(arr)
	fmt.Printf("Even numbers sorted: %v\n", even)
	fmt.Printf("Odd numbers sorted: %v\n\n", odd)

    fmt.Printf("Mean: %v\n", mean(arr))
    fmt.Printf("Median: %v\n", median(arrAsc))
    fmt.Printf("Mode: %v\n", mode(arr))
}