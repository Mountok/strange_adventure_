package setname

import (
	"fmt"
	"strangeadventure/data/frames"
)


func Render() {
	fmt.Printf(frames.GetScreen["inter_name"]())
}