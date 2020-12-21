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
			val:   "1",
			left:  nil,
			right: nil,
		},
	}
	arith.Cecode103()
}
