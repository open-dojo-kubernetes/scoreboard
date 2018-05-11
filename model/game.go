package model

import "fmt"

type Game struct {
	Id string `json:"id"`
	LeftSideScore Score `json:"left_score"`
	RightSideScore Score `json:"right_score"`
	Winner string `json:"winner"`
}

type Score struct {
	Game int `json:"game"`
	Sets int `json:"sets"`
}

func (game Game) String() {
	fmt.Sprintf("Game: %v, Winner: %v", game.Id, game.Winner)
}