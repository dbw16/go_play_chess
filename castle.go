package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

//*******Castle struct and methods
type Castle struct {
	colour Colour
	rep string
	tex *sdl.Texture

}

func NewCastle(c Colour, renderer *sdl.Renderer) Piece {
	if c == black{
		return &Castle{colour: c, rep: "B",tex: textureFromBMP(renderer, "images/bc.bmp")}
	}
	return &Castle{colour: c, rep: "B",tex: textureFromBMP(renderer, "images/wc.bmp")}
}

func (c Castle) Colour() Colour  {
	return c.colour
}

func (c Castle) ValidMove(move Move, board ChessBoard) bool {
	//Castles can move in a infite stright line in all directions
	if move.src.x != move.dst.x && move.src.y != move.dst.y{
		return false
	}

	if board.IsFriendlyPieceHere(move) {
		return false
	}

	//stops it from ignoring peaices
	for _, coord := range move.Journey(){
		if board.IsPieceHere(coord) == true{
			return false
		}
	}



	return true
}

func (c Castle) Rep()  {
	if c.colour == black{
		fmt.Printf(Red, "C")
	} else{
		fmt.Printf(White, "C")
	}
}

func (p Castle) Tex() *sdl.Texture  {
	return p.tex
}