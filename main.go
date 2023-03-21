package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yushengguo557/snake/snake"
)

func main() {
	os.Setenv("CGO_ENABLED", "1")
	game := snake.NewGame()
	ebiten.SetWindowSize(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowTitle("Snake")
	var icon image.Image
	var err error
	if _, icon, err = ebitenutil.NewImageFromFile("assets/icon.png"); err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowIcon([]image.Image{icon})
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
