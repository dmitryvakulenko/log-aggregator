package comparator

func Levenshtein(x, y string) int {
	s1, s2 := []rune(x), []rune(y)
	len1, len2 := len(s1), len(s2)

	if len1 == 0 {
		return len2
	}

	if len2 == 0 {
		return len1
	}

	var cost int
	if s1[len1 - 1] == s2[len2 - 1] {
		cost = 0
	} else {
		cost = 1
	}

	shortX, shortY := string(s1[0:len1 - 1]), string(s2[0:len2 - 1])
	return min(Levenshtein(x, shortY) + 1, Levenshtein(shortX, y) + 1, Levenshtein(shortX, shortY) + cost)
}

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	}

	if y < x && y < z {
		return y
	}

	return z
}

func initMatrix(strLen, colLen int) [][]int {
	d := make([][]int, strLen)
	for i := 0; i < strLen; i++ {
		d[i] = make([]int, colLen)
	}

	return d
}