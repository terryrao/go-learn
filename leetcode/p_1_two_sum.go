package leetcode

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return nums
		}
		return nil
	}

	var result []int
	for idx1, first := range nums {
		for idx2, sec := range nums {
			if idx1 == idx2 {
				continue
			}
			if first+sec == target {
				return []int{first, sec}
			}
		}
	}
	return result

}

func twoSum2(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return nums
		}
		return nil
	}

	intMap := make(map[int]int)
	var result []int
	for idx, v := range nums {
		d := target - v
		if sec, ok := intMap[d]; ok {
			return []int{idx, sec}
		} else {
			intMap[v] = idx
		}
	}
	return result

}
