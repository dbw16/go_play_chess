package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const Red = "\033[1;31m%s\033[0m"
const White = "\033[1;37m%s\033[0m"

const (
	screenWidth = 480
	screenHeight = 480
	imageboardSize = 400
	boardSize = 480
	peiceSize = 60
)


var (
	black Colour = Colour{colour: 1}
	white Colour = Colour{colour: 0}
)

type Colour struct {
	colour int
}

type Coord struct {
	x int
	y int
}

type Move struct {
	src Coord
	dst Coord
}

func min_max(a, b int) (min int, max int) {
	if a < b {
		return a, b
	}
	return b, a
}

func (move Move) ValidDigaonalMove() bool{
	var xMoves []int
	var yMoves []int
	if move.src.x == move.dst.x {
		return false
	}
	if move.src.y == move.dst.y {
		return false
	}

	for x := move.src.x; x != move.dst.x ; {
		xMoves = append(xMoves, x)
		if move.dst.x < move.src.x{
			x--
		}else {
			x++
		}
	}

	for y := move.src.y; y != move.dst.y;  {
		yMoves = append(yMoves, y)
		if move.dst.y < move.src.y{
			y--
		}else {
			y++
		}

	}

	if len(xMoves) != len(yMoves){
		return false
	}

	return true
}

func (move Move) Journey() []Coord{
	var coords []Coord
	//first case stight line in y direction
	if move.src.x == move.dst.x{
		min, max := min_max(move.src.y, move.dst.y)
		for y := min +1; y < max; y++{
			coords = append(coords, Coord{x: move.src.x, y: y})
		}
	}
	//case stight line in x direction
	if move.src.y == move.dst.y{
		min, max := min_max(move.src.x, move.dst.x)
		for x := min +1; x < max; x++{
			coords = append(coords, Coord{x: x, y: move.src.y})
		}
	}else {

		var xMoves []int
		var yMoves []int

		for x := move.src.x; x != move.dst.x ; {
			xMoves = append(xMoves, x)
			if move.dst.x < move.src.x{
				x--
			}else {
				x++
			}
		}

		for y := move.src.y; y != move.dst.y;  {
			yMoves = append(yMoves, y)
			if move.dst.y < move.src.y{
				y--
			}else {
				y++
			}

		}
		if len(xMoves) != len(yMoves){
			return coords
		}
		for i :=1; i< len(xMoves); i++{
			coords = append(coords, Coord{x: xMoves[i], y: yMoves[i]} )
		}
	}

	return coords

}

type Piece interface {
	Colour() Colour
	Rep()
	ValidMove(Move, ChessBoard) bool
	Tex() *sdl.Texture
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture{
	if renderer == nil{
		return nil
	}
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}
	return tex
}

func mouseClickPos() Coord{
	x, y, _ := sdl.GetMouseState()
	return Coord{int(x /peiceSize), int(y/peiceSize) }
}

func main()  {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	myBoard := newBoard(renderer)
	myBoard.ShowBoard()
	//myBoard.ShowBoard()
	//myBoard.move(Coord{7,7},Coord{6,6})
	//myBoard.ShowBoard()

	move := Move{src: Coord{0,0}, dst: Coord{0,4}}
	move2 := Move{src: Coord{7,6}, dst: Coord{7,4}}
	move3 := Move{src: Coord{6,6}, dst: Coord{0,3}}

	myBoard.move(move)
	myBoard.move(move2)
	myBoard.move(move3)

	myBoard.ShowBoard()
	var clickedSquare, secondClickedSquare Coord
	renderer.SetDrawColor(255, 255, 255, 255)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.MouseButtonEvent:
				clickedSquare = mouseClickPos()
				secondClickedSquare = mouseClickPos()
				if myBoard.validMove(Move{src: clickedSquare, dst: secondClickedSquare}){
					println("a second mouse event happened?" )
					secondClickedSquare = mouseClickPos()
				} else{
					secondClickedSquare = Coord{0,0}
				}

				println("a mouse event happened?" )
			}
		}

		renderer.Clear()
		myBoard.draw(renderer)
		myBoard.drawPeices(renderer)
		myBoard.drawclickedSquare(renderer, clickedSquare)
		myBoard.drawclickedSquare(renderer, secondClickedSquare)

		myBoard.drawPossibleMoves(renderer, clickedSquare)
		renderer.Present()
	}






}