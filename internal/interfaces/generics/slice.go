package generics

// IsInSlice godoc
func IsInSlice[T comparable](val T, list []T) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

// RemoveIndex remove removes the first instance of s in slice
func RemoveIndex[T comparable](s T, slice []T) []T {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
