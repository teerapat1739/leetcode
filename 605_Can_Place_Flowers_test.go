package leetcode

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "Example3",
			args: args{
				flowerbed: []int{1, 0, 1, 0, 1, 0, 1},
				n:         1,
			},
			want: false,
		},
		{
			name: "Example4",
			args: args{
				flowerbed: []int{0, 0, 1, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Example5",
			args: args{
				flowerbed: []int{0, 0, 1, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Example6",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1, 0, 0},
				n:         2,
			},
			want: true,
		},
		{
			name: "Example7",
			args: args{
				flowerbed: []int{0},
				n:         1,
			},
			want: true,
		},
		{
			name: "Example8",
			args: args{
				flowerbed: []int{0, 0, 1, 0, 0},
				n:         2,
			},
			want: true,
		},
		{
			name: "Example9",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "Example10",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 0, 0, 0, 0, 1},
				n:         3,
			},
			want: true,
		},
		{
			name: "Example11",
			args: args{
				flowerbed: []int{0, 0, 1, 0, 0},
				n:         1,
			},
			want: true,
		},
		{
			name: "Example12",
			args: args{
				flowerbed: []int{0, 0, 0, 0, 1},
				n:         2,
			},
			want: true,
		},
		{
			name: "Example13",
			args: args{
				flowerbed: []int{0},
				n:         0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
