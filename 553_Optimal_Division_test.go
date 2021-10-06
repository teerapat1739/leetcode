package leetcode

import "testing"

func Test_optimalDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{
					100,
				},
			},
			want: "100",
		},
		{
			name: "case 2",
			args: args{
				nums: []int{
					100, 2,
				},
			},
			want: "100/2",
		},
		{
			name: "case 3",
			args: args{
				nums: []int{
					2, 3, 4,
				},
			},
			want: "2/(3/4)",
		},
		{
			name: "case 4",
			args: args{
				nums: []int{
					100, 2, 4, 5, 5, 23, 5, 45, 2, 3,
				},
			},
			want: "100/(2/4/5/5/23/5/45/2/3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}
