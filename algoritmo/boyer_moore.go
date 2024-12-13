package algorithms

func buildBadCharTable(pattern string) map[rune]int {
	table := make(map[rune]int)
	pLen := len(pattern)
	for i, char := range pattern {
		table[char] = pLen - 1 - i
	}
	return table
}

func buildGoodSuffixTable(pattern string) []int {
	pLen := len(pattern)
	table := make([]int, pLen)
	lastPrefixPosition := pLen

	for i := pLen - 1; i >= 0; i-- {
		if isPrefix(pattern, i+1) {
			lastPrefixPosition = i + 1
		}
		table[pLen-1-i] = lastPrefixPosition - i + pLen - 1
	}

	for i := 0; i < pLen-1; i++ {
		lengthSuffix := suffixLength(pattern, i)
		table[lengthSuffix] = pLen - 1 - i + lengthSuffix
	}

	return table
}

func isPrefix(pattern string, p int) bool {
	sub := pattern[p:]
	return pattern[:len(sub)] == sub
}

func suffixLength(pattern string, p int) int {
	length := 0
	for i, j := p, len(pattern)-1; i >= 0 && pattern[i] == pattern[j]; i, j = i-1, j-1 {
		length++
	}
	return length
}

func BoyerMoore(text, pattern string) []int {
	if len(pattern) == 0 || len(text) == 0 {
		return nil
	}

	badCharTable := buildBadCharTable(pattern)
	goodSuffixTable := buildGoodSuffixTable(pattern)
	indices := []int{}

	tLen, pLen := len(text), len(pattern)
	for i := 0; i <= tLen-pLen; {
		j := pLen - 1
		for j >= 0 && pattern[j] == text[i+j] {
			j--
		}
		if j < 0 {
			indices = append(indices, i)
			i += goodSuffixTable[0]
		} else {
			badCharShift := badCharTable[rune(text[i+j])]
			if badCharShift == 0 {
				badCharShift = pLen
			}
			i += max(goodSuffixTable[j], badCharShift)
		}
	}

	return indices
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
