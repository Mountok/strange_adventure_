package data

import (
	"strangeadventure/models/game"
	"strangeadventure/models/player"
	"strangeadventure/screen/menu"
	"strangeadventure/screen/profile"
	world0 "strangeadventure/screen/world_0"
)

var GetLogic = map[string]func(g *game.Game,p player.Player){
	"menu": menu.Menu,
	"profile": profile.Profile,
	"world_0": world0.World0,
}
