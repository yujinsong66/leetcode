package main

import "fmt"

func twoSum(sums []int, target int) []int {
	m := make(map[int]int)

	for index, value := range sums {
		diff := target - value
		v, ok := m[diff]
		if ok {
			return []int{v, index}
		} else {
			m[value] = index
		}
	}

	return nil
}

func main() {
	nums := []int{2, 5, 3}
	target := 5
	fmt.Println(twoSum(nums, target))
}
