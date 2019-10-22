package algo

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

//最长不重复子串的长度
func LengthOfLongestSubstring(s string) int {
	if utf8.RuneCountInString(s) == 0 {
		return 0
	}
	var left float64 = 0
	var max float64 = 0
	data := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if j, ok := data[s[i]]; ok {
			left = math.Max(float64(left), float64(j+1))
		}
		data[s[i]] = i
		max = math.Max(float64(max), float64(float64(i)-left+1))
	}
	return int(max)
}

//数组中两个数的和等于target
func TwoSum(nums []int, target int) []int {
	var data = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		if re, ok := data[t]; ok {
			return []int{re, i}
		}
		data[nums[i]] = i
	}
	return nil
}

//最长回文子串 中心扩散算法
func LongestPalindrome(s string) string {
	var start, end = 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		length := math.Max(float64(len1), float64(len2))
		if int(length) > end-start {
			start = i - (int(length)-1)/2
			end = i + int(length)/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	var L = left
	var R = right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

//Z字型变换
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	step := 2*numRows - 2
	index := 0
	length := len(s)
	add := 0
	ret := ""
	for i := 0; i < numRows; i++ {
		index = i
		add = i * 2
		for index < length {
			ret += string(s[index])
			add = step - add
			if i == 0 || i == numRows-1 {
				index += step
			} else {
				index += add
			}
		}
	}
	return ret
}

//整数翻转
func Reverse(x int) int {
	ans := 0
	for x != 0 {
		pop := x % 10
		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if ans < math.MinInt32/10 || (ans == math.MinInt32/10 && pop < -8) {
			return 0
		}
		ans = ans*10 + pop
		x = x / 10
	}
	return ans
}

//回文数
func IsPalindrome(x int) bool {
	//这里要考虑进去
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	temp := x
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	return temp == y
}

//盛最多水的容器
//这里是高度矮的乘以长度得到面积
func MaxArea(height []int) int {
	var maxArea float64 = 0
	l := 0
	r := len(height) - 1
	for l < r {
		maxArea = math.Max(float64(maxArea), math.Min(float64(height[l]), float64(height[r]))*float64(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return int(maxArea)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := &ListNode{}
	tempNode := resultNode
	sum := 0
	for l1 != nil || l2 != nil || sum != 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tempNode.Next = &ListNode{sum % 10, nil}
		tempNode = tempNode.Next
		sum /= 10
	}
	return tempNode.Next
}

//数字转罗马数字
func IntToRoman(num int) string {
	ret := ""
	if num < 1 || num > 3999 {
		return ret
	}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	letter := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(nums); {
		if num >= nums[i] {
			num -= nums[i]
			ret += letter[i]
		} else {
			i++
		}
	}
	return ret
}

func RomanToInt(s string) int {
	a := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	number := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && a[s[i:i+1]] < a[s[i+1:i+2]] {
			number -= a[s[i:i+1]]
		} else {
			number += a[s[i:i+1]]
		}
	}
	return number
}

//最长公共子串
func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	ans := strs[0]
	for i := 0; i < length; i++ {
		j := 0
		for ; j < len(ans) && j < len(strs[i]); j++ {
			if ans[j] != strs[i][j] {
				break
			}
		}
		ans = ans[0:j]
		if ans == "" {
			return ans
		}
	}
	return ans
}

type ids []int

func (id ids) Less(i, j int) bool {
	return id[i] < id[j]
}

func (id ids) Swap(i, j int) {
	id[i], id[j] = id[j], id[i]
}

func (id ids) Len() int {
	return len(id)
}

//三数之和
func ThreeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 || nums == nil {
		return nil
	}
	var ans [][]int
	var id = ids(nums)
	//先排序
	//如果这个数和前一个数相同 则跳过
	//如果第一个数大于0  三数之和不可能等于0
	sort.Sort(id)
	var (
		l int
		r int
	)
	for i := 0; i < length; i++ {
		if id[i] > 0 {
			break
		}
		//去重
		if i > 0 && id[i] == id[i-1] {
			continue
		}
		l = i + 1
		r = length - 1
		for l < r {
			sum := id[i] + id[l] + id[r]
			if sum == 0 {
				ans = append(ans, []int{id[i], id[l], id[r]})
				//去重
				for l < r && id[l] == id[l+1] {
					l++
				}
				for l < r && id[r] == id[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			}
		}
	}
	return ans
}

//最接近的三数之和
func ThreeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length < 3 || nums == nil {
		return 0
	}
	var id = ids(nums)
	var ans int
	var lite = 1<<63 - 1
	//比小  不比大
	sort.Sort(id)
	var (
		l int
		r int
	)
	for i := 0; i < length; i++ {
		l = i + 1
		r = length - 1
		for l < r {
			sum := id[i] + id[l] + id[r]
			ll := math.Abs(float64(target - sum))
			if int(ll) < lite {
				lite = int(ll)
				ans = sum
			}
			if target == sum {
				return ans
			} else if target > sum {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}

//解题思路
/**
1.先将第一个数字对应的字符串的字符加入ans
2.继续遍历  将之前的ans存入temp中  将ans置空
3.遍历temp,再遍历digits对应的字符 ,然后相加存入ans
 */
func letterCombinations(digits string) []string {
	var ans []string
	set := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	for i, v := range digits {
		if i == 0 {
			for _, vt := range set[v] {
				ans = append(ans, string(vt))
			}
		} else {
			temp := ans
			ans = nil
			for _, v1 := range temp {
				for _, v2 := range set[v] {
					ans = append(ans, v1+string(v2))
				}
			}
		}
	}
	return ans
}

//删除链表的倒数第N个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	node := &ListNode{Next: head}
	start, end := node, node
	for n != 0 {
		end = end.Next
		n--
	}
	for end.Next != nil {
		start = start.Next
		end = end.Next
	}
	start.Next = start.Next.Next
	return node.Next
}

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return nil
	}
	//排序
	sort.Ints(nums)
	var (
		l   int
		r   int
		ans [][]int
	)
	for i := 0; i < length-3; i++ {
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l = j + 1
			r = length - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r-1 && nums[l] == nums[l+1] {
						l++
					}
					for r < length-1 && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return ans
}

//有效的括号
func isValid(s string) bool {
	var stack []string
	//映射表
	foreg := map[string]string{")": "(", "]": "[", "}": "{"}
	for _, x := range s {
		if string(x) == "(" || string(x) == "[" || string(x) == "{" {
			stack = append(stack, string(x))
		} else if string(x) == ")" || string(x) == "]" || string(x) == "}" {
			//栈不为空
			if len(stack) != 0 && stack[len(stack)-1] == foreg[string(x)] {
				stack = stack[0:len(stack)-1]
			}else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

//合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prevHead := ListNode{}
	prev := &prevHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		}else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return prevHead.Next
}

func generateParenthesis(n int) []string{
	var ans []string
	var dfs func(s string,left int,right int)
	dfs = func(s string,left,right int) {
		if left > right {
			return
		}
		if left == 0 && right == 0 {
			ans = append(ans,s)
			return
		}
		if left > 0 {
			 dfs(s + "(",left-1,right)
		}
		if right > 0 {
			dfs(s + ")",left,right-1)
		}
	}
	dfs("",n,n)
	return ans
}

func removeElement(nums []int, val int) int {
	if nums == nil {
		return 0
	}
	var ans = 0
	for _,n := range nums {
		if n != val {
			nums[ans] = n
			ans++
		}
	}
	return ans
}

func FindSubstring(s string, words []string) []int {
	//需要由words中所有单词组合成的字符串
	//遍历子串是否都是由words组成
	w2num := make(map[string]uint8)
	allNum := len(words)
	if allNum == 0 {
		return []int{}
	}
	length :=  len(words[0])
	if length == 0 {
		return []int{}
	}
	for _, str := range words {
		w2num[str] += 1
	}
	tmp := make([]int, 0, 30)
	re := &tmp
	for i:=0; i<=len(s)-allNum*length && i<length; i++ {
		isSub(s[i:], w2num, length, allNum, re, 0, 0, i)
	}
	return *re
}


func isSub(s string, w2num map[string]uint8, length int, allNum int, re *[]int, start int, end int, offset int) {
	if len(s[end:]) - allNum*length < 0 {
		return
	}
	tmp, ok := w2num[s[end:end+length]]
	if ok && tmp>0 {
		//能进下一状态
		w2num[s[end:end+length]] -= 1
		allNum -= 1
		end += length
		if allNum == 0 {
			*re = append(*re, start+offset)
			w2num[s[start:start+length]] += 1
			allNum += 1
			start += length
			isSub(s, w2num, length, allNum, re, start, end, offset)
			start -= length
			allNum -= 1
			w2num[s[start:start+length]] -= 1
		} else {
			isSub(s, w2num, length, allNum, re, start, end, offset)
		}
		end -= length
		allNum += 1
		w2num[s[end:end+length]] += 1
	} else {
		if end > start {
			w2num[s[start:start+length]] += 1
			allNum += 1
			start += length
			isSub(s, w2num, length, allNum, re, start, end, offset)
			start -= length
			allNum -= 1
			w2num[s[start:start+length]] -= 1
		} else {
			isSub(s, w2num, length, allNum, re, start+length, end+length, offset)
		}
	}
}

//下一个排列
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i > 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		swap(nums,i,j)
	}
	reverse(nums,i+1)
}

func reverse(nums []int, start int) {
	l := start
	r := len(nums) - 1
	for l < r {
		swap(nums,l,r)
		l++
		r--
	}
}

func swap(nums []int, i, j int) {
	nums[i],nums[j] = nums[j],nums[i]
}

func longestValidParentheses(s string) int {
	var stack []interface{}
	stack = append(stack, -1)
	var max float64
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" {
			stack = append(stack, i)
		}else {
			if len(stack) > 0 {
				top := stack[len(stack)-1].(int)
				stack = stack[:len(stack)-1]
				max = math.Max(max,float64(i-top))
			}else {
				stack = append(stack,i)
			}
		}
	}
	return int(max)
}

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	var ans int
	length := len(nums)
	for i := 0; i < length;i++ {
		if i == (length-1) && target > nums[i] {
			ans = i + 1
			break
		}
		if nums[i] >= target {
			ans = i
			break
		}
	}
	return ans
}

func SearchInsert2(nums []int,target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	if nums[right] < target {
		return right + 1
	}

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid +1
		}else {
			right = mid - 1
		}
	}
	return left
}

func IsValidSudoku(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')  // 当前数字的 二进制数位 位置
			bi := i/3 + j/3*3  // 3x3的块索引号
			if (row[i] & cur) | (col[j] & cur) | (block[bi] & cur) != 0 { // 使用与运算查重
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}

func FirstMissingPositive(nums []int) int {
	sort.Ints(nums)
	count := 1
	for i:=0;i < len(nums);i++ {
		if nums[i] <= 0 {
			continue
		}
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if count != nums[i] {
			break
		}
		count++
	}
	return count
}

func ReverseWords(s string) string {
	if s == "" {
		return ""
	}
	var ans = ""
	temp := ""
	for i := 0; i < len(s);i++ {

		if i == len(s) - 1 {
			temp = string(s[i]) + temp
			ans = ans + " " + temp
			break
		}

		if string(s[i]) == " " {
			ans = ans + " " + temp
			temp = ""
			continue
		}
		temp = string(s[i]) + temp
	}
	return ans[1:]
}

func sa(s string) string {

	ss := []byte(s)

	for i, j := 0, 0; j < len(ss); j++ {

		if ss[j] == ' ' {
			reverseWord(ss, i, j - 1)
			i = j + 1
		} else if j == len(ss) - 1 {
			reverseWord(ss, i, j)
			break
		}
	}
	return string(ss)
}

func reverseWord(ss []byte, i int, j int) {
	if i >= j {
		return
	}
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
}

func grayCode(n int) []int {
	if n <= 0 {
		return nil
	}
	var ans []int
	var num = uint(n)
	t := 1 << num
	for i:= 0;i < t;i++ {
		ans = append(ans,i ^ (i>>1))
	}
	return ans
}

func SpiralOrder(matrix [][]int) []int {
	var ans []int
	if matrix == nil  || len(matrix) == 0 {
		return ans
	}
	m := len(matrix) //列
	n := len(matrix[0]) //行
	var count int
	if m < n {
		count = (m +1) / 2
	}else {
		count = (n+1)/2
	}
	i := 0
	for i < count {
		//从左往右 第i个到第n-i个
		for j := i;j < n - i;j++ {
			ans = append(ans,matrix[i][j])
		}
		//从上到下 第i+1行到第m-i行  第n-1-i列(定值)
		for j := i+1;j < m-i; j++{
			ans = append(ans,matrix[j][n-i-1])
		}
		//从右到左 第n-1-(i+1)个到 i个 第m-i-1行(定值)
		for j:=n-1-(i+1);j>=i && m-i-1 != i;j-- {
			ans = append(ans,matrix[m-1-i][j])
		}
		//从下到上 第m-1-(i+1)行 到i+1行 第i列
		for j := m -1 - (i+1);j>=i+1 && n-i-1 != i;j-- {
			ans = append(ans,matrix[j][i])
		}
		i++
	}
	return ans
}

func MyAtoi(str string) int {
	str = strings.TrimSpace(str)
	if str == "" || (len(str) == 1 && (str < "0" || str > "9")) {
		return 0
	}

	flag := ""

	if string(str[0]) == "-" {
		flag = "-"
		str = str[1:]
	}else if string(str[0]) == "+"  {
		str = str[1:]
	}

	res := "0"
	for i:=0;i<len(str);i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		res += string(str[i])
	}
	res = flag + res

	ans,err := strconv.ParseInt(res,10,32)

	const MaxUint = ^uint32(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	if err != nil {
		if flag == "" {
			return MaxInt
		}
		return MinInt
	}

	return int(ans)

}