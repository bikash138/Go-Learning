package main

import (
	"fmt"
	"slices"
)

func main() {
	//uninitlaised slice
	//By default the unitialised slices are nil
	var nums = make([]int, 3)
	fmt.Println(len(nums))

	//Creating slices with no nil
	var nums2 = make([]int, 2, 5)
	nums2[0] = 1
	nums2[1] = 2
	nums2 = append(nums2, 3)
	fmt.Println(nums2)

	//Copy function -> It does not allocates more space unlike append
	copy(nums, nums2)
	fmt.Println(nums)

	//Slice opearator -> the mentioned index inside the slice opearor dont gets included
	nums3 := []int{1,2,3, 4, 5}
	fmt.Println(nums3[1:2])
	nums4 := []int{4,5,6}
	fmt.Println(slices.Equal(nums3, nums4))
}