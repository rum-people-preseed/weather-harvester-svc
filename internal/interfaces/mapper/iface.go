package mapper

type BaseMapper[T, E any] interface {
	ToDTO(T) E
	ToManyDTO([]T) []E
	FromDTO(E) T
	FromManyDTO([]E) []T
}
