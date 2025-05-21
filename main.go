package main

import (
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	avatar    *ebiten.Image
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
		g.obstacleX = float64(rand.Intn(screenWidth - obstacleLength))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// player
	// vector.DrawFilledRect(screen, float32(g.playerX), playerHeight, obstacleLength, obstacleLength, color.White, false)
	op := &ebiten.DrawImageOptions{}

	// 1. 원본 이미지 크기 확인
	w, h := g.avatar.Bounds().Dx(), g.avatar.Bounds().Dy()

	// 2. 20x20으로 축소 비율 계산
	scaleX := 20.0 / float64(w)
	scaleY := 20.0 / float64(h)

	// 3. 스케일 적용
	op.GeoM.Scale(scaleX, scaleY)

	// 4. 위치 설정 (예: g.playerX, playerHeight 위치에 그림)
	op.GeoM.Translate(g.playerX, playerHeight)

	// 5. 그리기
	screen.DrawImage(g.avatar, op)

	// obstacle
	vector.DrawFilledRect(screen, float32(g.obstacleX), float32(g.obstacleY), obstacleLength, obstacleLength, color.RGBA{255, 0, 0, 255}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		playerX:   float64(screenWidth / 2),
		obstacleX: 100,
		obstacleY: 0,
	}

	imgFile, err := os.Open("assets/gopher.png")
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	g.avatar = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Dodge the Falling Object")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
