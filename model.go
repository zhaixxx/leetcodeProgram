package arith

import "github.com/labstack/gommon/log"

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func quicksort(num []int) int {
	left := 0
	right := len(num) - 1
	key := (num)[left]
	log.Info("left: ", left, ", right: ", right, ", key", key)
	for left < right {
		for num[right] >= key && right > left {
			right--
		}
		if num[left] > num[right] {
			num[left] = num[right]
		}
		for num[left] < key && right > left {
			left++
		}
		if num[left] <= num[right] {
			num[right] = num[left]
		}
	}
	(num)[left] = key
	return left
}
