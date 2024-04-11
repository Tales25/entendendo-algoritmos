package main

import "fmt"

func main() {
	list := createList(10000)

	pos, err := binarySearch(list, 10000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found in position: ", pos)
	}
}

func createList(size int) []int {
	var list []int
	for i := 0; i < size; i++ {
		list = append(list, i)
	}
	return list
}

func binarySearch(list []int, target int) (int, error) {
	if len(list) == 0 {
		return 0, fmt.Errorf("Empty list")
	}

	first := 0
	last := len(list) - 1

	for first <= last {
		mid := (last - first) / 2
		num := list[mid]

		if target > num {
			first = mid
		} else if target < num {
			last = mid
		} else {
			return mid, nil
		}
	}
	return 0, fmt.Errorf("Number not found")
}
