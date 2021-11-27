package scene

import (
	"github.com/bjatkin/chessRTS/control"
	"github.com/bjatkin/chessRTS/entity"
	"github.com/bjatkin/chessRTS/errors"
	"github.com/bjatkin/chessRTS/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Editor struct {
	Manager  *entity.Manger
	Camera   math.IVec2
	Keyboard *control.Keyboard
	Mouse    *control.Mouse
	Done     bool
}

func NewEditor() *Editor {
	Keyboard := control.NewKeyboard(control.MapWASD)
	Mouse := control.NewMouse()

	return &Editor{
		Manager:  entity.NewManger(Keyboard, Mouse),
		Camera:   math.IVec2{},
		Keyboard: control.NewKeyboard(control.MapWASD),
		Mouse:    control.NewMouse(),
	}
}

func (s *Editor) Init() error {
	return nil
}

func (s *Editor) Update() error {
	for _, ent := range s.Manager.All() {
		err := ent.Update()
		if err != nil {
			return errors.Wrap(err, errors.UpdateError, "there was an error updating the editors entities").
				Tag("entity", ent)
		}
	}
	return nil
}

func (s *Editor) Draw(screen *ebiten.Image) {
	for _, ent := range s.Manager.All() {
		ent.Draw(screen)
	}
}

func (s *Editor) Finished() bool {
	if s.Done {
		return true
	}
	return false
}
