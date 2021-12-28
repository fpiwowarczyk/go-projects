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
	difficulty   int
	fps          float64
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
	scoreText = tl.NewText(0, 0, " Score: 0", tl.ColorBlack, tl.ColorWhite)

	mainLevel.AddEntity(border)
	mainLevel.AddEntity(snake)
	mainLevel.AddEntity(food)
	mainLevel.AddEntity(scoreText)
	return mainLevel
}

func setSimulationSpeed() {
	switch difficulty {
	case 1:
		fps = 10.0
	case 2:
		fps = 20.0
	case 3:
		fps = 30.0
	}
}

func init() {
	isFullscreen = flag.Bool("fullscreen", false, "Play fullscreen!")
	flag.Parse()
}

//Function start main menu
func main() {
	rand.Seed(time.Now().UnixNano())
	game = tl.NewGame()
	setSimulationSpeed()
	menu := NewMenu(isFullscreen)
	game.Screen().SetLevel(menu)
	game.Screen().SetFps(fps)
	game.Start()
}

func StartGame() {
	SetDifficulty()
	mainLevel := newMainLevel(isFullscreen)
	game.Screen().SetLevel(mainLevel)
}

func SetDifficulty() {
	difficulty = selected
	setSimulationSpeed()
	game.Screen().SetFps(fps)
}
