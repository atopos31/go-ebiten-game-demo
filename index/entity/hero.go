package entity

import (
	"bytes"
	"demo/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Hero struct {
	x             float64
	y             float64
	Img           *ebiten.Image
	Weight        int
	Height        int
	lastShootTime time.Time
	Bullets       map[*Bullet]any
}

func NewHero(weight int, height int) *Hero {
	img1, _, _ := ebitenutil.NewImageFromReader(bytes.NewReader(assets.Imghero1))
	img0, _, _ := ebitenutil.NewImageFromReader(bytes.NewReader(assets.Imghero0))
	heroWeight := img1.Bounds().Size().X
	heroHeight := img1.Bounds().Size().Y
	Hero := &Hero{
		x:       float64(weight-heroWeight) / 2,
		y:       float64(height - heroHeight),
		Img:     img1,
		Weight:  heroWeight,
		Height:  heroHeight,
		Bullets: map[*Bullet]any{},
	}
	go func() {
		//英雄机喷气效果
		ticker := time.NewTicker(70 * time.Millisecond)
		defer ticker.Stop()
		i := 1

		for range ticker.C {
			if i == 1 {
				Hero.Img = img1
			} else {
				Hero.Img = img0
			}
			i = 3 - i
		}
	}()

	return Hero
}

func (Hero *Hero) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) && Hero.x > 0 {
		Hero.x -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && Hero.x < float64(480-Hero.Weight) {
		Hero.x += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		Hero.y -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		Hero.y += 1.5
	}

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		if time.Since(Hero.lastShootTime).Seconds() >= 0.2 {
			bullet := newBullet(Hero.x+float64(Hero.Weight)/2+32, Hero.y+40)
			Hero.Bullets[bullet] = struct{}{}
			bullet = newBullet(Hero.x+float64(Hero.Weight)/2-32, Hero.y+40)
			Hero.Bullets[bullet] = struct{}{}
			Hero.lastShootTime = time.Now() // 更新时间戳
		}
	}
}

func (Hero *Hero) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Hero.x, Hero.y)
	screen.DrawImage(Hero.Img, op)
}
