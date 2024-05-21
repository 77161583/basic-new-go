package homework

import (
	"reflect"
	"testing"
)

func TestShrinkSlice(t *testing.T) {
	cases := []struct {
		name   string
		slice  []int
		newCap int
		want   []int
	}{
		{
			name:   "shrink to smaller capacity",
			slice:  []int{1, 2, 3, 4, 5},
			newCap: 3,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "no need to shrink",
			slice:  []int{1, 2, 3},
			newCap: 5,
			want:   []int{1, 2, 3},
		},
		{
			name:   "shrink to zero capacity",
			slice:  []int{1},
			newCap: 0,
			want:   []int{1},
		},
		{
			name:   "shrink to same capacity",
			slice:  []int{1, 2, 3, 4, 5},
			newCap: 5,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "empty slice",
			slice:  []int{},
			newCap: 3,
			want:   []int{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ShrinkSlice(c.slice, c.newCap)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("ShrinkSlice(%v, %d) = %v, want %v", c.slice, c.newCap, got, c.want)
			}
			if cap(got) != c.newCap && c.newCap > 0 {
				t.Errorf("Expected new capacity to be %d, got %d", c.newCap, cap(got))
			}
		})
	}
}
