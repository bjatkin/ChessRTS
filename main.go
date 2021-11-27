package main

import (
	"github.com/bjatkin/chessRTS/control"
	"github.com/bjatkin/chessRTS/entity"
	"github.com/bjatkin/chessRTS/game"
	"github.com/bjatkin/chessRTS/game/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// get the main game and setup ebiten
	main := &game.Game{
		Debug: true,
		Scene: &scene.TestScene{
			Manager: entity.NewManger(),
			Keyboard: control.NewKeyboard(
				map[ebiten.Key]int{
					ebiten.KeyW: control.W,
					ebiten.KeyA: control.A,
					ebiten.KeyS: control.S,
					ebiten.KeyD: control.D,
				},
			),
		},
	}

	game.Start(main)
}
