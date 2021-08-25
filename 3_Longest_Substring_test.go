package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "test2",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "test3",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "test4",
			args: args{s: " "},
			want: 1,
		},
		{
			name: "test5",
			args: args{s: "dvdf"},
			want: 3,
		},
		{
			name: "test6",
			args: args{s: "abba"},
			want: 2,
		},
		{
			name: "test7",
			args: args{s: "aab"},
			want: 2,
		},
		{
			name: "test8",
			args: args{s: "tmmzuxt"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
