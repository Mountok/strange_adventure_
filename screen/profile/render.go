package profile

import (
	"fmt"
	"strangeadventure/data/frames"
	"strangeadventure/models/player"
)


func Render(player player.Player) {
	fmt.Printf(frames.GetScreenWithPlayerInfo["profile"](player))
}