package leetcode

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				s: "aabb",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
