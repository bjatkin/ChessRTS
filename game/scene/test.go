package scene

import (
	"github.com/bjatkin/chessRTS/controller"
	"github.com/bjatkin/chessRTS/entity"
	"github.com/bjatkin/chessRTS/entity/prefabs"
	"github.com/bjatkin/chessRTS/errors"
	"github.com/bjatkin/chessRTS/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type TestScene struct {
	Manager  *entity.EntityManger
	Map      entity.ID
	Camera   math.IVec2
	Keyboard *controller.Keyboard
}

// nodes determines the map layout
var nodes = []prefabs.Node{
	{Pos: math.IVec2{X: 1, Y: 1}, Class: prefabs.BlackNode, AdjacentIdx: [8]int{1, 2, 3, 4, 5, 6, 7, 8}},

	{Pos: math.IVec2{X: 0, Y: 1}, Class: prefabs.WhiteNode, AdjacentIdx: [8]int{-1, 6, 0, 5, -1, -1, 2, 4}},
	{Pos: math.IVec2{X: 1, Y: 0}, Class: prefabs.WhiteNode, AdjacentIdx: [8]int{6, -1, 7, 0, 1, -1, -1, 3}},
	{Pos: math.IVec2{X: 2, Y: 1}, Class: prefabs.WhiteNode, AdjacentIdx: [8]int{0, 7, -1, 8, 4, 2, -1, -1}},
	{Pos: math.IVec2{X: 1, Y: 2}, Class: prefabs.WhiteNode, AdjacentIdx: [8]int{5, 0, 8, -1, -1, 1, 3, -1}},

	{Pos: math.IVec2{X: 0, Y: 2}, Class: prefabs.BlackNode, AdjacentIdx: [8]int{-1, 1, 4, -1, -1, -1, 0, -1}},
	{Pos: math.IVec2{X: 0, Y: 0}, Class: prefabs.BlackNode, AdjacentIdx: [8]int{-1, -1, 2, 1, -1, -1, -1, 0}},
	{Pos: math.IVec2{X: 2, Y: 0}, Class: prefabs.BlackNode, AdjacentIdx: [8]int{2, -1, -1, 3, 0, -1, -1, -1}},
	{Pos: math.IVec2{X: 2, Y: 2}, Class: prefabs.BlackNode, AdjacentIdx: [8]int{4, 3, -1, -1, -1, 0, -1, -1}},
}

func (s *TestScene) Init() error {
	king := prefabs.NewPiece(prefabs.Black, prefabs.King)
	king.Transform.Pos = math.IVec2{X: 241, Y: 238}
	s.Manager.Add(king)

	m := prefabs.NewMap(nodes)
	m.Transform.Pos = math.IVec2{X: 200, Y: 200}
	s.Map = s.Manager.Add(m)

	return nil
}

func (s *TestScene) Update() error {
	s.Keyboard.Update()

	if s.Keyboard.Down(controller.A) {
		s.Camera.X++
	}
	if s.Keyboard.Down(controller.D) {
		s.Camera.X--
	}
	if s.Keyboard.Down(controller.W) {
		s.Camera.Y++
	}
	if s.Keyboard.Down(controller.S) {
		s.Camera.Y--
	}

	// run all the entity updates
	for id, e := range s.Manager.All() {
		err := e.Update()
		if err != nil {
			return errors.Wrap(err, errors.UpdateError, "there was an error updateing entities in teh test scene").Tag("entity id", id)
		}
	}

	return nil
}

func (s *TestScene) Draw(screen *ebiten.Image) {
	// draw all the scene entities
	for _, e := range s.Manager.All() {
		tmp := e.Trans().Pos.Add(s.Camera)
		e.Draw(screen, &entity.Transform{Pos: tmp, Scale: e.Trans().Scale})
	}
}

func (s *TestScene) Finished() bool {
	// for now, lets set this scene up to never yield
	return false
}
