package examples

func Contains_string(haystack []string, needle string) bool {
	for _, el := range haystack {
		if el == needle {
			return true
		}
	}
	return false
}
