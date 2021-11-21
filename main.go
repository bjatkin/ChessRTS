package main

import (
	"github.com/bjatkin/chessRTS/controller"
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
			Manager: entity.NewEntityManger(),
			Keyboard: controller.NewKeyboard(
				[]controller.Mapping{
					{Key: ebiten.KeyW, Local: controller.W},
					{Key: ebiten.KeyA, Local: controller.A},
					{Key: ebiten.KeyS, Local: controller.S},
					{Key: ebiten.KeyD, Local: controller.D},
				},
			),
		},
	}

	game.Start(main)
}
