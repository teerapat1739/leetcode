package leetcode

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				n: 5,
			},
			want: 8,
		},
		{
			name: "case 4",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "case 5",
			args: args{
				n: 6,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs3(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
