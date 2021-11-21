package prefabs

import (
	"github.com/bjatkin/chessRTS/assets/images"
	"github.com/bjatkin/chessRTS/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type MousePointer struct {
	*entity.Transform

	gfx *ebiten.Image
}

func NewMousePointer() *MousePointer {
	return &MousePointer{
		gfx: images.MustEbiten(images.CursorBig),
	}
}

func (mp *MousePointer) Draw(screen *ebiten.Image) {
	screen.DrawImage(mp.gfx, mp.DrawImageOptions())
}

func (mp *MousePointer) Update() error {
	x, y := ebiten.CursorPosition()
	mp.Pos.X, mp.Pos.Y = x, y
	return nil
}
