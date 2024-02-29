package main

import (
	"bytes"
	"demo/assets"
	"demo/entity"
	"fmt"
	_ "image/png"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	hero      *entity.Hero
	mu        *sync.Mutex
	airPlanes map[*entity.AirPlane]any
	score     int
}

func (g *Game) Update() error {
	g.hero.Update()
	g.mu.Lock()
	for v := range g.hero.Bullets {
		v.Update()
	}
	for v := range g.airPlanes {
		v.Update()
	}
	g.mu.Unlock()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//背景图
	op := &ebiten.DrawImageOptions{}
	back,_,_:=ebitenutil.NewImageFromReader(bytes.NewReader(assets.ImgBack))
	screen.DrawImage(back, op)
	g.mu.Lock()
	for v := range g.hero.Bullets {
		if v.Y < -10 {
			delete(g.hero.Bullets, v)
		}
		v.Draw(screen)
	}
	for v := range g.airPlanes {
		if v.Y > 860 {
			delete(g.airPlanes, v)
		}
		v.Draw(screen)
	}
	for air := range g.airPlanes {
		for bul := range g.hero.Bullets {
			if checkCollision(air, bul) {
				delete(g.hero.Bullets, bul)
				delete(g.airPlanes, air)
				g.score++
			}
		}
	}
	g.mu.Unlock()
	g.hero.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nSCORE:%d", ebiten.ActualTPS(), ebiten.ActualFPS(), g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 852
}

func checkCollision(airplane *entity.AirPlane, bullet *entity.Bullet) bool {
	return !(airplane.X > bullet.X+float64(bullet.Weight) ||
		airplane.X+float64(airplane.Weight) < bullet.X ||
		airplane.Y > bullet.Y+float64(bullet.Height) ||
		airplane.Y+float64(airplane.Height) < bullet.Y)
}

func main() {
	ebiten.SetWindowSize(480, 852)
	ebiten.SetWindowTitle("飞机大战")
	air := map[*entity.AirPlane]any{}
	game := &Game{hero: entity.NewHero(480, 852), mu: &sync.Mutex{}, airPlanes: air}
	entity.StartAirPlane(air, game.mu)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
