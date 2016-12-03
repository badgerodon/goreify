package examples

// Contains_string returns whether or not the list contains the element
func Contains_string(haystack []string, needle string) bool {
	for _, el := range haystack {
		if el == needle {
			return true
		}
	}
	return false
}
