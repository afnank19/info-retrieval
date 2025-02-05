package sliceutil

// Currently doing any, but i think it would be better to have a type
func IntersectionM[T any](a, b []T) []T {
	var commonItems []T

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {

		}
	}

	return commonItems
}

func Intersection(a, b []string) []string {
	var commonItems []string

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				commonItems = append(commonItems, a[i])
			}
		}
	}

	return commonItems
}

// Some times, code is bestowed by God upon you
// This is one of those times
// It's simple, but don't ask me how it works
// Putting lots of trust in this
func Union(a, b []string) []string {
	result := a
	for i := 0; i < len(b); i++ {
		exists := false
		for j := 0; j < len(a); j++ {
			if b[i] == a[j] {
				exists = true
			}
		}
		if exists == false {
			result = append(result, b[i])
		}
	}

	return result
}

func Difference(a, b []string) []string {
	var result []string
	for i := 0; i < len(a); i++ {
		var foundInB = false
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				foundInB = true
			}
		}

		if !foundInB {
			result = append(result, a[i])
		}
	}
	return result
}

// Basically U - N, where U is the universal set
// This could probably turned into an O(N) function by using a map of some sorts
// My stupid self didn't recognise that the above func is exactly the same lol
// anywho
func UniversalDifference(universal, b []string) []string {
	var result []string

	for i := 0; i < len(universal); i++ {
		var found = false
		for j := 0; j < len(b); j++ {
			if universal[i] == b[j] {
				found = true
			}
		}
		if !found {
			result = append(result, universal[i])
		}
	}

	return result
}
