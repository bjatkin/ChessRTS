package control

import (
	"github.com/bjatkin/chessRTS/assets/images"
	"github.com/bjatkin/chessRTS/entity"
	"github.com/bjatkin/chessRTS/entity/layer"
	"github.com/bjatkin/chessRTS/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type MouseMapping struct {
	Local int
	Btn   ebiten.MouseButton
}

type Mouse struct {
	*entity.Base
	DownLast map[int]bool
	DownNow  map[int]bool
	Convert  map[ebiten.MouseButton]int
}

func NewMouse() *Mouse {
	return &Mouse{
		Base: &entity.Base{
			Transform: &entity.Transform{
				Pos:   math.IVec2{},
				Scale: 2,
			},
			GFX:       images.MustEbiten(images.CursorBig),
			DrawLayer: layer.Mouse,
		},
	}
}

func (m *Mouse) Draw(screen *ebiten.Image) {
	screen.DrawImage(m.GFX, m.Base.Transform.DrawImageOptions())
}

func (m *Mouse) Update() error {
	// update the mouse button state
	for i := range m.DownLast {
		m.DownLast[i] = false
	}

	for i := range m.DownNow {
		m.DownLast[i] = m.DownNow[i]
		m.DownNow[i] = false
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if btn, ok := m.Convert[ebiten.MouseButtonLeft]; ok {
			m.DownNow[btn] = true
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		if btn, ok := m.Convert[ebiten.MouseButtonMiddle]; ok {
			m.DownNow[btn] = true
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if btn, ok := m.Convert[ebiten.MouseButtonRight]; ok {
			m.DownNow[btn] = true
		}
	}

	// set the cursor position
	x, y := ebiten.CursorPosition()
	m.Base.Transform.Pos.X, m.Base.Transform.Pos.Y = x, y

	return nil
}

func (m *Mouse) Pressed(code int) bool {
	return !m.DownLast[code] && m.DownNow[code]
}

func (m *Mouse) Down(code int) bool {
	return m.DownNow[code]
}

func (m *Mouse) Release(code int) bool {
	return m.DownLast[code] && !m.DownNow[code]
}
