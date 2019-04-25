package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 三数之和

func main() {

	nums := []int{1, 5, 8, 3, 7, 11}
	num := 10
	n := twoSum(nums, num)
	fmt.Println(n)
}


//方法1：暴力破解 时间复杂度O(n^3)
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
func twoSum2(nums []int)  [][]int {
	sort.Ints(nums)
	var n [][]int
	mapk := make(map[string]int)

	for index :=0;index <len(nums)-2 ;index++  {
		if nums[index]==nums[index+2]&&nums[index]!=0 {
			continue
		}
		mapNum := make(map[int]int,len(nums))
		for index1 := index+1; index1 < len(nums); index1++ {
			flag := -(nums[index]+nums[index1])
			if  ins,ok := mapNum[nums[index1]];ok {
				strkey := strconv.Itoa(nums[index])+strconv.Itoa(nums[ins])+strconv.Itoa(nums[index1])
				if _,ok = mapk[strkey]; !ok{
					mapk[strkey] = 1
					n = append(n,[]int{nums[index],nums[ins],nums[index1]})
				}
			}
			mapNum[flag] = index1
		}
	}
	return n
}
