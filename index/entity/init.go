package entity

import "github.com/hajimehoshi/ebiten/v2"

type Eler interface {
	Update()
	Draw(screen *ebiten.Image)
}