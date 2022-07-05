package array

func getConcatenation(nums []int) []int {
	ans := make([]int, 2*len(nums))
	for i := 0; i < len(nums); i++ {
		j := i + len(nums)
		ans[i], ans[j] = nums[i], nums[i]
	}
	return ans
}
