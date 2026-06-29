package problems

// 75. Sort Colors
// https://leetcode.com/problems/sort-colors/description/

func DutchFlagPartition(pivotIndex int, nums []int) []int {
	pivot := nums[pivotIndex]
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch {
		case nums[mid] < pivot:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case nums[mid] == pivot:
			mid++
		default:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}

	return nums
}
