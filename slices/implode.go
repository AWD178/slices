package slices

func Implode[Arr ~[]T, T any](arr ...Arr) Arr {
	newSlice := make(Arr, 0)
	for _, s := range arr {
		newSlice = append(newSlice, s...)
	}
	return newSlice
}
