package leetcodego

// https://leetcode.com/problems/reverse-vowels-of-a-string
var vowels = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func isVowel(c rune) bool {
	_, ok := vowels[c]
	return ok
}

func reverseVowels(s string) string {
	runes := []rune(s)
	start, end := 0, len(s)-1
	for start < end {
		if !isVowel(runes[start]) {
			start++
			continue
		}

		if !isVowel(runes[end]) {
			end--
			continue
		}

		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	return string(runes)
}
