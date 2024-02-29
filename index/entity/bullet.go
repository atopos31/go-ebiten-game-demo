package entity

import (
	"bytes"
	"demo/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	X      float64
	Y      float64
	Img    *ebiten.Image
	Weight int
	Height int
}

func newBullet(x float64, y float64) *Bullet {
	img, _, _ := ebitenutil.NewImageFromReader(bytes.NewReader(assets.Imgbullet))
	bulletWeight := img.Bounds().Size().X
	bulletHeight := img.Bounds().Size().Y
	bullet := &Bullet{
		X:      x - float64(bulletWeight)/2 + 1,
		Y:      y - float64(bulletHeight)/2,
		Img:    img,
		Weight: bulletWeight,
		Height: bulletHeight,
	}

	return bullet
}

func (bullet *Bullet) Update() {
	bullet.Y -= 10
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.X, bullet.Y)
	screen.DrawImage(bullet.Img, op)
}
