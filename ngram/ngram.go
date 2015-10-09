package ngrams

func ngrams(s string, n int) []string {
	var result[]string

	for i := 0; i < len(s)-n+1; i++ {
			result.Append(s[i : i+n])
	}
	return result
}
