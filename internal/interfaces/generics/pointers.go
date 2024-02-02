package generics

// ToPointer creates a new pointer to the given value and returns it.
// It is a generic function that works with any type (T) that supports pointer creation.
func ToPointer[T any](v T) *T {
	return &v
}

// SetIfNotNil copies the value from the pointer to the target variable if the pointer is not nil.
// It is a generic function that works with any type (T) that supports pointer dereferencing.
// If the pointer is nil, the target variable remains unchanged.
// Use only for allocated target variables.
func SetIfNotNil[T any](ptr *T, target *T) {
	if ptr != nil {
		if target == nil {
			target = ptr
			return
		}
		*target = *ptr
	}
}

// DereferenceOrDefault safely dereferences a pointer and returns its value.
// It is a generic function that works with any type (T) that supports pointer dereferencing.
func DereferenceOrDefault[T any](p *T) (defaultValue T) {
	if p != nil {
		defaultValue = *p
	}
	return
}
