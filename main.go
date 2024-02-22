package main

import (
	"strangeadventure/data"
	"strangeadventure/models/game"
	"strangeadventure/models/player"
	"strangeadventure/utils/clearscreen"
)

var Game = game.MakeGame()
var Player = player.PersonConstructor("")

func main() {

	for Game.IsOpen {

		data.GetLogic[Game.GetScreen()](&Game,Player)
		// fmt.Println(Game.GetScreen())
		// time.Sleep(5*time.Second)
		clearscreen.CallClear()

	}

}
