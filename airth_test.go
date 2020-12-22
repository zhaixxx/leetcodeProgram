package arith

import (
	"fmt"
	"testing"
)

func TestArith_palindrom_LongestPalindrome(t *testing.T) {
	arith := Arith_palindrom{}
	str := arith.LongestPalindrome("ac")
	fmt.Println(str)
}
func TestArith_conver_Conver(t *testing.T) {
	str := "LEETCODEISHIRING"
	n := 3
	arith := Arith_conver{}
	arith.Coir(str, n)
}
func TestArith_lecode103_Cecode103(t *testing.T) {
	arith := Arith_lecode103{
		root: &TreeNode{
			val: 1,
			left: &TreeNode{
				val: 2,
				left: &TreeNode{
					val:   4,
					left:  nil,
					right: nil,
				},
				right: &TreeNode{
					val:   5,
					left:  nil,
					right: nil,
				},
			},
			right: &TreeNode{
				val:   3,
				left:  nil,
				right: nil,
			},
		},
	}
	arith.Cecode103()
}
func TestTrap_Leetcode42(t *testing.T) {
	num := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	arith := Arith_Trap{
		height: num,
		ans:    0,
	}
	arith.Leetcode42()

}
