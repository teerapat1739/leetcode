package leetcode

import "testing"

func Test_minSteps(t *testing.T) {
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
				n: 3,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				n: 4,
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				n: 6,
			},
			want: 5,
		},
		{
			name: "case 5",
			args: args{
				n: 8,
			},
			want: 6,
		},
		{
			name: "case 6",
			args: args{
				n: 10,
			},
			want: 7,
		},
		{
			name: "case 6",
			args: args{
				n: 9,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
