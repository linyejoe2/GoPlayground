package main

import (
	"fmt"
	"math/rand"
)

func initArray() {
	var arr [5]int

	arr[0] = 1
	
	// println(&arr[0])
	// println(&arr[1])
	// println(0xf0 - 0xe8)

	nums := []int{1, 3, 2, 5, 4}

	println(nums)
	fmt.Printf("%v\n", nums)

	randomN := randomAccess(nums)
	println(randomN)

	insert(&nums, 88, 0)
	fmt.Printf("%v\n", nums)
}

func randomAccess(nums []int) (randomNum int) {
	randomIndex := rand.Intn(len(nums))

	randomNum = nums[randomIndex]
	return
}

func insert(nums *[]int, num int, index int) {
	arr := *nums
	for i := index; i <= len(arr); i++ {
		if i == len(arr) { 
			*nums = append(arr, num)
			return
		}
		temp := arr[i]
		arr[i] = num
		num = temp
	}
}

