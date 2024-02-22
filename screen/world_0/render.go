package world0

import (
	"fmt"
	"strangeadventure/data/frames"
	"strangeadventure/utils/clearscreen"
)

func Render(selectPlace int) {
	clearscreen.CallClear()
	fmt.Printf(frames.GetScreen[fmt.Sprintf("world_0_%d",selectPlace)]())
}