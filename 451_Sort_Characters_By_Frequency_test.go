package leetcode

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "tree",
			},
			want: "eert",
		},
		{
			name: "case 2",
			args: args{
				s: "cccaaa",
			},
			want: "aaaccc",
		},
		{
			name: "case 3",
			args: args{
				s: "Aabb",
			},
			want: "bbAa",
		},
		{
			name: "case 4",
			args: args{
				s: "raaeaedere",
			},
			want: "eeeeaaarrd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
