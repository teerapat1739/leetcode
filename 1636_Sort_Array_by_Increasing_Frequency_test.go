package leetcode

import (
	"reflect"
	"testing"
)

func Test_frequencySort1636(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{
					1, 1, 2, 2, 2, 3,
				},
			},
			want: []int{
				3, 1, 1, 2, 2, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort1636(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("frequencySort1636() = %v, want %v", got, tt.want)
			}
		})
	}
}
