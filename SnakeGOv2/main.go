package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
	"github.com/nsf/termbox-go"
)

var (
	score        int
	game         *tl.Game
	border       *Border
	scoreText    *tl.Text
	isFullscreen *bool
)

type endgameScreen struct {
	*tl.BaseLevel
}

func (eg *endgameScreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			score = 0
			game.Screen().SetLevel(newMainLevel(isFullscreen))
		}
	}
}

func IncreaseScore(amount int) {
	score += amount
	scoreText.SetText(fmt.Sprint(" Score: ", score, " "))
}

func EndGame() {
	endLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorRed,
	})
	el := new(endgameScreen)
	el.BaseLevel = endLevel
	var PromptQuestion, PromptText *tl.Text
	PromptQuestion = tl.NewText(34, 17, " Play Again? ", tl.ColorBlue, tl.ColorWhite)
	PromptText = tl.NewText(34, 18, " Press Enter", tl.ColorBlue, tl.ColorWhite)
	scoreText.SetPosition(35, 14)
	scoreText.SetColor(tl.ColorBlue, tl.ColorWhite)
	el.AddEntity(scoreText)
	el.AddEntity(PromptQuestion)
	el.AddEntity(PromptText)

	game.Screen().SetLevel(el)
}

func newMainLevel(isFullscreen *bool) tl.Level {

	mainLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	width, height := 80, 30
	if *isFullscreen {
		termbox.Init()
		width, height = termbox.Size()
	}
	border = NewBorder(width, height)

	snake := NewSnake()
	food := NewFood()
	scoreText = tl.NewText(0, 0, " Score: 0", tl.ColorBlack, tl.ColorBlack)

	mainLevel.AddEntity(border)
	mainLevel.AddEntity(snake)
	mainLevel.AddEntity(food)
	mainLevel.AddEntity(scoreText)
	return mainLevel
}

func main() {
	isFullscreen = flag.Bool("fullscreen", false, "Play fullscreen!")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	game = tl.NewGame()

	mainLevel := newMainLevel(isFullscreen)

	game.Screen().SetLevel(mainLevel)
	game.Screen().SetFps(30)
	game.Start()
}
