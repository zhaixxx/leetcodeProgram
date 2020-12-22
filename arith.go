package arith

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"math"
	"strings"
)

type Arith_stringToInt struct {
	symbol int
	ansStr string
	num    int
	Ans    int
}

type Arith_palindrom struct {
	symbol int
	ansStr string
	num    int
	Ans    int
}
type Arith_conver struct {
	numrows    int
	str        map[int]interface{}
	goingwhere bool
	s          string
}
type Arith_lecode103 struct {
	root *TreeNode
}
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}
type Arith_Trap struct {
	height []int
	ans    int
}

/*
	接雨水，暴力解法，根据当前柱子的高度h，横向两边查找比h高的柱子，
	取两边最高柱子的最小值，时间复杂度O(n^2)。
	该题可以优化，通过一次for循环遍历整个数组，通过单向流水机制，
	从左到右取最大值，从右到左取最小值，最后取相同索引下的最小值进行相加
	时间复杂度O(n)
*/
func (p *Arith_Trap) Leetcode42() {
	for i := 0; i < len(p.height); i++ {
		max_lefy := p.height[i]
		max_right := p.height[i]
		for j := i; j >= 0; j-- {
			if max_lefy < p.height[j] {
				max_lefy = p.height[j]
			}
		}
		for j := i; j < len(p.height); j++ {
			if max_right < p.height[j] {
				max_right = p.height[j]
			}
		}
		if max_right >= max_lefy {
			p.ans = p.ans + max_lefy - p.height[i]
		} else {
			p.ans = p.ans + max_right - p.height[i]
		}
		log.Info("max_lefy: ", max_lefy, "___max_right: ", max_right, "___i: ", i)
		log.Info("ans: ", p.ans)
	}

}

func (p *Arith_lecode103) Cecode103() [][]int {
	var ans [][]int
	//边界处理
	if p.root == nil {
		return ans
	}
	queue := make([]*TreeNode, 1)
	//防止node与nil组合在一起
	if queue[0] == nil {
		queue[0] = p.root
	} else {
		queue = append(queue, p.root)
	}
	//判断是否从左向右遍历
	//若为false则从右向左遍历
	leftToRight := true
	for len(queue) > 0 {
		//得到队列长度
		lenn := len(queue)
		//用于存储同一深度节点的值
		res := make([]int, lenn)
		//遍历同一深度全部节点
		for i := 0; i < lenn; i++ {
			root := queue[i]
			if leftToRight {
				res[i] = root.val
			} else {
				res[lenn-i-1] = root.val
			}
			if root.left != nil {
				queue = append(queue, root.left)
			}
			if root.right != nil {
				queue = append(queue, root.right)
			}
		}
		log.Info("The res is : ", res)
		ans = append(ans, res)
		queue = queue[lenn:]
		leftToRight = !leftToRight
	}
	log.Info("End ans: ", ans)
	return ans
}

func (p *Arith_conver) Coir(str string, n int) string {
	if p.str == nil {
		p.str = make(map[int]interface{}, n+1)
	}
	p.goingwhere = false
	p.numrows = 0
	for _, i := range str {
		if p.str[p.numrows] == nil {
			p.str[p.numrows] = string(i)
		} else {
			p.str[p.numrows] = fmt.Sprint(p.str[p.numrows], string(i))
		}

		switch p.numrows {
		case 0:
			p.goingwhere = false
		case n - 1:
			p.goingwhere = true
		}
		//log.Info("num: ", num, "__________char: ", string(i))
		if p.goingwhere {
			p.numrows--
			continue
		}
		p.numrows++
	}
	for i := range p.str {
		p.s = fmt.Sprint(p.s, p.str[i])
	}
	log.Info(p.s)
	return p.s
}

func (p *Arith_stringToInt) VerifyString(str string) {
	//去除空格
	p.ansStr = strings.TrimSpace(str)
	//判断字符串正负，无法判断则返回空
	if p.ansStr == "" {
		return
	}

	switch p.ansStr[0] {
	case '-':
		p.symbol, p.ansStr = -1, p.ansStr[1:]
	case '+':
		p.symbol, p.ansStr = 1, p.ansStr[1:]
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		p.symbol, p.ansStr = 1, p.ansStr
	default:
		p.ansStr = ""
		return
	}
	//提取有效字符串
	for i, strs := range p.ansStr {
		if strs >= '0' && strs <= '9' {
			continue
		} else {
			p.ansStr = p.ansStr[:i]
			break
		}

	}
	return
}
func (p *Arith_stringToInt) StringToInt() {
	p.num = 0
	for _, ans := range p.ansStr {
		p.num = p.num*10 + int(ans-'0')
		switch {
		case p.num > math.MaxInt32 && p.symbol == 1:
			p.Ans = math.MaxInt32
			return
		case p.num*p.symbol < math.MinInt32 && p.symbol == -1:
			p.Ans = math.MinInt32
			return
		}
	}
	p.Ans = p.symbol * p.num
}

func (p *Arith_palindrom) LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	} else {
		end, start := 0, 0
		for i := 0; i < len(s); i++ {
			left, right := Expend(s, i, i)
			left2, right2 := Expend(s, i, i+1)
			if right-left > end-start {
				end = right
				start = left
			}
			if right2-left2 > end-start {
				end = right2
				start = left2
			}

		}
		return s[start : end+1]
	}

}

func Expend(s string, left int, right int) (int, int) {
	log.Info("-------------------------------------------")
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		log.Info("left:", left, "right:", right)
	}
	return left + 1, right - 1
}
