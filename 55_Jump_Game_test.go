package leetcode

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				nums: []int{2,3,1,1,4},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0},
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2},
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{0, 1},
			},
			want: false,
		},
		{
			name: "case 6",
			args: args{
				nums: []int{2, 0},
			},
			want: true,
		},
		{
			name: "case 7",
			args: args{
				nums: []int{1,1,2,2,0,1,1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
