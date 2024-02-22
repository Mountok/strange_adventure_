package profile

import (
	"strangeadventure/models/game"
	"strangeadventure/models/player"
	"strangeadventure/utils/clearscreen"
	"strangeadventure/utils/readkeyboard"

	"github.com/eiannone/keyboard"
)


func Profile(game *game.Game, player player.Player) {
	clearscreen.CallClear()
	Render(player)
	handleEsc(game)
}

func handleEsc(game *game.Game) {
	for {
		err, key := readkeyboard.ReadKeyBoard()
		if err != nil {panic(err)}
		if key == keyboard.KeyEsc {
			game.SetScreen("menu")
			break
		}
	}
}