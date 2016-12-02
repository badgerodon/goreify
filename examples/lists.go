package examples

import "github.com/badgerodon/goreify/generics"

// Contains returns whether or not the list contains the element
// reify:
//   types:
//     T1: numeric
func Contains(haystack []generics.T1, needle generics.T1) bool {
	for _, el := range haystack {
		if generics.Equal(el, needle) {
			return true
		}
	}
	return false
}
