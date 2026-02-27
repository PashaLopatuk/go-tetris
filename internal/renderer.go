package game

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Renderer struct {
	app          fyne.App
	window       fyne.Window
	container    *fyne.Container
	windowWidth  int
	windowHeight int
	blockSize    int
}

func (r *Renderer) Draw(board Board, currentPiece Tetromino) {
	r.container.RemoveAll()

	fyne.Do(func() {
		for y := range len(board.grid) {
			for x := range len(board.grid[y]) {
				cell := board.grid[y][x]
				if cell == nil {
					continue
				}

				r.DrawCell(cell, x, y)
			}
		}

		r.container.Refresh()
	})
}

func (r *Renderer) DrawCell(cell *Cell, x int, y int) {
	rect := CreateRectangle(CreateRectParams{
		X:      new(x * r.blockSize),
		Y:      new(y * r.blockSize),
		Width:  &r.blockSize,
		Height: &r.blockSize,
		Color:  cell.Color,
	})
	r.container.Add(rect)
}

func (r *Renderer) DrawGameOver() {

}

func CreateRenderer() *Renderer {
	a := app.New()
	w := a.NewWindow("Tetris")

	width := 400
	height := 800

	blockSize := width / 10

	w.Resize(fyne.NewSize(float32(width), float32(height)))

	c := container.NewWithoutLayout()

	w.SetContent(c)

	return &Renderer{
		app: a, window: w, container: c,
		windowWidth: width, windowHeight: height,
		blockSize: blockSize,
	}
}

type CreateRectParams struct {
	X      *int
	Y      *int
	Width  *int
	Height *int
	Color  *color.Color
}

func CreateRectangle(params CreateRectParams) *canvas.Rectangle {
	var x int
	var y int
	width := 50
	height := 50
	var rectColor color.Color = color.Black

	if params.X != nil {
		x = *params.X
	}
	if params.Y != nil {
		y = *params.Y
	}
	if params.Width != nil {
		width = *params.Width
	}
	if params.Height != nil {
		height = *params.Height
	}
	if params.Color != nil {
		rectColor = *params.Color
	}

	rect := canvas.NewRectangle(rectColor)
	rect.Resize(fyne.NewSize(float32(width), float32(height)))
	rect.Move(fyne.NewPos(float32(x), float32(y)))
	return rect
}
