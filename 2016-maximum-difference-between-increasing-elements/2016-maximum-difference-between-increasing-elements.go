func maximumDifference(nums []int) int {
    maxDiffVal := -1
	lowestIVal := nums[0]

	for _, num := range nums {
		diffVal := num - lowestIVal
		if diffVal > maxDiffVal && diffVal > 0 {
			maxDiffVal = diffVal
		}
		if num < lowestIVal {
			lowestIVal = num
		}
	}
	return maxDiffVal
}