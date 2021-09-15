package board

import (
	"unicode/utf8"
)

// castlePermToChar Convert castle perms to chars
func (pos *PositionStruct) castlePermToChar() (rune, rune, rune, rune) {
	var wKt rune = '-'
	var wQt rune = '-'
	var bKt rune = '-'
	var bQt rune = '-'

	if pos.CastlePerm&wkcastle != 0 {
		wKt = 'K'
	}

	if pos.CastlePerm&wqcastle != 0 {
		wQt = 'Q'
	}

	if pos.CastlePerm&bkcastle != 0 {
		bKt = 'k'
	}

	if pos.CastlePerm&bqcastle != 0 {
		bQt = 'q'
	}

	return wKt, wQt, bKt, bQt
}

// isPieceBig returns true if piece is big
func isPieceBig(piece int) bool {
	return piece != empty && piece != wP && piece != bP
}

// isPieceBig returns true if piece is major
func isPieceMajor(piece int) bool {
	return piece != empty && piece != wP && piece != wN && piece != wB &&
		piece != bP && piece != bN && piece != bB

}

// isPieceBig returns true if piece is minor
func isPieceMinor(piece int) bool {
	return piece == wN || piece == wB || piece == bN || piece == bB

}

// isPieceSlider returns true if piece is slider
//nolint
func isPieceSlider(piece int) bool {
	return piece == wB || piece == wR || piece == wQ ||
		piece == bB || piece == bR || piece == bQ
}

// GetPieceColor returns the color of a piece
func getPieceColor(piece int) int {
	if piece >= wP && piece <= wK {
		return white
	}

	if piece >= bP && piece <= bK {
		return black
	}
	return both
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
