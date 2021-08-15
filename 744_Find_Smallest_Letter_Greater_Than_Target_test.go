package leetcode

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "test1",
			args: args{
				letters: []byte("cfj"),
				target:  'a',
			},
			want: 'c',
		},
		{
			name: "test2",
			args: args{
				letters: []byte("cfj"),
				target:  'j',
			},
			want: 'c',
		},
		{
			name: "test3",
			args: args{
				letters: []byte("cfj"),
				target:  'c',
			},
			want: 'f',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
