package leetcode

import (
	"testing"
)

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "case 1",
		// 	args: args{
		// 		board: [][]byte{
		// 			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		// 			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		// 			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		// 			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		// 			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		// 			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		// 			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		// 			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		// 			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		// 		},
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "case 2",
		// 	args: args{
		// 		board: [][]byte{
		// 			{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		// 			{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		// 			{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		// 			{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		// 			{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		// 			{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		// 			{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		// 			{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		// 			{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
		// 		},
		// 	},
		// 	want: false,
		// },
		// {
		// 	name: "case 3",
		// 	args: args{
		// 		board: [][]byte{
		// 			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		// 			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		// 			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		// 			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		// 			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		// 			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		// 			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		// 			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		// 			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		// 		},
		// 	},
		// 	want: true,
		// },
		{
			name: "case 4",
			args: args{
				board: [][]byte{
					{'.', '8', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '2', '.', '.', '.', '.'},
					{'.', '6', '.', '.', '.', '.', '1', '.', '4'},
					{'.', '.', '.', '9', '.', '.', '7', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '4', '.'},
					{'.', '.', '1', '.', '.', '8', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '5', '.', '7', '.'},
					{'.', '.', '3', '.', '.', '5', '6', '.', '.'},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkBox(t *testing.T) {
	type args struct {
		board [9][9]int
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				board: [9][9]int{
					{0, 0, 0, 0, 5, 0, 0, 1, 0},
					{0, 4, 0, 3, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 3, 0, 0, 1},
					{8, 0, 0, 0, 0, 0, 0, 2, 0},
					{0, 0, 2, 0, 7, 0, 0, 0, 0},
					{0, 1, 5, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 2, 0, 0, 0},
					{0, 2, 0, 9, 0, 0, 0, 0, 0},
					{0, 0, 4, 0, 0, 0, 0, 0, 0},
				},
				i: 0,
				j: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkBox(tt.args.board, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("checkBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkCol(t *testing.T) {
	type args struct {
		board [9][9]int
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				board: [9][9]int{
					{0, 8, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0},
					{0, 6, 0, 0, 0, 0, 1, 0, 4},
					{0, 0, 0, 9, 0, 0, 7, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 4, 0},
					{0, 0, 1, 0, 0, 8, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 5, 0, 7, 0},
					{0, 0, 3, 0, 0, 5, 6, 0, 0},
				},
				i: 7,
				j: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkCol(tt.args.board, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("checkCol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkRow(t *testing.T) {
	type args struct {
		board [9][9]int
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				board: [9][9]int{
					{0, 8, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0},
					{0, 6, 0, 0, 0, 0, 1, 0, 4},
					{0, 0, 0, 9, 0, 0, 7, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 4, 0},
					{0, 0, 1, 0, 0, 8, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 5, 0, 7, 0},
					{0, 0, 3, 0, 0, 5, 6, 0, 0},
				},
				i: 7,
				j: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRow(tt.args.board, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("checkRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
