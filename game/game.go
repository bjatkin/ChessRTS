package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/bjatkin/chessRTS/errors"
)

func Start(game *Game) {
	err := ebiten.RunGame(game)
	if err != nil {
		Exit(err)
	}
}

func Exit(err error) {
	// run any clean up scripts here
	fmt.Println("~EXITING GAME~")
	fmt.Println(err)

	code := errors.Unknown
	if cErr, ok := err.(*errors.Err); ok {
		code = cErr.Code
	}

	panic(code)
}

type Game struct {
	Ticks        int64
	Debug        bool
	exitError    error
	Scene        Scene
	sceneStarted bool
	sceneQue     Scene
}

func NewGame(scene Scene) *Game {
	return &Game{
		Scene: scene,
	}
}

func (g *Game) Update() error {
	if g.exitError != nil {
		return g.exitError
	}
	g.Ticks++

	// make sure to start the scene
	if !g.sceneStarted {
		err := g.Scene.Init()
		if err != nil {
			errors.Wrap(err, errors.SetupError, "unable to init scene")
		}
		g.sceneStarted = true
	}

	err := g.Scene.Update()
	if err != nil {
		errors.Wrap(err, errors.UpdateError, "there was an error running the scene update")
	}

	if g.sceneQue != nil && g.Scene.Finished() {
		g.Scene = g.sceneQue
		g.sceneQue = nil
		g.sceneStarted = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Debug {
		// print out general debugging here
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	}

	g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func (g *Game) QueNextScene(new Scene) {
	g.sceneQue = new
}
