package problems

// Problem 1. Two Sum
// URL: https://leetcode.com/problems/two-sum/tabs/description

func twoSum(nums []int, target int) []int {
	n := len(nums)

	if n < 2 {
		return nil
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum1(nums []int, target int) []int {
	tbl := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := tbl[target-v]; ok {
			return []int{j, i}
		}

		tbl[v] = i
	}

	return nil
}
