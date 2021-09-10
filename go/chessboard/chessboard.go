package chessboard

// Rank stores if a square is occupied by a piece.
type Rank []bool

// Chessboard contains eight Ranks, accessed with values from 'A' to 'H'
type Chessboard map[byte]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func (cb Chessboard) CountInRank(rank byte) (ret int) {
	ret = 0
	for _, sqr := range cb[rank] {
		if sqr {
			ret++
		}
	}
	return ret
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func (cb Chessboard) CountInFile(file int) (ret int) {
	if file > 8 {
		return 0
	}
	ret = 0
	for _, row := range cb {
		if row[file-1] {
			ret++
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard
func (cb Chessboard) CountAll() (ret int) {
	ret = 0
	for _, row := range cb {
		for range row {
			ret++
		}

	}
	return ret
}

// CountOccupied returns how many squares are occupied in the chessboard
func (cb Chessboard) CountOccupied() (ret int) {
	ret = 0
	for _, row := range cb {
		for _, sqr := range row {
			if sqr {
				ret++
			}
		}

	}
	return ret
}
