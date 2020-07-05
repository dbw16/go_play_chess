package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

//*******qnuee struct and methods
type Queen struct {
	colour Colour
	rep string
	tex *sdl.Texture

}

func NewQueen(c Colour, renderer *sdl.Renderer) Piece {
	if c == black{
		return &Queen{colour: c, rep: "B",tex: textureFromBMP(renderer, "images/bq.bmp")}
	}
	return &Queen{colour: c, rep: "B",tex: textureFromBMP(renderer, "images/wq.bmp")}
}

func (q Queen) Colour() Colour  {
	return q.colour
}

func (q Queen) Rep()  {
	if q.colour == black{
		fmt.Printf(Red, "Q")
	} else{
		fmt.Printf(White, "Q")
	}
}

func (q Queen) ValidMove(move Move, board ChessBoard) bool {
	return true
}

func (p Queen) Tex() *sdl.Texture  {
	return p.tex
}