package valid

func IsValid(s string) bool {
	if s == "" {
		return true
	}
	var mid []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			mid = append(mid, s[i])
		}
		if s[i] == ')' {
			if len(mid) != 0 && mid[len(mid)-1] == '(' {
				mid = mid[:len(mid)-1]
			} else {
				return false
			}
		}
		if s[i] == ']' {
			if len(mid) != 0 && mid[len(mid)-1] == '[' {
				mid = mid[:len(mid)-1]
			} else {
				return false
			}
		}
		if s[i] == '}' {
			if len(mid) != 0 && mid[len(mid)-1] == '{' {
				mid = mid[:len(mid)-1]
			} else {
				return false
			}
		}
	}
	return len(mid) == 0
}

func FisValid(s string) bool {
	dic := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	var stack []string
	for _, v := range s {
		if _, ok := dic[string(v)]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != dic[string(v)] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(v))
		}
	}
	return len(stack) == 0
}
