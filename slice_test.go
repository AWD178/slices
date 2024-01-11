package slices_arrays

import (
	"reflect"
	"slices_arrays/slices"
	"testing"
)

func TestAppend(t *testing.T) {
	sliceFirst := []int{1, 2, 3}
	sliceSecond := []int{4, 5, 6}

	sliceFirstCap := cap(sliceFirst)
	sliceFirst = append(sliceFirst, sliceSecond...)
	if cap(sliceFirst)/sliceFirstCap != 2 {
		t.Errorf(`failed inc capacity`)
	}
	return
}

func TestGrow(t *testing.T) {
	var err error

	originalSlice := []int{1, 2, 3}
	originalSliceCap := cap(originalSlice)

	originalSlice, err = slices.Grow(originalSlice, originalSliceCap*2)
	if err != nil {
		t.Error(err)
		return
	}

	if cap(originalSlice) != originalSliceCap*2 {
		t.Errorf(`incorrect capacity`)
	}
	return
}

func TestImplode(t *testing.T) {
	sv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if reflect.DeepEqual(slices.Implode([]int{1, 2, 3}, []int{4, 5, 6}, []int{7}, []int{8, 9}), sv) == false {
		t.Error(`slices not equals`)
	}
}
