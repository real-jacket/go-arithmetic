package palindromic

// IsPalindrome 回文函数
func IsPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
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
