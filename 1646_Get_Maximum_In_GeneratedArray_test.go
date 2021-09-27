package leetcode

import "testing"

func Test_getMaximumGenerated(t *testing.T) {
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
				n: 7,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				n: 6,
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				n: 15,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGenerated(tt.args.n); got != tt.want {
				t.Errorf("getMaximumGenerated() = %v, want %v", got, tt.want)
			}
		})
	}
}
