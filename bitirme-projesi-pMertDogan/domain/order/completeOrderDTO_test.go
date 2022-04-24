package order

import (
	"testing"

)

func TestUniqueIntSlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"shouldOK empty slice", args{[]int{}}, true},
		{"shouldOK one element slice", args{[]int{1}}, true},
		{"shouldOK two element slice", args{[]int{1, 2}}, true},
		{"shouldOK three element slice", args{[]int{1, 2, 3}}, true},
		{"shouldFalse three element slice with duplicates", args{[]int{1, 2, 2}}, false},
		{"shouldFalse three element slice with duplicates", args{[]int{1, 2, 3, 3}}, false},
		{"shouldFalse three element slice with duplicates", args{[]int{1, 2, 3, 3, 3}}, false},
		{"shouldFalse three element slice with duplicates", args{[]int{1, 2, 3, 3, 3, 3}}, false},
		{"shouldFalse three element slice with duplicates", args{[]int{1, 2, 3, 3, 3, 3, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueIntSlice(tt.args.slice); got != tt.want {
				t.Errorf("UniqueIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
