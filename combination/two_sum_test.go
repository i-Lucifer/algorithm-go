package combination

import (
	"fmt"
	"testing"
)

// 求数之和
func TestTwoSum(t *testing.T) {
	var nums = []int{2, 7, 11, 15}
	// res := TwoSumFF(nums, 9)
	// res := TwoSumF(nums, 9)
	res := TwoSumf(nums, 9)
	fmt.Println(res)
}
