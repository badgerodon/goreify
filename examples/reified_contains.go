package examples

// ContainsString returns whether or not the list contains the element
func ContainsString(haystack []string, needle string) bool {
	for _, el := range haystack {
		if el == needle {
			return true
		}
	}
	return false
}
