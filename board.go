package main

import (
	"C"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

//*******ChessBoard struct and methods
type ChessBoard struct {
	square  [8][8]Piece
	tex *sdl.Texture
}

func newBoard(renderer *sdl.Renderer)  *ChessBoard{
	var board ChessBoard

	board.tex = textureFromBMP(renderer, "images/chessboard.bmp")

	for i, _ := range board.square{
		board.square[6][i] = NewPawn(black, renderer)
		board.square[1][i] = NewPawn(white, renderer)
	}

	//we can now set the other rows
	row := 0
	for _, colour := range [2]Colour{white, black}{
		board.square[row][0] = NewCastle(colour, renderer)
		board.square[row][1] = NewRook(colour, renderer)
		board.square[row][2] = NewBishop(colour, renderer)
		board.square[row][3] = NewQueen(colour, renderer)
		board.square[row][4] = NewKing(colour, renderer)
		board.square[row][5] = NewBishop(colour, renderer)
		board.square[row][6] = NewRook(colour, renderer)
		board.square[row][7] = NewCastle(colour, renderer)
		row = 7
	}

	return &board
}

func (cb ChessBoard) ShowBoard()   {
	for i:= len(cb.square) -1; i >= 0; i-- {
		for j, _ := range cb.square{
			if cb.square[i][j] != nil {
				cb.square[i][j].Rep()
			} else{
				fmt.Print("_")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (cb *ChessBoard) move(move Move)   {
	cb.square[move.dst.y][move.dst.x] = cb.square[move.src.y][move.src.x]
	cb.square[move.src.y][move.src.x] = nil
}

func (cb ChessBoard) validMove(move Move) bool{
	if move.src.x > 7 || move.src.y >7 {
		return false
	}

	return 	cb.square[move.src.y][move.src.x].ValidMove(move, cb)

}

func (cb ChessBoard) IsPieceHere(coord Coord) bool{
	if cb.square[coord.y][coord.x] != nil{
		return true
	}
	return false
}

func (cb ChessBoard) IsFriendlyPieceHere(move Move) bool{
	if cb.square[move.dst.y][move.dst.x] != nil{
		if cb.square[move.src.y][move.src.x].Colour() == cb.square[move.dst.y][move.dst.x].Colour() {
			return true
		}
	}
	return false
}

func (cb ChessBoard) ValidMoves(src_coord Coord) []Coord{
	var validMoves []Coord
	if cb.IsPieceHere(src_coord) == false{
		return nil
	}
	for i := 0; i < 8; i++ {
		for j :=0;  j <8;  j++ {
			if cb.validMove(Move{src_coord, Coord{i,j}}) == true{
				validMoves = append(validMoves, Coord{i,j})
			}
		}
	}
	return validMoves
}

func (cb *ChessBoard) draw(renderer *sdl.Renderer ){
	//convert coord system to middle of sprite
	renderer.Copy(cb.tex,
		&sdl.Rect{X: 0, Y: 0, W: imageboardSize, H: imageboardSize},
		&sdl.Rect{X: 0, Y: 0, W: boardSize, H: boardSize})
}

func (cb *ChessBoard)  drawPeices(renderer *sdl.Renderer ) {
	for i := len(cb.square) - 1; i >= 0; i-- {
		for j, _ := range cb.square {
			if cb.square[i][j] != nil {
				renderer.Copy(cb.square[i][j].Tex(),
					&sdl.Rect{X: 0, Y: 0, W: peiceSize, H: peiceSize},
					&sdl.Rect{X: int32(j * peiceSize), Y: int32(i * peiceSize), W: peiceSize, H: peiceSize})
			}
		}
	}
}

func (cb *ChessBoard) drawclickedSquare(renderer *sdl.Renderer, coord Coord)  {
	greenTex := textureFromBMP(renderer, "images/green_square.bmp")
	renderer.Copy(greenTex,
		&sdl.Rect{X: 0, Y: 0, W: peiceSize, H: peiceSize},
		&sdl.Rect{X: int32(coord.x * peiceSize), Y: int32(coord.y * peiceSize), W: peiceSize, H: peiceSize})

}

func (cb *ChessBoard) drawPossibleMoves(renderer *sdl.Renderer, coord Coord)  {

	for _, sqaure := range cb.ValidMoves(coord){
		greenTex := textureFromBMP(renderer, "images/green_square_thin.bmp")
		renderer.Copy(greenTex,
			&sdl.Rect{X: 0, Y: 0, W: peiceSize, H: peiceSize},
			&sdl.Rect{X: int32(sqaure.x * peiceSize), Y: int32(sqaure.y * peiceSize), W: peiceSize, H: peiceSize})

	}

}

