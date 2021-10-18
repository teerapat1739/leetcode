package leetcode

import "testing"

func Test_maxProfitWithFree(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				prices: []int{1, 3, 7, 5, 10, 3},
				fee:    3,
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				prices: []int{1, 3, 2, 8, 12, 20},
				fee:    2,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitWithFreeIII(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfitWithFree() = %v, want %v", got, tt.want)
			}
		})
	}
}
