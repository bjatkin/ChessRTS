package prefabs

import (
	"image"

	"github.com/bjatkin/chessRTS/assets/images"
	"github.com/bjatkin/chessRTS/entity"
	"github.com/bjatkin/chessRTS/entity/layer"

	"github.com/hajimehoshi/ebiten/v2"
)

// Piece Type
const (
	Pawn = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

// Piece Color
// use a large power of two so color can be used as a flag
const (
	Black = 0
	White = 128
)

type Piece struct {
	*entity.Base

	class int
	color int
}

func NewPiece(color, class int) *Piece {
	gfx := images.MustEbiten(images.Pieces)
	switch class | color {
	case Pawn | White:
		gfx = gfx.SubImage(image.Rect(100, 20, 119, 39)).(*ebiten.Image)
	case Pawn | Black:
		gfx = gfx.SubImage(image.Rect(0, 20, 19, 39)).(*ebiten.Image)
	case Knight | White:
		gfx = gfx.SubImage(image.Rect(60, 20, 79, 39)).(*ebiten.Image)
	case Knight | Black:
		gfx = gfx.SubImage(image.Rect(80, 20, 99, 39)).(*ebiten.Image)
	case Bishop | White:
		gfx = gfx.SubImage(image.Rect(40, 0, 59, 19)).(*ebiten.Image)
	case Bishop | Black:
		gfx = gfx.SubImage(image.Rect(60, 0, 79, 19)).(*ebiten.Image)
	case Rook | White:
		gfx = gfx.SubImage(image.Rect(80, 0, 99, 19)).(*ebiten.Image)
	case Rook | Black:
		gfx = gfx.SubImage(image.Rect(100, 0, 119, 19)).(*ebiten.Image)
	case Queen | White:
		gfx = gfx.SubImage(image.Rect(20, 20, 39, 39)).(*ebiten.Image)
	case Queen | Black:
		gfx = gfx.SubImage(image.Rect(40, 20, 59, 39)).(*ebiten.Image)
	case King | White:
		gfx = gfx.SubImage(image.Rect(0, 0, 19, 19)).(*ebiten.Image)
	case King | Black:
		gfx = gfx.SubImage(image.Rect(20, 0, 39, 19)).(*ebiten.Image)
	}

	return &Piece{
		Base:  entity.NewBase(gfx, layer.Foreground),
		class: class,
		color: color,
	}
}
