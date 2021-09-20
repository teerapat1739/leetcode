package leetcode

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				nums: []int{
					2, 3, 1, 1, 4,
				},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{
					2, 3, 0, 1, 4,
				},
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{
					1, 1, 2, 2, 0, 1, 1,
				},
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{
					1,
				},
			},
			want: 0,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{
					1, 2, 1, 1, 1,
				},
			},
			want: 3,
		},
		{
			name: "case 6",
			args: args{
				nums: []int{
					2, 0, 2, 0, 1,
				},
			},
			want: 2,
		},
		{
			name: "case 7",
			args: args{
				nums: []int{
					1, 2,
				},
			},
			want: 1,
		},
		{
			name: "case 8",
			args: args{
				nums: []int{
					10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
