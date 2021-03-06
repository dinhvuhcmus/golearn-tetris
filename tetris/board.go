package tetris

import (
	"github.com/nsf/termbox-go"
)

type ColorMap map[Vector]termbox.Attribute

func (cm ColorMap) contains(v Vector) bool {
	_, ok := cm[v]
	return ok
}

type Board struct {
	cells           ColorMap
	currentPiece    *Piece
	currentPosition Vector
}

func newBoard() *Board {
	board := new(Board)
	board.cells = make(ColorMap)
	return board
}

func (board *Board) currentPieceInCollision() bool {
	for _, point := range board.currentPiece.instance() {
		attemptedPoint := point.plus(board.currentPosition)
		if attemptedPoint.x < 0 || attemptedPoint.x >= width ||
			attemptedPoint.y < 0 || attemptedPoint.y >= height ||
			board.cells.contains(attemptedPoint) {
			return true
		}
	}
	return false
}

func (board *Board) moveIfPossible(translation Vector) bool {
	position := board.currentPosition
	board.currentPosition = board.currentPosition.plus(translation)
	if board.currentPieceInCollision() {
		board.currentPosition = position
		return false
	}
	return true
}

// Merge the blocks of the current piece into the game board and remove the current piece.
func (board *Board) mergeCurrentPiece() {
	for _, point := range board.currentPiece.instance() {
		board.cells[point.plus(board.currentPosition)] = board.currentPiece.color
	}
	board.currentPiece = nil
}

func (board *Board) clearRows() {
	rowsCleared := 0
	y := height - 1
	for y >= 0 {
		for board.rowComplete(y) {
			rowsCleared += 1
			board.collapseRow(y)
		}
		y -= 1
	}
}

func (board *Board) rowComplete(y int) bool {
	for x := 0; x < width; x++ {
		if !board.cells.contains(Vector{x, y}) {
			return false
		}
	}
	return true
}

func (board *Board) collapseRow(rowY int) {
	for y := rowY - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if color, ok := board.cells[Vector{x, y}]; ok {
				board.cells[Vector{x, y + 1}] = color
			} else {
				delete(board.cells, Vector{x, y + 1})
			}
		}
	}
	// Clear the top row completely
	for x := 0; x < width; x++ {
		delete(board.cells, Vector{x, 0})
	}
}

func (board *Board) clearedRows() []int {
	cleared := make([]int, 0)
	for y := 0; y < height; y++ {
		if board.rowComplete(y) {
			cleared = append(cleared, y)
		}
	}
	return cleared
}

func (board *Board) CellColor(position Vector) termbox.Attribute {
	if color, ok := board.cells[position]; ok {
		return color
	}
	if board.currentPiece == nil {
		return backgroundColor
	}
	for _, point := range board.currentPiece.instance() {
		if point.plus(board.currentPosition).equals(position) {
			return board.currentPiece.color
		}
	}
	return backgroundColor
}
