package setname

import (
	"strangeadventure/models/game"
	"strangeadventure/models/player"
	"strangeadventure/utils/clearscreen"
	"strangeadventure/utils/readinter"
)


func InterName(g *game.Game, p player.Player) {
	clearscreen.CallClear()
	Render()
	
	name := GetName()
	p.SetName(name)
	g.SetScreen("menu")
}

func GetName() string {
	name, err := readinter.ReadString()
	if err != nil {panic(err)}
	return name
}