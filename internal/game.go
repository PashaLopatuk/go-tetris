package game

import (
	"context"
	"time"

	"fyne.io/fyne/v2"
)

type Game struct {
	board        *Board
	input        *InputController
	renderer     *Renderer
	state        *State
	currentPiece *Tetromino
}

func Create() *Game {
	renderer := CreateRenderer()

	controller := NewInputController(func(action *Action) {
		CurrentActionChan <- action
	})

	renderer.window.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		controller.HandleKey(ev)
	})

	state := CreateState()
	board := CreateBoard(10, 20)

	return &Game{
		renderer: renderer,
		state:    state,
		board:    board,
	}
}

var CurrentActionChan = make(chan *Action)

func (g *Game) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	g.currentPiece = TetrominoFactory()

	go func(ctx context.Context) {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				g.UpdatePhysics()
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case a, ok := <-CurrentActionChan:
				if !ok {
					return
				}
				g.UpdateInput(a)
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	g.renderer.window.ShowAndRun()
}

func (g *Game) Pause() {

}

func (g *Game) UpdateInput(a *Action) {
	if !g.state.IsRunning {
		return
	}
	if a != nil {
		g.HandleInput(*a)
	}
	g.renderer.Draw(*g.board, *g.currentPiece)
}

func (g *Game) UpdatePhysics() {
	if !g.state.IsRunning {
		return
	}
	if g.board.IsValidPosition(g.currentPiece, 0, 1) {
		g.MoveCurrentPiece(0, 1)
	} else {
		g.currentPiece = TetrominoFactory()
	}
	g.HandleFilledFields(*g.board)

	g.renderer.Draw(*g.board, *g.currentPiece)
}

func (g *Game) Reset() {

}

func (g *Game) HandleFilledFields(board Board) {
	for fieldRow := range board.grid {
		isFilled := true
		for px := range board.AtRow(fieldRow) {
			isFilled = isFilled && board.At(px, fieldRow).IsFilled
		}
		if isFilled {
			for py := fieldRow; py > 0; py-- {
				for px := range board.AtRow(fieldRow) {
					board.At(px, py).Clear()
					if board.At(px, py-1).IsFilled {
						board.At(px, py).Fill(
							board.At(px, py-1).Tetromino,
						)
					}
				}
			}
		}
	}
}

func (g *Game) HandleInput(a Action) {
	switch a.ActionType {
	case MOVE_LEFT:
		g.MoveCurrentPiece(-1, 0)
	case MOVE_RIGHT:
		g.MoveCurrentPiece(1, 0)
	case MOVE_DOWN:
		g.MoveCurrentPiece(0, 1)
	case ROTATE:
		g.HandleRotate()
	}
}

func (g *Game) HandleRotate() {
	for py, _ := range g.currentPiece.Shape {
		for px, _ := range g.currentPiece.Shape[py] {
			g.board.At(g.currentPiece.X+px, g.currentPiece.Y+py).Clear()
		}
	}
	g.currentPiece.Rotate90()
}

func (g *Game) MoveCurrentPiece(dx, dy int) {
	if !g.board.IsValidPosition(g.currentPiece, dx, dy) {
		return
	}
	for py, _ := range g.currentPiece.Shape {
		for px, _ := range g.currentPiece.Shape[py] {
			if !g.currentPiece.Shape[py][px] {
				continue
			}

			oldX := g.currentPiece.X + px
			oldY := g.currentPiece.Y + py

			if oldY >= 0 && oldY < len(g.board.grid) &&
				oldX >= 0 && oldX < len(g.board.grid[0]) {
				g.board.At(oldX, oldY).Clear()
			}
		}
	}

	g.currentPiece.X += dx
	g.currentPiece.Y += dy

	for py, _ := range g.currentPiece.Shape {
		for px, _ := range g.currentPiece.Shape[py] {
			if !g.currentPiece.Shape[py][px] {
				continue
			}

			boardX := g.currentPiece.X + px
			boardY := g.currentPiece.Y + py

			if boardY >= 0 && boardY < len(g.board.grid) &&
				boardX >= 0 && boardX < len(g.board.AtRow(boardY)) {
				g.board.At(boardX, boardY).Fill(g.currentPiece)
			}
		}
	}

}
