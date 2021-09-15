package board

import "fmt"

var bitTables = [64]int{
	63, 30, 3, 32, 25, 41, 22, 33, 15, 50, 42, 13, 11, 53, 19, 34, 61, 29, 2,
	51, 21, 43, 45, 10, 18, 47, 1, 54, 9, 57, 0, 35, 62, 31, 40, 4, 49, 5, 52,
	26, 60, 6, 23, 44, 46, 27, 56, 16, 7, 39, 48, 24, 59, 14, 12, 55, 38, 28,
	58, 20, 37, 17, 36, 8}

var setMask [64]uint64
var clearMask [64]uint64

//Bitboard file maseks
var fileMasks [8]uint64
var rankMasks [8]uint64

//Passed and isolated pawns masks
var blackPassedMasks [64]uint64
var whitePassedMasks [64]uint64
var isolatedMasks [64]uint64

//initBitMasks Initialize the bit masks
func initBitMasks() {
	for i := 0; i < 64; i++ {
		setMask[i] |= uint64(1) << uint64(i)
		clearMask[i] = ^setMask[i]
	}

	for r := rank8; r >= rank1; r-- {
		for f := fileA; f <= fileH; f++ {
			sq := r*8 + f
			fileMasks[f] |= (uint64(1) << sq)
			rankMasks[r] |= (uint64(1) << sq)
		}
	}

	for sq := 0; sq < 64; sq++ {

		tsq := sq + 8
		for tsq < 64 {
			whitePassedMasks[sq] |= (uint64(1) << uint64(tsq))
			tsq += 8
		}

		tsq = sq - 8
		for tsq >= 0 {
			blackPassedMasks[sq] |= (uint64(1) << uint64(tsq))
			tsq -= 8
		}

		if filesBoard[sq64ToSq120[sq]] > fileA {
			isolatedMasks[sq] |= fileMasks[filesBoard[sq64ToSq120[sq]]-1]

			tsq = sq + 7
			for tsq < 64 {
				whitePassedMasks[sq] |= (uint64(1) << uint64(tsq))
				tsq += 8
			}

			tsq = sq - 9
			for tsq >= 0 {
				blackPassedMasks[sq] |= (uint64(1) << uint64(tsq))
				tsq -= 8
			}
		}

		if filesBoard[sq64ToSq120[sq]] < fileH {
			isolatedMasks[sq] |= fileMasks[filesBoard[sq64ToSq120[sq]]+1]

			tsq = sq + 9
			for tsq < 64 {
				whitePassedMasks[sq] |= (uint64(1) << uint64(tsq))
				tsq += 8
			}

			tsq = sq - 7
			for tsq >= 0 {
				blackPassedMasks[sq] |= (uint64(1) << uint64(tsq))
				tsq -= 8
			}
		}
	}
}

// printBitBoard Will print a visual representation of a bitboard to screen
//nolint
func printBitBoard(bitboard uint64) {
	fmt.Print("\n")

	for rank := rank8; rank >= rank1; rank-- {
		for file := fileA; file <= fileH; file++ {
			sq := fileRankToSquare(file, rank)
			sq64 := sq120ToSq64[sq]

			if ((uint64(1) << sq64) & bitboard) != 0 {
				fmt.Print(" X ")
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}

//popBit Pop first 1 bit off and return its index
func popBit(bitboard *uint64) int {
	var board uint64 = *bitboard ^ (*bitboard - 1)
	var fold uint32 = uint32((board & 0xffffffff) ^ (board >> 32))
	*bitboard &= (*bitboard - 1)
	return bitTables[(fold*0x783a9b23)>>26]
}

//countBits Count the number of 1 bits in a bitboard
func countBits(board uint64) int {
	var r int
	for r = 0; board > 0; r++ {
		board &= board - 1
	}
	return r
}

//clearBit removes given square from bitboard
func clearBit(bitboard *uint64, square int) {
	*bitboard &= clearMask[square]
}

//setBit sets bit to given square
func setBit(bitboard *uint64, square int) {
	*bitboard |= setMask[square]
}
