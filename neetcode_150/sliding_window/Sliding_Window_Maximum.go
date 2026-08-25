func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	currMax := -10001

	start := 0
	end := -1
	for end-start+1 != k {
		end++
		if nums[end] > currMax {
			currMax = nums[end]
		}
	}
	result = append(result, currMax)

	for end < len(nums)-1 {
		start++
		end++
		if nums[start-1] == currMax {
			currMax = -10001
			for i := start; i < end+1; i++ {
				if nums[i] > currMax {
					currMax = nums[i]
				}
			}
		} else if nums[end] > currMax {
			currMax = nums[end]
		}
		result = append(result, currMax)
	}
	return result
}

