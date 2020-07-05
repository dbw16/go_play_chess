package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

//*******Rook struct and methods
type Rook struct {
	colour Colour
	rep string
	tex *sdl.Texture

}

func NewRook(c Colour, renderer *sdl.Renderer) Piece {
	if c == black{
		return &Rook{colour: c, rep: "B",tex: textureFromBMP(renderer, "images/br.bmp")}
	}
	return &Rook{colour: c, rep: "B",tex: textureFromBMP(renderer, "images/wr.bmp")}
}

func (r Rook) Colour() Colour  {
	return r.colour
}

func (r Rook) ValidMove(move Move, board ChessBoard) bool {
	return true
}

func (r Rook) Rep()  {
	if r.colour == black{
		fmt.Printf(Red, "R")
	} else{
		fmt.Printf(White, "R")
	}
}

func (p Rook) Tex() *sdl.Texture  {
	return p.tex
}