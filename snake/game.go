package snake

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/bitmapfont/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 600
	boardRows    = 20
	boardCols    = 20
)

var (
	backgroundColor = color.RGBA{128, 128, 128, 255}
	snakeHeadColor  = color.RGBA{255, 99, 71, 255}
	snakeBodyColor  = color.RGBA{149, 236, 105, 255}
	// foodColor       = color.RGBA{128, 0, 128, 255}
)

type Game struct {
	input *Input
	board *Board
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
		board: NewBoard(boardRows, boardCols),
	}
}

func (g *Game) Update() error {
	return g.board.Update(g.input)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	if g.board.gameOver {
		// 游戏结束 显示分数
		text.Draw(
			screen,
			fmt.Sprintf("Game Over. Score: %d", g.board.scores),
			bitmapfont.Face,
			ScreenWidth/2-50, ScreenHeight/2-50,
			color.Black,
		)
	} else {
		width := ScreenHeight / boardRows
		// 画🐍身
		snakeColor := snakeBodyColor
		for i, p := range g.board.snake.body {
			if i == len(g.board.snake.body)-1 {
				snakeColor = snakeHeadColor
			}
			ebitenutil.DrawRect(screen, float64(p.y*width)+float64(width*1/20), float64(p.x*width)+float64(width*1/20), float64(width)*9/10, float64(width)*9/10, snakeColor)

		}
		// 画食物
		// if g.board.food != nil {
		// 	ebitenutil.DrawRect(screen, float64(g.board.food.y*width), float64(g.board.food.x*width), float64(width), float64(width), foodColor)
		// }
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.board.scores))
		var foodImg *ebiten.Image
		var err error
		if foodImg, _, err = ebitenutil.NewImageFromFile("./foods/apple.png"); err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{}
		sx, sy := foodImg.Size()
		propx := float64(width) / float64(sy)
		propy := float64(width) / float64(sx)
		op.GeoM.Scale(propx, propy)
		op.GeoM.Translate(float64(g.board.food.y*width), float64(g.board.food.x*width))
		screen.DrawImage(foodImg, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
