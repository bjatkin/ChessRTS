package fonts

import (
	"fmt"
	"image"

	"github.com/bjatkin/chessRTS/errors"
	"github.com/bjatkin/chessRTS/game"
	"github.com/bjatkin/chessRTS/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Rune struct {
	Ref rune

	// width and height of the charater measured from the bottom left corner
	Size math.IVec2

	// horizontal and vertical offsets measured from the bottom left corner
	// used for spacing
	Offset math.IVec2

	// don't print these characters
	WhiteSpace bool
	NewLine    bool

	// gfx will be automatically set by the writer
	GFX *ebiten.Image
}

func (r *Rune) Print(pt math.IVec2, screen *ebiten.Image) {
	if r.WhiteSpace || r.NewLine {
		return
	}

	screen.DrawImage(r.GFX, pt.DrawImageOptions())
}

type Font struct {
	Atlas      *ebiten.Image
	SpriteSize math.IVec2
	Runes      map[rune]Rune
	LineHeight int
}

func NewFont(spriteSize math.IVec2, runeAtlas *ebiten.Image, runes []Rune) *Font {
	font := &Font{
		Atlas:      runeAtlas,
		SpriteSize: spriteSize,
		Runes:      make(map[rune]Rune, len(runes)),
	}

	// set up the runes
	width, _ := runeAtlas.Size()
	widthCount := width / spriteSize.X
	for i := 0; i < len(runes); i++ {
		r := &runes[i]

		x0 := (i % widthCount) * spriteSize.X
		y0 := (i / widthCount) * spriteSize.Y

		x1 := x0 + spriteSize.X
		y1 := y0 + spriteSize.Y

		r.GFX = runeAtlas.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
		if r.GFX == nil {
			game.Exit(errors.New(errors.SetupError, "rune could not find a valid sub image from the rune atlas").Tag("ref", r.Ref))
		}
		_, ok := font.Runes[r.Ref]
		if ok {
			game.Exit(errors.New(errors.SetupError, "rune ref already exists").Tag("ref", r.Ref))
		}

		font.Runes[r.Ref] = runes[i]

		// set the deafult line height
		h := int(float64(r.Size.Y) * 1.33)
		if h > font.LineHeight {
			font.LineHeight = h
		}
	}

	return font
}

func (f *Font) Print(pt math.IVec2, s string, screen *ebiten.Image) {
	offset := math.IVec2{}
	for _, r := range s {
		aRune, ok := f.Runes[r]
		if !ok {
			fmt.Println(errors.New(errors.DrawError, "can't print unknown character").Tag("character", string(r)))
			panic(errors.DrawError)
		}

		offset.Y -= aRune.Offset.Y
		aRune.Print(pt.Add(offset), screen)
		offset.Y += aRune.Offset.Y

		// move the cursor
		offset.X += aRune.Size.X + aRune.Offset.X
		if aRune.NewLine {
			offset.X = 0
			offset.Y += f.LineHeight
		}
	}
}
