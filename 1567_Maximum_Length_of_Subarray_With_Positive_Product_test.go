package leetcode

import "testing"

func Test_getMaxLen(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{
					1, -2, -3, 4,
				},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{
					0, 1, -2, -3, -4,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxLen(tt.args.nums); got != tt.want {
				t.Errorf("getMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
