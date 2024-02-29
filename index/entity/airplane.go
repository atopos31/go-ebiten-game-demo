package entity

import (
	"bytes"
	"demo/assets"
	"math/rand"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AirPlane struct {
	X      float64
	Y      float64
	Img    *ebiten.Image
	Weight int
	Height int
}

func newAirPlane(x float64, y float64) *AirPlane {
	img, _, _ := ebitenutil.NewImageFromReader(bytes.NewReader(assets.ImgAirPlane))
	heroWeight := img.Bounds().Size().X
	heroHeight := img.Bounds().Size().Y
	airplane := &AirPlane{
		X:      x,
		Y:      y,
		Img:    img,
		Weight: heroWeight,
		Height: heroHeight,
	}

	return airplane
}

func StartAirPlane(airplanes map[*AirPlane]any, mu *sync.Mutex) {
	go func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			x := r.Int63n(430)
			mu.Lock()
			airplanes[newAirPlane(float64(x), -20)] = struct{}{}
			mu.Unlock()
		}

	}()
}

func (airplane *AirPlane) Update() {
	airplane.Y += 1.5
}

func (airplane *AirPlane) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(airplane.X, airplane.Y)
	screen.DrawImage(airplane.Img, op)
}
