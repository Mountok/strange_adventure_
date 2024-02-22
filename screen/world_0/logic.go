package world0

import (
	"strangeadventure/models/game"
	"strangeadventure/models/player"
	"strangeadventure/utils/readkeyboard"

	"github.com/eiannone/keyboard"
)

func World0(g *game.Game, p player.Player) {
	var selectPlace int = 0
	Render(selectPlace)
	handleArrowsClick(selectPlace)
	g.SetScreen("menu")
}


func handleArrowsClick(selectIndex int) (int,error) {
	for {
		err, key := readkeyboard.ReadKeyBoard()
		if err != nil {panic(err)}
		switch key {
		case keyboard.KeyArrowUp:
			if selectIndex > 0{
				selectIndex -= 1
				Render(selectIndex)
			}
		case keyboard.KeyArrowDown:
			if selectIndex < 3 {
				selectIndex += 1
				Render(selectIndex)
			}
		case keyboard.KeyEnter:
			return selectIndex, nil
		}
	}
}