package game

type Board struct {
	grid [][]*Cell
}

func CreateBoard(w, h int) *Board {
	b := &Board{
		grid: make([][]*Cell, h),
	}

	for y := range b.grid {
		b.grid[y] = make([]*Cell, w)
		for x := range b.grid[y] {
			b.grid[y][x] = CreateCell()
		}
	}

	return b
}

func (b *Board) At(x, y int) *Cell {
	return b.grid[y][x]
}

func (b *Board) AtRow(y int) []*Cell {
	return b.grid[y]
}

func (b *Board) IsValidPosition(piece *Tetromino, dx, dy int) bool {
	x := piece.X
	y := piece.Y

	for py := range len(piece.Shape) {
		for px := range len(piece.Shape[py]) {
			boardX := x + px + dx
			boardY := y + py + dy

			// fmt.Printf("boardX %d\n", boardX)
			// fmt.Printf("boardY %d\n", boardY)

			if boardX < 0 || boardX >= len(b.grid[0]) {
				return false
			}
			if boardY < 0 || boardY >= len(b.grid) {
				return false
			}

			if !piece.Shape[py][px] || b.grid[boardY][boardX].Tetromino == piece {
				continue
			}

			if b.grid[boardY][boardX].IsFilled {
				return false
			}
		}
	}
	return true
}

func (b *Board) LockPiece() {

}

func (b *Board) ClearFull() {

}
