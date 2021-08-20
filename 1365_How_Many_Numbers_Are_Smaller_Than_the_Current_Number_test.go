package leetcode

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test1",
		// 	args: args{
		// 		nums: []int{5, 2, 6, 1},
		// 	},
		// 	want: []int{3, 2, 1, 0},
		// },
		{
			name: "test2",
			args: args{
				nums: []int{6, 5, 4, 8},
			},
			want: []int{2, 1, 0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
