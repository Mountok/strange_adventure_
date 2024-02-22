package game


type Game struct {
	Screen string
	IsOpen bool
	W,H int
}

func (g *Game) Close() {
	g.IsOpen = false
}
func (g *Game) SetScreen(screen string) {
	g.Screen = screen
}
func (g *Game) GetScreen() string{
	return g.Screen 
}

func MakeGame() Game {
	game := Game{
		IsOpen: true,
		Screen: "menu",
		W: 101,
		H: 55,
	}
	return game
}