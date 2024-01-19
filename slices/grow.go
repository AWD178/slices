package slices

import (
	"errors"
)

func Grow[Arr ~[]T, T any](arr Arr, n int) (Arr, error) {
	if n < cap(arr) {
		return arr, errors.New(`need capacity must be greater than current array capacity`)
	}
	newArray := make(Arr, n)
	copy(newArray, arr)
	return newArray, nil
}
