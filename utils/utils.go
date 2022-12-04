package utils

func ContainsInStringArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsInIntArray(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsRuneInString(s string, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
