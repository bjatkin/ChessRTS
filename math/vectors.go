package math

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type FVec3 struct {
	X, Y, Z float64
}

func (vec *FVec3) FVec2() *FVec2 {
	return &FVec2{
		X: vec.X,
		Y: vec.Y,
	}
}

func (vec *FVec3) IVec2() *IVec2 {
	return &IVec2{
		X: int(vec.X),
		Y: int(vec.Y),
	}
}

func (vec *FVec3) IVec3() *IVec3 {
	return &IVec3{
		X: int(vec.X),
		Y: int(vec.Y),
		Z: int(vec.Z),
	}
}

type IVec3 struct {
	X, Y, Z int
}

func (vec *IVec3) FVec2() *FVec2 {
	return &FVec2{
		X: float64(vec.X),
		Y: float64(vec.Y),
	}
}

func (vec *IVec3) IVec2() *IVec2 {
	return &IVec2{
		X: vec.X,
		Y: vec.Y,
	}
}

func (vec *IVec3) FVec3() *FVec3 {
	return &FVec3{
		X: float64(vec.X),
		Y: float64(vec.Y),
		Z: float64(vec.Z),
	}
}

type FVec2 struct {
	X, Y float64
}

func (vec *FVec2) IVec2() *IVec2 {
	return &IVec2{
		X: int(vec.X),
		Y: int(vec.Y),
	}
}

func (vec *FVec2) IVec3() *IVec3 {
	return &IVec3{
		X: int(vec.X),
		Y: int(vec.Y),
	}
}

func (vec *FVec2) FVec3() *FVec3 {
	return &FVec3{
		X: vec.X,
		Y: vec.Y,
	}
}

type IVec2 struct {
	X, Y int
}

func (vec *IVec2) DrawImageOptions() *ebiten.DrawImageOptions {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(vec.X), float64(vec.Y))
	return opt
}

func (vec *IVec2) Add(add IVec2) IVec2 {
	return IVec2{
		X: vec.X + add.X,
		Y: vec.Y + add.Y,
	}
}

func (vec *IVec2) FVec2() *FVec2 {
	return &FVec2{
		X: float64(vec.X),
		Y: float64(vec.Y),
	}
}

func (vec *IVec2) IVec3() *IVec3 {
	return &IVec3{
		X: int(vec.X),
		Y: int(vec.Y),
	}
}

func (vec *IVec2) FVec3() *FVec3 {
	return &FVec3{
		X: float64(vec.X),
		Y: float64(vec.Y),
	}
}
