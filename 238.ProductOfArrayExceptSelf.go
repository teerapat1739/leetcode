package leetcode

/*
Runtime: 24 ms, faster than 96.60% of Go online submissions for Product of Array Except Self.
Memory Usage: 7.8 MB, less than 81.47% of Go online submissions for Product of Array Except Self.
*/

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))

	for i, _ := range products {
		products[i] = 1
	}

	leftRunning := 1
	for i := 0; i < len(nums); i++ {
		products[i] *= leftRunning
		leftRunning *= nums[i]
	}

	rightRunning := 1
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] *= rightRunning
		rightRunning *= nums[i]
	}
	return products
}
