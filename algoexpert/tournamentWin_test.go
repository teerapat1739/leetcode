package algoexpert

import "testing"

func TestTournamentWinner(t *testing.T) {
	type args struct {
		competitions [][]string
		results      []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				competitions: [][]string{
					{"HTML", "C#"},
					{"C#", "Python"},
					{"Python", "HTML"},
				},
				results: []int{
					0, 0, 1,
				},
			},
			want: "Python",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TournamentWinner(tt.args.competitions, tt.args.results); got != tt.want {
				t.Errorf("TournamentWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
