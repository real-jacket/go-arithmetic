package longestCommonPrefix

import "testing"

func TestGetLongestCommonPrefix(t *testing.T) {
	strs1 := []string{"flower", "flow", "flight"}
	str1 := GetLongestCommonPrefix(strs1)
	if str1 != "fl" {
		t.Errorf("%s failed", str1)
	}
	strs2 := []string{"dog", "racecar", "car"}
	str2 := GetLongestCommonPrefix(strs2)
	if str2 != "" {
		t.Errorf("%s failed", str2)
	}
	strs3 := []string{}
	str3 := GetLongestCommonPrefix(strs3)
	if str3 != "" {
		t.Errorf("%s failed", str3)
	}
	strs4 := []string{"aa", "a"}
	str4 := GetLongestCommonPrefix(strs4)
	if str3 != "" {
		t.Errorf("%s failed", str4)
	}
}
