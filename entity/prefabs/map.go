package prefabs

import (
	"github.com/bjatkin/chessRTS/assets/images"
	"github.com/bjatkin/chessRTS/entity"
	"github.com/bjatkin/chessRTS/entity/layer"
	"github.com/bjatkin/chessRTS/math"

	"github.com/hajimehoshi/ebiten/v2"
)

// all the posible node classes
const (
	WhiteNode = iota
	BlackNode
	WhiteResourceNode
	BlackResourceNode
)

// Node is the basic tile on the map
// 5, 1, 6
// 0, X, 2
// 4, 3, 7
type Node struct {
	Pos         math.IVec2
	AdjacentIdx [8]int
	Adjacent    [8]*Node
	Class       int
}

// Map is an RTS map where a game is played
type Map struct {
	*entity.Base

	Nodes []Node
	GFX   *ebiten.Image
}

func NewMap(nodes []Node) *Map {
	// connect all the nodes with pointers
	var width, height int
	for i, node := range nodes {
		var idx int
		// left
		if idx = node.AdjacentIdx[0]; idx >= 0 {
			nodes[i].Adjacent[0] = &nodes[idx]
		}
		// top
		if idx = node.AdjacentIdx[1]; idx >= 0 {
			nodes[i].Adjacent[1] = &nodes[idx]
		}
		// right
		if idx = node.AdjacentIdx[2]; idx >= 0 {
			nodes[i].Adjacent[2] = &nodes[idx]
		}
		// bottom
		if idx = node.AdjacentIdx[3]; idx >= 0 {
			nodes[i].Adjacent[3] = &nodes[idx]
		}
		// bottom left
		if idx = node.AdjacentIdx[4]; idx >= 0 {
			nodes[i].Adjacent[4] = &nodes[idx]
		}
		// top left
		if idx = node.AdjacentIdx[5]; idx >= 0 {
			nodes[i].Adjacent[5] = &nodes[idx]
		}
		// top right
		if idx = node.AdjacentIdx[6]; idx >= 0 {
			nodes[i].Adjacent[6] = &nodes[idx]
		}
		// bottom right
		if idx = node.AdjacentIdx[7]; idx >= 0 {
			nodes[i].Adjacent[7] = &nodes[idx]
		}

		if node.Pos.X > width {
			width = node.Pos.X
		}
		if node.Pos.Y > height {
			height = node.Pos.Y
		}
	}

	// build the map gfx to make drawing it quicker
	lightSquare := images.MustEbiten(images.LightSquare)
	darkSquare := images.MustEbiten(images.DarkSquare)
	gfx := ebiten.NewImage(20*(width+1)+300, 20*(height+1)+300)
	for _, node := range nodes {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(node.Pos.FVec2().X*20, node.Pos.FVec2().Y*20)

		switch node.Class {
		case BlackNode:
			gfx.DrawImage(darkSquare, opt)
		case WhiteNode:
			gfx.DrawImage(lightSquare, opt)
		case BlackResourceNode:
			gfx.DrawImage(darkSquare, opt)
		case WhiteResourceNode:
			gfx.DrawImage(lightSquare, opt)
		}
	}

	return &Map{
		Base: &entity.Base{
			Transform: &entity.Transform{
				Pos:   math.IVec2{X: 0, Y: 0},
				Scale: 2,
			},
			GFX:       gfx,
			DrawLayer: layer.BG,
		},

		Nodes: nodes,
	}
}
