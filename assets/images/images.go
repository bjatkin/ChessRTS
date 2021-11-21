package images

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/bjatkin/chessRTS/errors"
	"github.com/bjatkin/chessRTS/game"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed png/small_pointer.png
	CursorSmall []byte

	//go:embed png/big_pointer.png
	CursorBig []byte

	//go:embed png/font.png
	Font []byte

	//go:embed png/pieces.png
	Pieces []byte

	//go:embed png/light_square.png
	LightSquare []byte
	//go:embed png/dark_square.png
	DarkSquare []byte
)

func Ebiten(imgBytes []byte) (*ebiten.Image, error) {
	if len(imgBytes) == 0 {
		return nil, errors.New(errors.InvalidArgument, "unable to convert bytes into an ebiten image, recieved empty byte slice")
	}

	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, errors.Wrap(err, errors.SetupError, "unable to convert bytes into an ebiten image")
	}

	return ebiten.NewImageFromImage(img), err
}

func MustEbiten(imgBytes []byte) *ebiten.Image {
	if len(imgBytes) == 0 {
		game.Exit(errors.New(errors.InvalidArgument, "unable to convert bytes into an ebiten image, recieved empty byte slice"))
	}

	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		game.Exit(errors.Wrap(err, errors.SetupError, "unable to convert bytes into an ebiten image"))
	}

	return ebiten.NewImageFromImage(img)
}
