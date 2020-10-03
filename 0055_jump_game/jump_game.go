package jumpGame

//https://leetcode.com/problems/jump-game/
func CanJump(nums []int) bool {
	n := len(nums) - 1
	iMax := nums[0]

	for i := 1; i <= iMax && iMax <= n; i++ {
		if i + nums[i] > iMax {
			iMax = i + nums[i]
		}
	}

	return iMax >= n
}
