package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{4, 2, 3, 1, 5},
			},
			want: []int{5, 3, 5, 5, -1},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, -1, 2},
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2, 3, 4, 3},
			},
			want: []int{2, 3, 4, -1, 4},
		},
		{
			name: "test4",
			args: args{
				nums: []int{1, 2, 3, 2, 1},
			},
			want: []int{2, 3, -1, 3, 2},
		},
		{
			name: "test5",
			args: args{
				nums: []int{1, 8, -1, -100, -1, 222},
			},
			want: []int{8, 222, 222, -1, 222, -1},
		},
		{
			name: "test6",
			args: args{
				nums: []int{1, 8, -1, -100, -1, 222, 1111111, -111111},
			},
			want: []int{8, 222, 222, -1, 222, 1111111, -1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
