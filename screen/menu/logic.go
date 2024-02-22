package menu

import (
	"strangeadventure/models/game"
	"strangeadventure/models/player"
	"strangeadventure/screen/setname"
	"strangeadventure/utils/readkeyboard"

	"github.com/eiannone/keyboard"
)


func Menu(g *game.Game, p player.Player) {
	var menuIndex int = 1
	Render(menuIndex)
	
	for {
		menuIndex, err := handleArrowsClick(menuIndex)
		if err != nil {panic(err)}
		handleMenuSelect(menuIndex,g,p)
		break
	}

	// game.SetScreen("profile")	
}


func handleArrowsClick(menuIndex int) (int,error) {
	for {
		err, key := readkeyboard.ReadKeyBoard()
		if err != nil {panic(err)}
		switch key {
		case keyboard.KeyArrowUp:
			if menuIndex > 1{
				menuIndex -= 1
				Render(menuIndex)
			}
		case keyboard.KeyArrowDown:
			if menuIndex < 4 {
				menuIndex += 1
				Render(menuIndex)
			}
		case keyboard.KeyEnter:
			return menuIndex, nil
		}
	}
}
func handleMenuSelect(menuIndex int,game *game.Game,p player.Player) {
	if menuIndex == 1 {
		if p.GetInfo().Name == "" {
			setname.InterName(game,p)
			
		} else {
			game.SetScreen("world_0")
		}
	} else if menuIndex == 2 {
		game.SetScreen("profile")
	} else if menuIndex == 4 {
		game.Close()
	}
}