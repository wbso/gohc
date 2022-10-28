package str

func SliceContains[T comparable](haystack []T, needle T) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}
