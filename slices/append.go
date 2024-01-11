package slices

func Append[T comparable](slice []T, elems ...T) []T {
	if cap(slice) < len(slice)+len(elems) {
		newSlice := make([]T, len(slice), (len(slice)+len(elems))*2)
		copy(newSlice, slice)
		return append(newSlice, elems...)
	}
	return append(slice, elems...)
}
