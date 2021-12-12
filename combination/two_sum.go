package combination

// 求两数之和
// 时间复杂度O(n2)
// 空间复杂度O(2n)
func TwoSumFF(nums []int, target int) []int {
	for _, x := range nums {
		for _, y := range nums {
			if x+y == target {
				return []int{x, y}
			}
		}
	}
	return nil
}

// 求两数之和
// 时间复杂度O(2n)
// 空间复杂度O(n)
func TwoSumF(nums []int, target int) []int {
	var temps = make(map[int]int)
	for index, x := range nums {
		temps[x] = index
	}

	for _, y := range nums {
		z := target - y
		if _, ok := temps[z]; ok {
			return []int{y, z}
		}
	}
	return nil
}

// 求两数之和
// 时间复杂度O(n)
// 空间复杂度O(n)
func TwoSumf(nums []int, target int) []int {
	var temps = make(map[int]int)
	for index, y := range nums {
		z := target - y
		if _, ok := temps[z]; ok {
			return []int{y, z}
		}
		temps[y] = index
	}
	return nil
}
