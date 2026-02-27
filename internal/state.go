package game

type State struct {
	IsRunning  bool
	IsGameOver bool
}

func CreateState() *State {
	return &State{IsRunning: true, IsGameOver: false}
}
