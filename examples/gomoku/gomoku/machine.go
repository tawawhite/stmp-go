// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-08 23:34:40
package gomoku

import (
	"fmt"
	"github.com/acrazing/stmp-go/examples/gomoku/gomoku/num"
)

type Hand struct {
	Row uint32
	Col uint32
	// the time offset from createdAt, 0.1s as 1
	Ts uint32
}

type History []Hand

func (h History) String() string {
	var buf []byte
	for _, h := range h {
		buf = append(buf, num.Digits[h.Row], num.Digits[h.Col])
		buf = append(buf, num.Uint32tos(h.Ts)...)
		buf = append(buf, '.')
	}
	if len(buf) > 0 {
		return string(buf)
	}
	return string(buf)
}

const (
	PieceNone  = 0
	PieceBlack = 1
	PieceWhite = 2
)

type Code int

const (
	CodeOk = iota
	CodeInvalidLocation
	CodeForbiddenLocation
	CodeUnavailableLocation
	CodeInvalidPiece
)

type Machine struct {
	Rows           uint32
	Cols           uint32
	History        History
	Board          [][]uint32
	Winner         uint32
	WinnerLocation *Hand
	CreatedAt      uint64
}

func NewMachine(rows uint32, cols uint32, history History) *Machine {
	m := new(Machine)
	if rows < 5 || cols < 5 || rows > 32 || cols > 32 {
		panic(fmt.Sprintf("invalid board size %d x %d.", rows, cols))
	}
	m.History = history
	m.Board = make([][]uint32, rows, rows)
	for i := uint32(0); i < rows; i++ {
		m.Board[i] = make([]uint32, cols, cols)
	}
	for i, h := range history {
		m.Board[h.Row][h.Col] = uint32(i%2 + 1)
	}
	return m
}

func (m *Machine) IsForbidden(row, col uint32) bool {
	return false
}

func (m *Machine) Play(piece uint32, h *Hand) Code {
	if piece != uint32(len(m.History))%2+1 {
		return CodeInvalidPiece
	}
	if h.Col > m.Cols || h.Row > m.Rows {
		return CodeInvalidLocation
	}
	if m.Board[h.Row][h.Col] != PieceNone {
		return CodeUnavailableLocation
	}
	if m.IsForbidden(h.Row, h.Col) {
		return CodeForbiddenLocation
	}
	m.Board[h.Row][h.Col] = piece
	m.Compute()
	return CodeOk
}

func (m *Machine) Compute() {
}
