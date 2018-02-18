package stringutils

// NOTE: PACKAGE SCOPE:
// This function is not visible outside of the stringutils package.
// However, it is visible to other files within the package itself.
// (Starts with lower case letter)
func reverseV2(i_str string) string {
	r := []rune(i_str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
