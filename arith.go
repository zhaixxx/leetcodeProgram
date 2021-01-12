package arith

import (
	"container/list"
	"fmt"
	"github.com/labstack/gommon/log"
	"math"
	"path"
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

type Arith_daily struct {
	temp []int
	ans  []int
}
type Arith_profit struct {
	prices []int
	ans    int
}
type Arith_substring struct {
	str    string
	ans    string
	length int
}
type Arith_nnum struct {
	num  []int
	num2 []int
	ans  []int
}
type Arith_clean struct {
	path    string
	relPath string
}
type Arith_VaildStack struct {
	pushed []int //12345
	poped  []int //32154
}
type ListNode struct {
	Val  int
	Next *ListNode
}
type Lru struct {
	LruCache map[int]*list.Element
	LruList  *list.List
	Len      int
}
type Point struct {
	key   int
	value string
}
type Arith_Num struct {
	num [3]int
}

func (p *Arith_Num) ThreeSum() {

}
func (p *Arith_Num) QuickSort(low, high int) {
	if low < high {
		mid := 0
		if len(p.num) >= high {
			mid = quicksort(p.num[low:high]) + low
		} else {
			mid = quicksort(p.num[low:high+1]) + low
		}

		p.QuickSort(low, mid-1)
		p.QuickSort(mid+1, high)
	}

}
func (l *Lru) Get(n int) string {
	element, ok := l.LruCache[n]
	if ok {
		l.LruList.MoveToFront(element)
		return element.Value.(Point).value
	}
	return "nil"
}
func (l Lru) Put(n int, str string) {
	ele, ok := l.LruCache[n]
	if ok {
		l.LruList.MoveToFront(ele)
		ele.Value = Point{
			key:   n,
			value: str,
		}
	} else {
		if l.LruList.Len() == l.Len {
			delete(l.LruCache, l.LruList.Back().Value.(Point).key)
			l.LruList.Remove(l.LruList.Back())
		}
		l.LruList.PushFront(Point{
			key:   n,
			value: str,
		})
		l.LruCache[n] = l.LruList.Front()
	}
}

func (p *ListNode) ReverseList(head *ListNode) *ListNode {
	var foot *ListNode = nil
	node := head
	next := node.Next
	for node != nil {
		log.Info("Start node: ", node.Val)
		node.Next = foot
		foot = node
		node = next
		if next != nil {
			next = next.Next
		}
	}
	return foot
}
func (p *ListNode) getKthFromEnd(head *ListNode, k int) *ListNode {
	first := head
	second := head
	for i := 0; i < k; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
}
func (p *ListNode) swapTwoNode(root *ListNode) *ListNode {
	if root.Next == nil {
		return root
	}
	hair := ListNode{}
	hair.Next = root
	firstNode := root
	secondNood := root.Next
	node := root.Next
	head := &hair
	for node != nil {
		node = secondNood.Next
		log.Info("第", +1, "次交换：first: ", firstNode.Val, "，second: ", secondNood.Val)
		head.Next = secondNood
		secondNood.Next = firstNode
		firstNode.Next = node
		if node != nil {
			head = firstNode
			secondNood = node.Next
			firstNode = node
			node = secondNood
		}
		log.Info("第", +1, "次交换后：first: ", firstNode.Val, "，second: ", secondNood.Val)
	}

	return hair.Next

}
func (p *Arith_VaildStack) validateStackSequences() bool {
	stack := []int{}
	for i := 0; i < len(p.pushed); i++ {
		stack = append(stack, p.pushed[i])
		for p.poped[0] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			//copy(stack, stack[:len(stack)-1])
			//stack = append()
			p.poped = p.poped[1:]
			//copy(p.poped, p.poped[1:])
			if len(stack) == 0 {
				break
			}
		}
	}
	if len(stack) == 0 && len(p.poped) == 0 {
		return true
	}
	return false
}
func (p *Arith_clean) LeetcodeClean() {
	p.relPath = path.Clean(p.path)
}

func (p *Arith_substring) SubstringFromLongTew() int {
	sigle := make(map[string]int)
	for i, j := 0, 0; j < len(p.str); j++ {
		if num, ok := sigle[string(p.str[j])]; ok {
			i = max(num+1, i)
		}
		p.length = max(j-i+1, p.length)
		sigle[string(p.str[j])] = j
	}
	log.Info("length: ", p.length)
	return p.length
}
func (p *Arith_substring) SubstringFromLong() string {
	sigle := make(map[string]int)
	str := ""
	count := -1
	left := 0
	for i := 0; i < len(p.str); i++ {
		if c, ok := sigle[string(p.str[i])]; !ok || c < count {
			str = str + string(p.str[i])
			sigle[string(p.str[i])] = i
		} else {
			sigle[string(p.str[i])] = i
			if len(str)-left > len(p.ans) {
				p.ans = str[left:]
			}
			str = str + string(p.str[i])
			//str = str[c+1:]
			left = c + 1
			count = c + 1
		}
	}
	if len(p.ans) < len(str)-left {
		p.ans = str[left:]
	}
	log.Info("string is :", p.ans)
	return p.ans
}

func (p *Arith_nnum) MergeTwoNum() []int {
	lena := len(p.num)
	lenb := len(p.num2)
	p.ans = make([]int, 0)
	for i, j := 0, 0; i+j < lena+lenb; {
		if i >= lena {
			p.ans = append(p.ans, p.num2[j:]...)
			break
		}
		if j >= lenb {
			p.ans = append(p.ans, p.num[i:]...)
			break
		}
		if p.num[i] >= p.num2[j] {
			p.ans = append(p.ans, p.num2[j])
			j++
		} else if p.num[i] < p.num2[j] {
			p.ans = append(p.ans, p.num[i])
			i++
		}
	}
	log.Info(p.ans)
	return p.ans

}

/*
	变种，可以购买多个股票，但是只能单线操作，不能手持多股
	把每一次上坡都计算出并相加即可 需要求边界情况
*/
func (p *Arith_profit) Leetcode122() {
	if len(p.prices) == 0 {
		return
	}
	minprice := p.prices[0]
	sumprice := 0
	for i := 1; i < len(p.prices); i++ {
		if p.prices[i-1] > p.prices[i] {
			p.ans += sumprice
			minprice = p.prices[i]
			sumprice = 0
		} else {
			sumprice = p.prices[i] - minprice
		}
		log.Info("minprice: ", minprice, ", p.prices: ", p.prices[i], ", p.ans: ", p.ans, ", sumprice: ", sumprice)
	}
	p.ans += sumprice
	log.Info("ans: ", p.ans)
}

/*
	股票取最大利润，可以暴力求解，时间复杂度应该是O(n^2)
	但是还有一种解法，假设当前是第i天，我们用一个变量存储[0-i）天的最小价格，
	那么 price[i]-minprice就是我们需要的当天的最大利润，以此类推遍历整个数组
	时间复杂度O(n)
*/
func (p *Arith_profit) Leetcode121() {
	if len(p.prices) == 0 {
		return
	}
	lenn := len(p.prices)
	minprice := p.prices[0]

	for i := 0; i < lenn; i++ {
		if minprice > p.prices[i] {
			minprice = p.prices[i]
		} else if p.ans > p.prices[i]-minprice {
			p.ans = p.prices[i] - minprice
		}
	}
	log.Info("maxPrice: ", p.ans)
}

/*
	单调栈
	将当前读取的数据和栈内数据进行比较，
	如果当前数据比栈内数据大则输出，否则放入栈
	因此，栈内元素的值一定是单调的
*/
func (p *Arith_daily) Leetcode739() {
	stack := []int{}
	p.ans = make([]int, len(p.temp))
	for i := 0; i < len(p.temp); i++ {
		temp := p.temp[i]
		for len(stack) > 0 && temp > p.temp[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			p.ans[index] = i - index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	log.Info("ans: ", p.ans)
}

/*
	接雨水，暴力解法，根据当前柱子的高度h，横向两边查找比h高的柱子，
	取两边最高柱子的最小值，时间复杂度O(n^2)。
	该题可以优化，
	1、通过一次for循环遍历整个数组，通过单向流水机制，
	从左到右取最大值，从右到左取最小值，最后取相同索引下的最小值进行相加
	时间复杂度O(n)
	2、单调栈
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
