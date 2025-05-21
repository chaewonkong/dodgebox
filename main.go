package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth    = 320
	screenHeight   = 240
	playerHeight   = 200
	obstacleLength = 20
)

type Game struct {
	playerX   float64
	obstacleX float64
	obstacleY float64
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 2
	}

	g.obstacleY += 2
	if g.obstacleY > screenHeight {
		g.obstacleY = 0
		g.obstacleX = float64(rand.Intn(screenWidth - obstacleLength))
	}

	if g.obstacleY > playerHeight && g.playerX < g.obstacleX+obstacleLength && g.playerX > g.obstacleX-obstacleLength {
		g.obstacleY = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// player
	ebitenutil.DrawRect(screen, g.playerX, playerHeight, obstacleLength, obstacleLength, color.White)

	// obstacle
	ebitenutil.DrawRect(screen, g.obstacleX, g.obstacleY, obstacleLength, obstacleLength, color.RGBA{255, 0, 0, 255})

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := &Game{
		playerX:   float64(screenWidth / 2),
		obstacleX: 100,
		obstacleY: 0,
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Dodge the Falling Object")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
