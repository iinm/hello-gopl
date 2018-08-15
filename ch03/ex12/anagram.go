package main

import "fmt"

func main() {
	fmt.Println(isAnagran("anagrams", "ars magna"))
	fmt.Println(isAnagran("ars magna", "anagrams"))
	fmt.Println(isAnagran("ars magna", "anagram"))
	fmt.Println(isAnagran("anagram", "ars magna"))
}

func isAnagran(s1, s2 string) bool {
	if s1 == `` || s2 == `` {
		return false
	}

	// 文字の頻度
	runeCount1 := countRunes(s1)
	runeCount2 := countRunes(s2)
	for r, count1 := range runeCount1 {
		if count2, ok := runeCount2[r]; !ok && count1 != count2 {
			return false
		}
	}
	for r, count2 := range runeCount2 {
		if count1, ok := runeCount1[r]; !ok && count1 != count2 {
			return false
		}
	}
	return true
}

func countRunes(s string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range s {
		if r == ' ' {
			continue
		}
		count[r]++
	}
	return count
}
