package entity

import (
	"math/rand"

	"github.com/bjatkin/chessRTS/entity/layer"
	"github.com/bjatkin/chessRTS/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Draw(*ebiten.Image)
	Update() error
	Trans() *Transform
	Layer() int
}

type Base struct {
	Transform *Transform
	GFX       *ebiten.Image
	Tick      int
	DrawLayer int
}

func NewBase(gfx *ebiten.Image, layer int) *Base {
	return &Base{
		Transform: &Transform{Pos: math.IVec2{}, Scale: 2},
		GFX:       gfx,
		DrawLayer: layer,
	}
}

func (e *Base) Draw(screen *ebiten.Image) {
	if e.GFX != nil {
		screen.DrawImage(e.GFX, e.Trans().DrawImageOptions())
	}
}

func (e *Base) Update() error {
	e.Tick++
	return nil
}

func (e *Base) Trans() *Transform {
	return e.Transform
}

func (e *Base) Layer() int {
	return e.DrawLayer
}

type Transform struct {
	Pos   math.IVec2
	Scale int
}

func (t *Transform) DrawImageOptions() *ebiten.DrawImageOptions {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(t.Scale), float64(t.Scale))
	op.GeoM.Translate(float64(t.Pos.X), float64(t.Pos.Y))
	return &op
}

type ID uint64

type Manger struct {
	Entities map[ID]entityLocation
	Layers   [][]Entity
}

type entityLocation struct {
	layer int
	index int
}

func NewManger(add ...Entity) *Manger {
	manager := &Manger{
		Entities: make(map[ID]entityLocation),
		Layers:   make([][]Entity, layer.Total),
	}

	for _, ent := range add {
		manager.Add(ent)
	}

	return manager
}

func (em *Manger) Add(ent Entity) ID {
	// get a unique random id
	id := rand.Int()
	for _, ok := em.Entities[ID(id)]; ok; id = rand.Int() {
	}

	em.Layers[ent.Layer()] = append(em.Layers[ent.Layer()], ent)
	idx := len(em.Layers) - 1
	em.Entities[ID(id)] = entityLocation{layer: ent.Layer(), index: idx}
	return ID(id)
}

func (em *Manger) Del(id ID) bool {
	location, ok := em.Entities[id]
	if !ok {
		return false
	}

	// remove the entity from the layer slice
	em.Layers[location.layer] = append(
		em.Layers[location.layer][:location.index],
		em.Layers[location.layer][location.index+1:]...,
	)

	// remove the location info from the map
	delete(em.Entities, id)
	return true
}

func (em *Manger) Get(id ID) (Entity, bool) {
	location, ok := em.Entities[id]
	if !ok {
		return nil, false
	}

	return em.Layers[location.layer][location.index], true
}

func (em *Manger) All() []Entity {
	var entities []Entity
	for _, layer := range em.Layers {
		for _, ent := range layer {
			entities = append(entities, ent)
		}
	}

	return entities
}
