package kakao2018

import (
	"reflect"
	"testing"
)

func TestHappy9(t *testing.T) {
	cases := []struct {
		m, n  int
		board []string
		want  int
	}{
		{4, 5, []string{"CCBDE", "AAADE", "AAABF", "CCBBF"}, 14},
		{6, 6, []string{"TTTANT", "RRFACC", "RRRFCC", "TRRRAA", "TTMMMF", "TMMTTJ"}, 15},
	}

	for _, c := range cases {
		rz := calcScore(c.m, c.n, c.board)
		if rz != c.want {
			t.Fatalf("invalid result. rz:%d want:%d", rz, c.want)
		}
	}
}

func calcScore(m int, n int, board []string) int {
	score1, board := CalcScoreAndGetNewBoard(m, n, board)
	nextBoard := ShirinkBoard(board)
	score2, board := CalcScoreAndGetNewBoard(m, n, nextBoard)
	return score1 + score2
}

type ScoreAndBoard struct {
	score int
	board []string
}

func TestCalcScoreAndGetNewBoard(t *testing.T) {
	cases := []struct {
		m, n  int
		board []string
		want  ScoreAndBoard
	}{
		{4, 5, []string{"CCBDE", "AAADE", "AAABF", "CCBBF"},
			ScoreAndBoard{6, []string{"CCBDE", "   DE", "   BF", "CCBBF"}}},
	}

	for _, c := range cases {
		score, board := CalcScoreAndGetNewBoard(c.m, c.n, c.board)
		if score != c.want.score {
			t.Fatalf("invalid score. rz:%d want:%d", score, c.want.score)
		}
		if !reflect.DeepEqual(board, c.want.board) {
			t.Fatalf("invalid board. rz:%v want:%v", board, c.want.board)
		}
	}
}

func CalcScoreAndGetNewBoard(m int, n int, board []string) (int, []string) {
	tmpBoard := make([][]bool, m)
	for i, _ := range tmpBoard {
		tmpBoard[i] = make([]bool, n)
	}

	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if board[i][j] == ' ' {
				continue
			}

			if board[i][j] == board[i+1][j] &&
				board[i+1][j] == board[i][j+1] &&
				board[i][j+1] == board[i+1][j+1] {
				tmpBoard[i][j] = true
				tmpBoard[i+1][j] = true
				tmpBoard[i][j+1] = true
				tmpBoard[i+1][j+1] = true
			}
		}
	}

	score := 0
	result := make([]string, len(board))
	for i, row := range tmpBoard {
		for j, isDeleted := range row {
			if isDeleted {
				score += 1
				result[i] += " "
			} else {
				result[i] += string(board[i][j])
			}
		}
	}

	return score, result
}

func TestShrinkBoard(t *testing.T) {
	cases := []struct {
		board []string
		want  []string
	}{
		{[]string{"CCBDE", "   DE", "   BF", "CCBBF"}, []string{"   DE", "   DE", "CCBBF", "CCBBF"}},
	}

	for _, c := range cases {
		rz := ShirinkBoard(c.board)
		if !reflect.DeepEqual(rz, c.want) {
			t.Fatalf("invalid board. rz:%v want:%v", rz, c.board)
		}
	}
}

func ShirinkBoard(board []string) []string {
	for i := len(board) - 1; i >= 0; i-- {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == ' ' {
				for k := i - 1; k >= 0; k-- {
					if board[k][j] != ' ' {
						out1 := []rune(board[i])
						out1[j] = rune(board[k][j])

						out2 := []rune(board[k])
						out2[j] = rune(board[i][j])

						board[i] = string(out1)
						board[k] = string(out2)
					}
				}
			}
		}
	}
	return board
}

