package board

import "fmt"

// TakeMove Take back the last move
func (pos *PositionStruct) TakeMove() error {

	pos.HisPly--
	pos.Ply--

	move := pos.History[pos.HisPly].Move
	from := GetFrom(move)
	to := GetTo(move)

	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}
		if !squareOnBoard(from) || !squareOnBoard(to) {
			return fmt.Errorf("From: %d or To: %d square for move is off the board", from, to)
		}
	}

	if pos.EnPassant != noSquare {
		pos.hashEnPas()
	}
	pos.hashCastle()

	pos.CastlePerm = pos.History[pos.HisPly].CastlePerm
	pos.FiftyMove = pos.History[pos.HisPly].FiftyMove
	pos.EnPassant = pos.History[pos.HisPly].EnPassant

	if pos.EnPassant != noSquare {
		pos.hashEnPas()
	}
	pos.hashCastle()

	// Flipping side to move
	pos.Side ^= 1
	pos.hashSide()

	if MoveFlagEP&move != 0 {
		if pos.Side == white {
			err := pos.addPiece(to-10, bP)
			if err != nil {
				return err
			}
		} else {
			err := pos.addPiece(to+10, wP)
			if err != nil {
				return err
			}
		}
	} else if MoveFlagCA&move != 0 {
		var err error
		switch to {
		case c1:
			err = pos.movePiece(d1, a1)
		case c8:
			err = pos.movePiece(d8, a8)
		case g1:
			err = pos.movePiece(f1, h1)
		case g8:
			err = pos.movePiece(f8, h8)
		default:
			return fmt.Errorf("Invalid castle move to %d", to)
		}
		if err != nil {
			return err
		}
	}

	err := pos.movePiece(to, from)
	if err != nil {
		return err
	}

	if pos.Pieces[from] == wK || pos.Pieces[from] == bK {
		pos.KingSquare[pos.Side] = from
	}

	captured := GetCapture(move)
	if captured != empty {
		if DEBUG && !pieceValid(captured) {
			return fmt.Errorf("Invalid capture piece %d", captured)
		}
		err := pos.addPiece(to, captured)
		if err != nil {
			return err
		}
	}

	promotedPiece := GetPromoted(move)
	if promotedPiece != empty {
		if !pieceValid(promotedPiece) || !isPieceBig(promotedPiece) {
			return fmt.Errorf("Invalid promotion piece of %d", promotedPiece)
		}
		err = pos.clearPiece(from)
		if err != nil {
			return err
		}
		if getPieceColor(promotedPiece) == white {
			err = pos.addPiece(from, wP)
		} else {
			err = pos.addPiece(from, bP)
		}
		if err != nil {
			return err
		}
	}

	if DEBUG {
		err = pos.CheckBoard()
		if err != nil {
			return err
		}
	}

	return nil
}
