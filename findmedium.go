package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1 == 0 && len2 == 0 {
		return 0.0
	}

	result := make([]int, 0)

	i := 0
	j := 0

	for i < len1 && j < len2 {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	for i < len1 {
		result = append(result, nums1[i])
		i++
	}
	for j < len2 {
		result = append(result, nums2[j])
		j++
	}

	lens := len1 + len2

	if lens%2 == 0 {
		v1 := result[lens/2-1]
		v2 := result[lens/2]
		return float64((float64(v1) + float64(v2)) / 2)
	} else {
		return float64(result[lens/2])
	}

}

func main() {
	nums1 := []int{2}
	nums2 := []int{}

	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println("result == ")
	fmt.Println(result)
}
