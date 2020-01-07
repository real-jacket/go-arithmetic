package longestCommonPrefix

func GetLongestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == "" {
		return ""
	}
	str := strs[0]
	for i := 0; i < len(str); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || str[i] != strs[j][i] {
				return str[0:i]
			}
		}
	}
	return str
}
