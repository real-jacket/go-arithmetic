package palindromic

import "testing"

func TestIsPalindrome(t *testing.T) {
	num1 := 1002
	num2 := 10001
	num3 := 0
	num4 := 456
	num5 := 1001001
	if IsPalindrome(num1) != false {
		t.Errorf("'%v' test faild", num1)
	}
	if IsPalindrome(num2) != true {
		t.Errorf("'%v' test faild", num2)
	}
	if IsPalindrome(num3) != true {
		t.Errorf("'%v' test faild", num2)
	}
	if IsPalindrome(num4) != false {
		t.Errorf("'%v' test faild", num4)
	}
	if IsPalindrome(num5) != true {
		t.Errorf("'%v' test faild", num5)
	}
}
