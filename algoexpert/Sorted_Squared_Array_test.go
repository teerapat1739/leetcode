package algoexpert

import (
	"reflect"
	"testing"
)

func TestSortedSquaredArray(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				array: []int{
					1, 2,
				},
			},
			want: []int{1, 4},
		},
		{
			name: "case 2",
			args: args{
				array: []int{
					-3, 1, 2, 5,
				},
			},
			want: []int{1, 4, 9, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedSquaredArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedSquaredArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
