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
	val   string
	left  *TreeNode
	right *TreeNode
}

func (p *Arith_lecode103) Cecode103() [][]int {
	var ans [][]int
	if p.root == nil {
		return ans
	}
	queue := make([]*TreeNode, 1)
	if queue[0] == nil {
		queue[0] = p.root
	} else {
		queue = append(queue, p.root)
	}
	log.Info("Start queue: ", queue)
	for len(queue) > 0 {
		//得到队列长度
		len := len(queue)
		//用于存储同一深度节点的值
		res := make([]int, 1)
		//遍历同一深度全部节点
		for i := 0; i < len; i++ {

		}
		queue = queue[1:]
	}
	log.Info("End queue: ", queue)
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
