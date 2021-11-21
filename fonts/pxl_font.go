package fonts

import (
	"github.com/bjatkin/chessRTS/assets/images"
	"github.com/bjatkin/chessRTS/errors"
	"github.com/bjatkin/chessRTS/game"
	"github.com/bjatkin/chessRTS/math"
)

var PixelFont *Font

func Init() {

	// Set Up the Pixel Font
	fontImg, err := images.Ebiten(images.Font)
	if err != nil {
		game.Exit(errors.Wrap(err, errors.SetupError, "failed to init fonts"))
	}

	// Font Rune
	runes := []Rune{
		{Ref: 'A', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'B', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'C', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'D', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'E', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'F', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'G', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'H', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'I', Size: math.IVec2{X: 3, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'J', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'K', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'L', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'M', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'N', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'O', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'P', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'Q', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'R', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'S', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'T', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'U', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'V', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'W', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'X', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'Y', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'Z', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'a', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'b', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'c', Size: math.IVec2{X: 3, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'd', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'e', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'f', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'g', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: -2}},
		{Ref: 'h', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'i', Size: math.IVec2{X: 1, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'j', Size: math.IVec2{X: 3, Y: 7}, Offset: math.IVec2{X: 1, Y: -2}},
		{Ref: 'k', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'l', Size: math.IVec2{X: 2, Y: 6}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'm', Size: math.IVec2{X: 5, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'n', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'o', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'p', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: -1}},
		{Ref: 'q', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: -1}},
		{Ref: 'r', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 's', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 't', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'u', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'v', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'w', Size: math.IVec2{X: 5, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'x', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: 'y', Size: math.IVec2{X: 4, Y: 6}, Offset: math.IVec2{X: 1, Y: -1}},
		{Ref: 'z', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '~', Size: math.IVec2{X: 5, Y: 4}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '`', Size: math.IVec2{X: 2, Y: 8}, Offset: math.IVec2{X: 1, Y: 1}},
		{Ref: '!', Size: math.IVec2{X: 3, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '@', Size: math.IVec2{X: 7, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '#', Size: math.IVec2{X: 8, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '$', Size: math.IVec2{X: 7, Y: 7}, Offset: math.IVec2{X: 1, Y: 1}},
		{Ref: '%', Size: math.IVec2{X: 8, Y: 8}, Offset: math.IVec2{X: 1, Y: 1}},
		{Ref: '^', Size: math.IVec2{X: 5, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '&', Size: math.IVec2{X: 7, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '*', Size: math.IVec2{X: 4, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '(', Size: math.IVec2{X: 3, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: ')', Size: math.IVec2{X: 3, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '-', Size: math.IVec2{X: 5, Y: 4}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '_', Size: math.IVec2{X: 5, Y: 2}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '+', Size: math.IVec2{X: 6, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '=', Size: math.IVec2{X: 5, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '|', Size: math.IVec2{X: 2, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '{', Size: math.IVec2{X: 4, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '}', Size: math.IVec2{X: 4, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '\\', Size: math.IVec2{X: 5, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '/', Size: math.IVec2{X: 5, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: ';', Size: math.IVec2{X: 2, Y: 6}, Offset: math.IVec2{X: 1, Y: -1}},
		{Ref: ':', Size: math.IVec2{X: 2, Y: 5}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '"', Size: math.IVec2{X: 3, Y: 8}, Offset: math.IVec2{X: 1, Y: 1}},
		{Ref: '\'', Size: math.IVec2{X: 1, Y: 8}, Offset: math.IVec2{X: 1, Y: 1}},
		{Ref: ',', Size: math.IVec2{X: 2, Y: 2}, Offset: math.IVec2{X: 1, Y: -1}},
		{Ref: '.', Size: math.IVec2{X: 2, Y: 2}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '?', Size: math.IVec2{X: 6, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '[', Size: math.IVec2{X: 3, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: ']', Size: math.IVec2{X: 3, Y: 8}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '1', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '2', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '3', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '4', Size: math.IVec2{X: 4, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '5', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '6', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '7', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '8', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '9', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: '0', Size: math.IVec2{X: 5, Y: 7}, Offset: math.IVec2{X: 1, Y: 0}},
		{Ref: ' ', Size: math.IVec2{X: 5, Y: 7}, WhiteSpace: true},
		{Ref: '\t', Size: math.IVec2{X: 20, Y: 7}, WhiteSpace: true},
		{Ref: '\n', NewLine: true},
	}

	PixelFont = NewFont(
		math.IVec2{X: 8, Y: 8},
		fontImg,
		runes,
	)
}
