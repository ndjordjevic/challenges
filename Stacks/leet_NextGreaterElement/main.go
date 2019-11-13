package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement2([]int{2, 4}, []int{1, 2, 3, 4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	length := len(nums2)

	for _, num1 := range nums1 {
		found := false
		for i, num2 := range nums2 {
			if num2 == num1 {
				found = true
			}
			if num2 > num1 && found == true {
				stack = append(stack, num2)
				break
			}
			if i == length-1 {
				stack = append(stack, -1)
			}
		}
	}

	return stack
}

func nextGreaterElement2(findNums []int, nums []int) []int {
	var mp = make(map[int]int)
	var ans []int
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}

	for i := 0; i < len(findNums); i++ {
		key := mp[findNums[i]]
		for j := key; j < len(nums); j++ {
			if nums[j] > findNums[i] {
				ans = append(ans, nums[j])
				break
			}
			if j == len(nums)-1 {
				ans = append(ans, -1)
			}
		}
	}

	return ans
}
