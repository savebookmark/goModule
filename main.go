package main

import (
	"log"

	"github.com/Workiva/go-datastructures/set"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game ..
type Game struct{}

// Update ..
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw ..
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout ..
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	list := arraylist.New()
	list.Add("list")

	set := set.New()
	set.Add("set")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
