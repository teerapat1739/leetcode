package leetcode

import "testing"

func Test_maxProfitII(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitII(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitII() = %v, want %v", got, tt.want)
			}
		})
	}
}
