package slices_arrays

import (
	"reflect"
	"slices_arrays/slices"
	"testing"
)

func BenchmarkImplode(b *testing.B) {
	sv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		if reflect.DeepEqual(slices.Implode([]int{1, 2, 3}, []int{4, 5, 6}, []int{7}, []int{8, 9}), sv) == false {
			b.Error(`slices not equals`)
		}
	}
}

func BenchmarkImplodeW(b *testing.B) {
	sv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		if reflect.DeepEqual(slices.ImplodeWithoutCalc([]int{1, 2, 3}, []int{4, 5, 6}, []int{7}, []int{8, 9}), sv) == false {
			b.Error(`slices not equals`)
		}
	}
}
