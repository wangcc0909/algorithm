package algo

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 8, 9, 11, 13, 15, 17, 18, 21}
	var target = 31
	result := TwoSum(nums, target)
	if result != nil {
		fmt.Println(result)
		return
	}
	t.Errorf("not found %d", target)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "asdjflaksjada"
	result := LongestPalindrome(s)
	t.Log(result)
}

func TestConvert(t *testing.T) {
	s := "LEETCODEISHIRING"
	result := Convert(s, 5)
	t.Log(result)
}

func TestReverse(t *testing.T) {
	i := -1469989897
	result := Reverse(i)
	t.Log(result)
}

func TestIsPalindrome(t *testing.T) {
	i := 10
	result := IsPalindrome(i)
	t.Log(result)
}

func TestMaxArea(t *testing.T) {
	r := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := MaxArea(r)
	t.Log(result)
}

func TestIntToRoman(t *testing.T) {
	result := IntToRoman(10)
	t.Log(result)
}

func TestLongestCommonPrefix(t *testing.T) {
	r := []string{"lasjf", "palf", "plasfgasj"}
	result := LongestCommonPrefix(r)
	t.Log(result)
}

func TestThreeSum(t *testing.T) {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	result := ThreeSum(nums)
	t.Log(result)
}

func TestThreeSumClosest(t *testing.T) {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	result := ThreeSumClosest(nums, 4)
	t.Log(result)
}

func TestFindSubstring(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "word"}
	result := FindSubstring(s, words)
	t.Log(result)
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	result := SearchInsert(nums, target)
	t.Log(result)
}

func TestIsValidSudoku(t *testing.T) {
	var board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	result := IsValidSudoku(board)
	t.Log(result)
}

func TestFirstMissingPositive(t *testing.T) {
	var nums = []int{7,8,9,11,12}
    result := FirstMissingPositive(nums)
    t.Log(result)
}

func TestReverseWords(t *testing.T) {
	var s = "Let's take LeetCode contest"
	r := ReverseWords(s)
	t.Log(r)
}

/*func TestInvalidTransactions(t *testing.T) {
	transactions := []string{
		"lee,886,1785,beijing","alex,763,1157,amsterdam","lee,277,129,amsterdam","bob,770,105,amsterdam","lee,603,926,amsterdam","chalicefy,476,50,budapest","lee,924,859,barcelona","alex,302,590,amsterdam","alex,397,1464,barcelona","bob,412,1404,amsterdam","lee,505,849,budapest",
	}

	r := InvalidTransactions(transactions)
	t.Log(r)
}*/

func TestMyAtoi(t *testing.T) {
	s := "9223372036854775808"
	r := MyAtoi(s)
	t.Log(r)
}
