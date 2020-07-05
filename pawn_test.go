package main

import "testing"

func TestPawn_Colour(t *testing.T) {
	var (
		tables = []struct {
			pawn   Pawn
			colour Colour
		}{
			{Pawn{colour: black}, black},
			{Pawn{colour: white}, white},
		}
	)
	for _, table := range tables {
		if table.pawn.Colour() != table.colour {
			t.Errorf("Pawm colour was incorrect, got: %d, want: %d.", table.pawn.Colour(), table.colour)
		}
	}

}

func TestPawn_ValidMove(t *testing.T) {
	var (
		tables = []struct {
			pawn   Pawn
			move   Move
			result bool
		}{
			{Pawn{colour: white}, Move{src: Coord{x: 1, y: 1}, dst: Coord{x: 2, y: 1}}, false},
			{Pawn{colour: white}, Move{src: Coord{x: 1, y: 1}, dst: Coord{x: 1, y: 2}}, true},
			{Pawn{colour: white}, Move{src: Coord{x: 1, y: 1}, dst: Coord{x: 1, y: 3}}, false},
			{Pawn{colour: black}, Move{src: Coord{x: 1, y: 1}, dst: Coord{x: 1, y: 0}}, true},
			{Pawn{colour: black}, Move{src: Coord{x: 1, y: 1}, dst: Coord{x: 1, y: 3}}, false}}
	)

	for _, table := range tables {
		var cb ChessBoard
		cb.square[table.move.src.y][table.move.src.x] = table.pawn
		result := table.pawn.ValidMove(table.move, cb)
		if result != table.result{
			t.Errorf("InvalidMove() was incorrect,using %v got: %t, want: %t.",table, result, table.result)
		}
	}
}

func TestNewPawn(t *testing.T) {
	pawn:= Pawn{colour: black, tex: nil}

	newPawn := NewPawn(black,nil)
	 if pawn.colour != newPawn.Colour(){
		 t.Errorf("InvalidMove() was incorrect")
	 }

	pawn2:= Pawn{colour: white, tex: nil}

	newPawn2 := NewPawn(white,nil)
	if pawn2.colour == newPawn2.Colour(){
		t.Errorf("InvalidMove() was incorrect")
	}
}
