package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

//*******King struct and methods
type King struct {
	colour Colour
	rep string
	tex *sdl.Texture

}

func NewKing(c Colour, renderer *sdl.Renderer) Piece {
	if c == black{
		return &King{colour: c, rep: "B",tex: textureFromBMP(renderer, "images/bk.bmp")}
	}
	return &King{colour: c, rep: "B",tex: textureFromBMP(renderer, "images/wk.bmp")}
}

func (k King) Colour() Colour  {
	return k.colour
}

func (k King) Rep()  {
	if k.colour == black{
		fmt.Printf(Red, "K")
	} else{
		fmt.Printf(White, "K")
	}
}

func (k King) ValidMove(move Move, board ChessBoard) bool {
	return true
}

func (p King) Tex() *sdl.Texture  {
	return p.tex
}
