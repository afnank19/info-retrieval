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
