package control

import (
	"github.com/bjatkin/chessRTS/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var MapWASD = map[ebiten.Key]int{
	ebiten.KeyW: W,
	ebiten.KeyA: A,
	ebiten.KeyS: S,
	ebiten.KeyD: D,
}

// local key codes
const (
	W = iota
	A
	S
	D
)

type Keyboard struct {
	*entity.Base
	keys     []ebiten.Key
	DownLast map[int]bool
	DownNow  map[int]bool
	Convert  map[ebiten.Key]int
}

func NewKeyboard(mapping map[ebiten.Key]int) *Keyboard {
	return &Keyboard{
		Convert:  mapping,
		DownNow:  make(map[int]bool),
		DownLast: make(map[int]bool),
	}
}

func (k *Keyboard) Update() error {
	k.keys = inpututil.AppendPressedKeys(k.keys[:0])
	for i := range k.DownLast {
		k.DownLast[i] = false
	}

	for i := range k.DownNow {
		k.DownLast[i] = k.DownNow[i]
		k.DownNow[i] = false
	}

	for _, key := range k.keys {
		k.DownNow[k.Convert[key]] = true
	}

	// required for keyboard to be an entity
	return nil
}

func (k *Keyboard) Pressed(code int) bool {
	return !k.DownLast[code] && k.DownNow[code]
}

func (k *Keyboard) Down(code int) bool {
	return k.DownNow[code]
}

func (k *Keyboard) Release(code int) bool {
	return k.DownLast[code] && !k.DownNow[code]
}
