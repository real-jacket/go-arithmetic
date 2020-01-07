package longestCommonPrefix

import "unicode"

func GetLongestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == "" {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var str = strs[0]
	var i = 0
	for i < len(str) {
		if unicode.IsUpper(rune(str[i])) {
			return str[0:i]
		} else {
			for _, v := range strs {
				if v == "" || i >= len(v) || str[i] != v[i] {
					return str[0:i]
				}
			}
		}
		i++
	}
	return str[0:i]
}
