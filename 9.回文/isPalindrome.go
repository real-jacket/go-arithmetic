package palindromic

import "strconv"

// IsPalindrome 回文函数
func IsPalindrome(x int) bool {
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}
	var r, v int
	r = x
	for v < r {
		v = v*10 + r%10
		r = r / 10
	}
	return (v/10 == r) || (v == r)
}

func IsPalindromeMath(x int) bool {
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}
	var div int = 1
	for x/div >= 10 {
		div *= 10
	}
	var left, right int
	for x > 0 {
		left = x / div
		right = x % 10
		if left != right {
			return false
		}
		x = (x % div) / 10
		div /= 100
	}
	return true
}

func IsPalindromeString(x int) bool {
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}
	var y string = strconv.Itoa(x)
	var head, tail int
	for head <= tail {
		tail = len(y) - head - 1
		if y[head] != y[tail] {
			return false
		}
		head++
	}
	return true
}
