package menu

import (
	"fmt"
	"strangeadventure/data/frames"
	"strangeadventure/utils/clearscreen"
)


func Render(menuIndex int) {
	clearscreen.CallClear()
	var str string = frames.GetScreen[fmt.Sprintf("menu_%d",menuIndex)]();
	fmt.Println(str)
}