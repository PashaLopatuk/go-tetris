package game

import "fyne.io/fyne/v2"

const (
	MOVE_LEFT = iota
	MOVE_RIGHT
	ROTATE
	MOVE_DOWN
)

type Action struct {
	ActionType int
}

type InputController struct {
	OnAction func(action *Action)
}

func NewInputController(onAction func(action *Action)) *InputController {
	return &InputController{
		OnAction: onAction,
	}
}

func (i *InputController) HandleKey(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyLeft:
		i.OnAction(&Action{ActionType: MOVE_LEFT})
	case fyne.KeyRight:
		i.OnAction(&Action{ActionType: MOVE_RIGHT})
	case fyne.KeyUp:
		i.OnAction(&Action{ActionType: ROTATE})
	case fyne.KeyDown:
		i.OnAction(&Action{ActionType: MOVE_DOWN})
	}
}
