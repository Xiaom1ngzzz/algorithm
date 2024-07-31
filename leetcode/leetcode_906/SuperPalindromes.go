package leetcode906

import (
	"math"
	"sort"
	"strconv"
)

// 如果一个正整数自身是回文数，而且它也是一个回文数的平方，那么我们称这个数为超级回文数。
// 现在，给定两个正整数 L 和 R （以字符串形式表示），
// 返回包含在范围 [L, R] 中的超级回文数的数目。
// 1 <= len(L) <= 18
// 1 <= len(R) <= 18
// L 和 R 是表示 [1, 10^18) 范围的整数的字符串

// [left, right]有多少超级回文数
func superpalindromesInRange(left string, right string) int {
	l, _ := strconv.ParseInt(left, 10, 64)
	r, _ := strconv.ParseInt(right, 10, 64)

	// x根号，范围limit
	limit := int64(math.Sqrt(float64(r)))
	// seed ： 枚举量很小，10^8 -> 10^9 => 10^5
	// seed ： 奇数长度回文，偶数长度回文
	// seed := 1
	var seed int64 = 1
	// num ： 根号x，num^2
	// num := 0
	var num int64 = 0
	ans := 0

	for {
		// seed生成偶数长度回文数字
		// 123 -> 123321
		num = evenEnlarge(seed)
		if check(num*num, l, r) {
			ans++
		}
		// seed生成奇数长度回文数字
		// 123 -> 12321
		num = oddEnlarge(seed)
		if check(num*num, l, r) {
			ans++
		}
		seed++

		if num >= limit {
			break
		}
	}
	return ans
}

// 根据种子扩充到偶数长度的回文数字并返回
func evenEnlarge(seed int64) int64 {
	ans := seed

	for seed != 0 {
		ans = ans*10 + seed%10
		seed /= 10
	}
	return ans
}

// 根据种子扩充到奇数长度的回文数字并返回
func oddEnlarge(seed int64) int64 {
	ans := seed

	seed /= 10
	for seed != 0 {
		ans = ans*10 + seed%10
		seed /= 10
	}
	return ans
}

// 判断ans是不是属于[l,r]且是不是回文数字
func check(num, l, r int64) bool {
	return num <= r && num >= l && isPalindrome(num)
}

func isPalindrome(num int64) bool {
	var offset int64 = 1

	// 为了防止溢出
	for num/offset >= 10 {
		offset *= 10
	}

	// 首尾判断
	for num != 0 {
		if num/offset != num%10 {
			return false
		}
		num = (num % offset) / 10
		offset /= 100
	}
	return true
}

var records = []int64{
	1,
	4,
	9,
	121,
	484,
	10201,
	12321,
	14641,
	40804,
	44944,
	1002001,
	1234321,
	4008004,
	100020001,
	102030201,
	104060401,
	121242121,
	123454321,
	125686521,
	400080004,
	404090404,
	10000200001,
	10221412201,
	12102420121,
	12345654321,
	40000800004,
	1000002000001,
	1002003002001,
	1004006004001,
	1020304030201,
	1022325232201,
	1024348434201,
	1210024200121,
	1212225222121,
	1214428244121,
	1232346432321,
	1234567654321,
	4000008000004,
	4004009004004,
	100000020000001,
	100220141022001,
	102012040210201,
	102234363432201,
	121000242000121,
	121242363242121,
	123212464212321,
	123456787654321,
	400000080000004,
	10000000200000001,
	10002000300020001,
	10004000600040001,
	10020210401202001,
	10022212521222001,
	10024214841242001,
	10201020402010201,
	10203040504030201,
	10205060806050201,
	10221432623412201,
	10223454745432201,
	12100002420000121,
	12102202520220121,
	12104402820440121,
	12122232623222121,
	12124434743442121,
	12321024642012321,
	12323244744232321,
	12343456865434321,
	12345678987654321,
	40000000800000004,
	40004000900040004,
	1000000002000000001,
	1000220014100220001,
	1002003004003002001,
	1002223236323222001,
	1020100204020010201,
	1020322416142230201,
	1022123226223212201,
	1022345658565432201,
	1210000024200000121,
	1210242036302420121,
	1212203226223022121,
	1212445458545442121,
	1232100246420012321,
	1232344458544432321,
	1234323468643234321,
	4000000008000000004,
}

func superpalindromesInRange2(left string, right string) int {
	l, _ := strconv.ParseInt(left, 10, 64)
	r, _ := strconv.ParseInt(right, 10, 64)

	i := 0
	for ; i < len(records); i++ {
		if records[i] >= l {
			break
		}
	}
	j := len(records) - 1
	for ; j >= 0; j-- {
		if records[j] <= r {
			break
		}
	}
	return j - i + 1
}

func collect() []int64 {
	var l int64 = 1
	var r int64 = math.MaxInt64
	var limit int64 = int64(math.Sqrt(float64(r)))
	var seed int64 = 1
	var enlarge int64 = 0
	ans := make([]int64, 0)

	for {
		enlarge = evenEnlarge(seed)
		if check(enlarge*enlarge, l, r) {
			ans = append(ans, enlarge*enlarge)
		}

		enlarge = oddEnlarge(seed)
		if check(enlarge*enlarge, l, r) {
			ans = append(ans, enlarge*enlarge)
		}
		seed++

		if enlarge >= limit {
			break
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}
