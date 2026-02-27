package game

import (
	"image/color"
	"math/rand"
)

type Tetromino struct {
	Shape [][]bool
	Color color.Color
	X     int
	Y     int
}

var tetrominos = []*Tetromino{
	{ // I
		Shape: [][]bool{
			{true, true, true, true},
		},
	},
	{ // O
		Shape: [][]bool{
			{true, true},
			{true, true},
		},
	},
	{ // T
		Shape: [][]bool{
			{false, true, false},
			{true, true, true},
		},
	},
	{ // S
		Shape: [][]bool{
			{false, true, true},
			{true, true, false},
		},
	},
	{ // Z
		Shape: [][]bool{
			{true, true, false},
			{false, true, true},
		},
	},
	{ // J
		Shape: [][]bool{
			{true, false, false},
			{true, true, true},
		},
	},
	{ // L
		Shape: [][]bool{
			{false, false, true},
			{true, true, true},
		},
	},
}

func TetrominoFactory() *Tetromino {
	t := *tetrominos[rand.Intn(len(tetrominos))]

	t.Color = Colors[rand.Intn(len(Colors))]

	t.X = 0
	t.Y = 0
	return &t
}

func (t *Tetromino) Rotate90() {
	h := len(t.Shape)
	w := len(t.Shape[0])

	newShape := make([][]bool, w)
	for i := range w {
		newShape[i] = make([]bool, h)
	}

	for y := range h {
		for x := range w {
			newShape[x][h-1-y] = t.Shape[y][x]
		}
	}

	t.Shape = newShape
}
