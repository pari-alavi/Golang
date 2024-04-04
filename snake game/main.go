package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

const SnakeSymbol = 0x2588
const AppleSymbol = 0x25CF
const GameFrameWidth = 40
const GameFrameHeight = 20
const GameFrameSymbol = '║'

type Point struct {
	row, col int
}

type Snake struct {
	parts          []*Point
	velRow, velCol int
	symbol         rune
}

type Apple struct {
	point  *Point
	symbol rune
}

var screen tcell.Screen
var GamePaused bool
var GameOver bool
var debugLog string
var score int
var PointsToClear []*Point
var snake *Snake
var apple *Apple

func main() {

	// Create a new random number generator with a custom seed (e.g., current time)
	rand.Seed(time.Now().UnixNano())

	InitScreen()
	InitGameState()
	inputChan := InitUserInput()

	for !GameOver {

		UpdateState()
		DrawState()
		HandleUserInput(ReadInput(inputChan))

		time.Sleep(200 * time.Millisecond)
	}
	screenWidth, screenHeight := screen.Size()
	PrintStringCentered(screenHeight/2, screenWidth/2, "Game Over")
	PrintStringCentered(screenHeight/2+1, screenWidth/2, fmt.Sprintf("your score is %d...", score))
	screen.Show()
	time.Sleep(3 * time.Second)
	screen.Fini()

}

func DrawState() {

	// if GamePaused {
	// 	return
	// }

	//screen.Clear()
	ClearScreen()

	PrintString(0, 0, debugLog)
	PrintGameFrame()

	DrawSnake()
	DrawApple()

	screen.Show()
}

func ClearScreen() {
	for _, p := range PointsToClear {
		PrintFilledRectInGameFrame(p.row, p.col, 1, 1, ' ')
	}

	PointsToClear = []*Point{}
}

func InitScreen() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
}

func InitGameState() {

	snake = &Snake{
		parts: []*Point{
			{row: 9, col: 3}, //tail
			{row: 8, col: 3},
			{row: 7, col: 3},
			{row: 6, col: 3},
			{row: 5, col: 3}, //head
		},
		velRow: -1,
		velCol: 0,
		symbol: SnakeSymbol,
	}
	apple = &Apple{
		point:  &Point{row: 10, col: 10},
		symbol: AppleSymbol,
	}
}

func HandleUserInput(key string) {
	if key == "Rune[q]" {
		screen.Fini()
		os.Exit(0)
	} else if key == "Rune[p]" {
		GamePaused = !GamePaused
	} else if key == "Up" && snake.velRow != 1 {
		snake.velRow = -1
		snake.velCol = 0
	} else if key == "Down" && snake.velRow != -1 {
		snake.velRow = 1
		snake.velCol = 0
	} else if key == "Right" && snake.velCol != -1 {
		snake.velRow = 0
		snake.velCol = 1
	} else if key == "Left" && snake.velCol != 1 {
		snake.velRow = 0
		snake.velCol = -1
	}

}

func InitUserInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventKey:
				inputChan <- ev.Name()
			}
		}
	}()
	return inputChan
}

func ReadInput(inputChan chan string) string {
	var key string
	select {
	case key = <-inputChan:
	default:
		key = ""
	}
	return key
}

func PrintGameFrame() {
	screenWidth, screenHeight := screen.Size()
	row, col := screenHeight/2-GameFrameHeight/2-1, screenWidth/2-GameFrameWidth/2-1
	width, height := GameFrameWidth+2, GameFrameHeight+2
	PrintUnfilledRect(row, col, width, height, GameFrameSymbol)

}

func DrawSnake() {
	for _, p := range snake.parts {
		PrintFilledRectInGameFrame(p.row, p.col, 1, 1, SnakeSymbol)
		PointsToClear = append(PointsToClear, p)
	}
}

func DrawApple() {
	PrintFilledRectInGameFrame(apple.point.row, apple.point.col, 1, 1, AppleSymbol)
	PointsToClear = append(PointsToClear, apple.point)
}

func UpdateState() {
	if GamePaused {
		return
	}

	//update snake+Apple
	UpdateSnake()
	UpdateApple()

}

func GetSnakeHead() *Point {

	return snake.parts[len(snake.parts)-1]
}

func UpdateSnake() {
	//Add new element
	Head := GetSnakeHead()
	snake.parts = append(snake.parts, &Point{
		row: Head.row + snake.velRow,
		col: Head.col + snake.velCol,
	})
	//Delete last element
	if !AppleInsideSnake() {
		snake.parts = snake.parts[1:]
	} else {
		score++
	}
	if SnakeHittingWall() || SnakeEatItself() {
		GameOver = true
	}

}

func SnakeHittingWall() bool {
	Head := GetSnakeHead()
	return Head.row < 0 ||
		Head.row >= GameFrameHeight ||
		Head.col < 0 ||
		Head.col >= GameFrameWidth
}

func SnakeEatItself() bool {
	head := GetSnakeHead() //[0:snakeHeadIndex)
	for _, p := range snake.parts[:len(snake.parts)-1] {
		if p.row == head.row && p.col == head.col {
			return true
		}
	}
	return false
}

func PrintString(row, col int, str string) {
	for _, c := range str {
		PrintFilledRect(row, col, 1, 1, c)
		col += 1
	}

}

func PrintStringCentered(row, col int, str string) {
	col = col - len(str)/2
	PrintString(row, col, str)

}

func PrintFilledRectInGameFrame(row, col, width, height int, ch rune) {
	//get game frame top lef point(row, col)
	screenWidth, screenHeight := screen.Size()
	r, c := screenHeight/2-GameFrameHeight/2, screenWidth/2-GameFrameWidth/2

	PrintFilledRect(row+r, col+c, width, height, ch)

}

func PrintFilledRect(row, col, width, height int, ch rune) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			screen.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}
	}
}

func PrintUnfilledRect(row, col, width, height int, ch rune) {
	// print first row
	for c := 0; c < width; c++ {
		screen.SetContent(col+c, row, ch, nil, tcell.StyleDefault)
	}
	//for each row
	//    print first coll and last coll
	for r := 1; r < height-1; r++ {
		screen.SetContent(col, row+r, ch, nil, tcell.StyleDefault)
		screen.SetContent(col+width-1, row+r, ch, nil, tcell.StyleDefault)

	}
	//print laste row
	for c := 0; c < width; c++ {
		screen.SetContent(col+c, row+height-1, ch, nil, tcell.StyleDefault)
	}

}

func UpdateApple() {
	//do this while apple is inside snake
	for AppleInsideSnake() {

		apple.point.row, apple.point.col = rand.Intn(GameFrameHeight), rand.Intn(GameFrameWidth)
	}

}

func AppleInsideSnake() bool {
	for _, p := range snake.parts {
		if p.row == apple.point.row && p.col == apple.point.col {
			return true
		}
	}
	return false
}
