package controller

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// local key codes
const (
	W = iota
	A
	S
	D
)

type Mapping struct {
	Local int
	Key   ebiten.Key
}

type Keyboard struct {
	keys     []ebiten.Key
	DownLast map[int]bool
	DownNow  map[int]bool
	Convert  []Mapping
}

func NewKeyboard(mapping []Mapping) *Keyboard {
	return &Keyboard{
		Convert:  mapping,
		DownNow:  make(map[int]bool),
		DownLast: make(map[int]bool),
	}
}

func (k *Keyboard) Update() {
	k.keys = inpututil.AppendPressedKeys(k.keys[:0])
	for i := range k.DownLast {
		k.DownLast[i] = false
	}

	for i := range k.DownNow {
		k.DownLast[i] = k.DownNow[i]
		k.DownNow[i] = false
	}

	for _, key := range k.keys {
		for _, c := range k.Convert {
			if key == c.Key {
				k.DownNow[c.Local] = true
				break
			}
		}
	}
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
