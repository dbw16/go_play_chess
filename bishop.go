package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

//*******Bishop struct and methods
type Bishop struct {
	colour Colour
	tex *sdl.Texture
}

func NewBishop(c Colour, renderer *sdl.Renderer ) Piece {
	if c == black{
		return &Bishop{colour: c, tex: textureFromBMP(renderer, "images/bb.bmp")}
	}
	return &Bishop{colour: c, tex: textureFromBMP(renderer, "images/wb.bmp")}
}

func (b Bishop) Colour() Colour  {
	return b.colour
}

func (b Bishop) ValidMove(move Move, board ChessBoard) bool {
	return true
}

func (b Bishop) Rep()  {
	if b.colour == black{
		fmt.Printf(Red, "B")
	} else{
		fmt.Printf(White, "B")
	}
}

func (p Bishop) Tex() *sdl.Texture  {
	return p.tex
}