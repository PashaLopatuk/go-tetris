package game

import (
	"image/color"
)

type Cell struct {
	IsFilled  bool
	Color     *color.Color
	Tetromino *Tetromino
}

func (c *Cell) Clear() {
	c.IsFilled = false
	c.Color = nil
	c.Tetromino = nil
}

func (c *Cell) Fill(piece *Tetromino) {
	c.IsFilled = true
	c.Color = &piece.Color
	c.Tetromino = piece
}

func CreateCell() *Cell {
	return &Cell{
		IsFilled: false,
	}
}

var PastelPink = color.RGBA{R: 255, G: 182, B: 193, A: 255}
var PastelRose = color.RGBA{R: 250, G: 200, B: 200, A: 255}
var PastelPeach = color.RGBA{R: 255, G: 218, B: 185, A: 255}
var PastelApricot = color.RGBA{R: 255, G: 225, B: 205, A: 255}

var PastelYellow = color.RGBA{R: 255, G: 245, B: 200, A: 255}
var PastelMint = color.RGBA{R: 200, G: 255, B: 220, A: 255}
var PastelGreen = color.RGBA{R: 190, G: 240, B: 200, A: 255}

var PastelCyan = color.RGBA{R: 200, G: 240, B: 255, A: 255}
var PastelBlue = color.RGBA{R: 190, G: 210, B: 255, A: 255}
var PastelLavender = color.RGBA{R: 220, G: 200, B: 255, A: 255}

var PastelLilac = color.RGBA{R: 235, G: 210, B: 255, A: 255}
var PastelPurple = color.RGBA{R: 210, G: 190, B: 240, A: 255}
var PastelPlum = color.RGBA{R: 225, G: 200, B: 230, A: 255}

var PastelBeige = color.RGBA{R: 245, G: 235, B: 220, A: 255}
var PastelSand = color.RGBA{R: 240, G: 230, B: 210, A: 255}
var PastelGray = color.RGBA{R: 220, G: 220, B: 220, A: 255}

var Colors = []color.Color{
	PastelPink,
	PastelRose,
	PastelPeach,
	PastelApricot,
	PastelYellow,
	PastelMint,
	PastelGreen,
	PastelCyan,
	PastelBlue,
	PastelLavender,
	PastelLilac,
	PastelPurple,
	PastelPlum,
	PastelBeige,
	PastelSand,
	PastelGray,
}
