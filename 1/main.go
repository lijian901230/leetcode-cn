package main

import (
	"fmt"
)

// 两数之和

func main() {

	nums := []int{1, 5, 8, 3, 7, 11}
	num := 10
	n := twoSum(nums, num)
	fmt.Println(n)
}


//方法1：暴力破解 时间复杂度O(n^2)
func twoSum(nums []int, target int) []int {
	var n []int
	for index := 0; index < len(nums); index++ {
		for index1 := index + 1; index1 < len(nums); index1++ {
			if target == (nums[index] + nums[index1]) {
				n = append(n, index, index1)
				return n
			}
		}
	}
	return n
}

//方法2：
func twoSum2(nums []int, target int)  []int {
	var n []int
	mapNum := make(map[int]int,len(nums))
	for index := 0; index < len(nums); index++ {
		flag := target-nums[index]
		if  ins,ok := mapNum[flag];ok {
			return []int{ins,index}
		}
		mapNum[nums[index]] = index
	}
	return n
}
