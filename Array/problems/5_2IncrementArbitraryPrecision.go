package problems

// 66. Plus One
// https://leetcode.com/problems/plus-one/description/

func PlusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {

		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
