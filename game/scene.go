package game

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Init() error
	Update() error
	Draw(screen *ebiten.Image)
	Finished() bool
}
