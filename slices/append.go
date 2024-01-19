package slices

const rateFactor = 2

func Append[T comparable](slice []T, elems ...T) []T {
	newCap := len(slice) + len(elems)
	if cap(slice) < newCap {
		if newCap < rateFactor*len(slice) {
			newCap = rateFactor * len(slice)
		}
		newSlice := make([]T, len(slice)+len(elems), newCap)
		copy(newSlice, slice)
		copy(newSlice[len(slice):], elems)
		return newSlice
	}
	slice = slice[:len(slice)+len(elems)]
	copy(slice[len(slice)-len(elems):], elems)
	return slice
}
