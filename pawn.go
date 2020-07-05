package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

//*******Pawn struct and methods
type Pawn struct {
	colour Colour
	tex *sdl.Texture
}

func NewPawn(c Colour, renderer *sdl.Renderer) Piece {
	if c == black{
		return &Pawn{colour: c, tex: textureFromBMP(renderer, "images/bp.bmp")}
	}
	return &Pawn{colour: c, tex: textureFromBMP(renderer, "images/wp.bmp")}
}

func (p Pawn) Colour() Colour  {
	return p.colour
}

func (p Pawn) Rep()  {
	if p.colour == black{
		fmt.Printf(Red, "P")
	} else{
		fmt.Printf(White, "P")
	}
}

func (p Pawn) ValidMove(move Move, board ChessBoard) bool {
	// Pawns cant move sideways
	if move.src.x != move.dst.x{
		return false
	}

	//Black pawns can move down one square
	//White pawns can move up one square
	if p.colour == black{
		if move.src.y -1 != move.dst.y{
			return false
		}
	} else
	if move.src.y +1 != move.dst.y {
		return false
	}

	return true
}

func (p Pawn) Tex() *sdl.Texture  {
	return p.tex
}