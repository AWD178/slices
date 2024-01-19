package slices_arrays

import (
	"fmt"
	"reflect"
	"slices_arrays/slices"
	"testing"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		a       func() []int
		b       func() []int
		wantLen int
		wantCap int
	}{
		{
			a: func() []int {
				return nil
			},
			b: func() []int {
				return nil
			},
			wantLen: 0,
			wantCap: 0,
		},
		{
			a: func() []int {
				return []int{}
			},
			b: func() []int {
				return []int{}
			},

			wantLen: 0,
			wantCap: 0,
		},
		{
			a: func() []int {
				return nil
			},
			b: func() []int {
				return []int{1}
			},

			wantLen: 1,
			wantCap: 1,
		},
		{
			a: func() []int {
				return nil
			},
			b: func() []int {
				return []int{1, 2}
			},
			wantLen: 2,
			wantCap: 2,
		},
		{
			a: func() []int {
				return nil
			},
			b: func() []int {
				return []int{1, 2, 3}
			},
			wantLen: 3,
			wantCap: 4,
		},
		{
			a: func() []int {
				return []int{1, 2}
			},
			b: func() []int {
				return []int{1, 2, 3}
			},
			wantLen: 5,
			wantCap: 8,
		},
		{
			a: func() []int {
				r := make([]int, 0, 5)
				r = append(r, 1)
				r = append(r, 2)
				return r
			},
			b: func() []int {
				return []int{3, 4, 5}
			},
			wantLen: 5,
			wantCap: 5,
		},
		{
			a: func() []int {
				r := make([]int, 0, 5)
				r = append(r, 1)
				r = append(r, 2)
				r = append(r, 3)
				return r
			},
			b: func() []int {
				return []int{4, 5, 6}
			},
			wantLen: 6,
			wantCap: 10,
		},
		{
			a: func() []int {
				r := make([]int, 0, 3)
				r = append(r, 1)
				r = append(r, 2)
				return r
			},
			b: func() []int {
				return []int{3, 4, 5}
			},
			wantLen: 5,
			wantCap: 6,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("len=%d cap=%d", tt.wantLen, tt.wantCap), func(t *testing.T) {
			got := slices.Append(tt.a(), tt.b()...)
			if len(got) != tt.wantLen {
				t.Errorf("got len %d; want %d", len(got), tt.wantLen)
			}
			if cap(got) != tt.wantCap {
				t.Errorf("got cap %d; want %d", cap(got), tt.wantCap)
			}
		})
	}
}

func TestGrow(t *testing.T) {
	b := make([]int, 256)

	tests := []struct {
		a       []int
		n       int
		wantCap int
	}{
		{
			a:       []int{1, 2, 3},
			n:       10,
			wantCap: 10,
		},
		{
			a:       nil,
			n:       1001,
			wantCap: 1001,
		},
		{
			a:       []int{},
			n:       1002,
			wantCap: 1002,
		},

		{
			a:       []int{1, 2, 3},
			n:       1000,
			wantCap: 1000,
		},
		{
			a:       b,
			n:       300,
			wantCap: 300,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("cap=%d", tt.wantCap), func(t *testing.T) {
			got, err := slices.Grow(tt.a, tt.n)
			if err != nil {
				t.Error(err)
				return
			}
			if cap(got) != tt.wantCap {
				t.Errorf("got cap %d; want %d", cap(got), tt.wantCap)
			}
		})
	}
	return
}

func TestImplode(t *testing.T) {
	sv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if reflect.DeepEqual(slices.Implode([]int{1, 2, 3}, []int{4, 5, 6}, []int{7}, []int{8, 9}), sv) == false {
		t.Error(`slices not equals`)
	}
}
