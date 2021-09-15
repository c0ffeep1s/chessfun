package board

import (
	"errors"
	"math/rand"
	"time"
)

var pieceKeys [13][120]uint64
var sideKey uint64
var castleKeys [16]uint64

// initHashKeys Initialize the hash keys
func initHashKeys() {
	rand.Seed(time.Now().UTC().UnixNano())

	for x := 0; x < 13; x++ {
		for y := 0; y < 120; y++ {
			pieceKeys[x][y] = rand.Uint64()
		}
	}
	sideKey = rand.Uint64()
	for i := 0; i < 16; i++ {
		castleKeys[i] = rand.Uint64()
	}
}

// generatePosKey generates a unique key for the current position
func (pos *PositionStruct) generatePosKey() (uint64, error) {
	var finalKey uint64
	var piece int

	// Pieces
	for sq := 0; sq < SquareNumber; sq++ {
		piece = pos.Pieces[sq]
		if piece != noSquare && piece != empty && piece != offBoard {
			if !(piece >= wP && piece <= bK) {
				return 0, errors.New("Piece value out of bounds")
			}
			finalKey ^= pieceKeys[piece][sq]
		}
	}

	// Side
	if pos.Side == white {
		finalKey ^= sideKey
	}

	// EnPassant
	if pos.EnPassant != noSquare {
		if !(pos.EnPassant >= 0 && pos.EnPassant < SquareNumber) {
			return 0, errors.New("EnPassant value out of bounds")
		}
		finalKey ^= pieceKeys[empty][pos.EnPassant]
	}

	// CastlePerm
	if !(pos.CastlePerm >= 0 && pos.CastlePerm <= 15) {
		return 0, errors.New("CastlePerm value out of bounds")
	}
	finalKey ^= castleKeys[pos.CastlePerm]

	return finalKey, nil
}

// hashPiece update hash with pieces new square
func (pos *PositionStruct) hashPiece(piece, sq int) {
	pos.PosKey ^= (pieceKeys[piece][sq])
}

// hashCastle update hash with castle perms
func (pos *PositionStruct) hashCastle() {
	pos.PosKey ^= (castleKeys[pos.CastlePerm])
}

// hashSide update hash with new side
func (pos *PositionStruct) hashSide() {
	pos.PosKey ^= (sideKey)
}

// hashEnPas update hash for EnPas square
func (pos *PositionStruct) hashEnPas() {
	pos.PosKey ^= (pieceKeys[empty][pos.EnPassant])
}
