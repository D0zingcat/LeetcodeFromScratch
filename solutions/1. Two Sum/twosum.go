package ___Two_Sum

func twoSumS1(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i; j < len(nums); j++ {
			if (nums[i] + nums[j] == target) {
				return []int{i, j}
			}
		}
	}
	return nil
}

