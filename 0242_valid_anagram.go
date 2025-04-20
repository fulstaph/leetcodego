package leetcodego

const englishAlphabetLetterCount = 26

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	lettersCount := make(map[rune]int, englishAlphabetLetterCount)
	for _, c := range s {
		lettersCount[c] += 1
	}

	for _, c := range t {
		lettersCount[c] -= 1
	}

	for _, v := range lettersCount {
		if v > 0 {
			return false
		}
	}

	return true
}
