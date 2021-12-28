package main

import tl "github.com/JoelOtter/termloop"

type menu struct {
	tl.BaseLevel
}

var selected int = 1
var (
	easy   *tl.Text
	medium *tl.Text
	hard   *tl.Text
)

func NewMenu(isFullscreen *bool) tl.Level {

	menuScreen := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlue,
	})
	menu := new(menu)
	menu.BaseLevel = *menuScreen

	easy = tl.NewText(45, 20, "Easy", tl.ColorWhite, tl.ColorCyan)
	medium = tl.NewText(45, 21, "Medium", tl.ColorWhite, tl.ColorCyan)
	hard = tl.NewText(45, 22, "Hard", tl.ColorWhite, tl.ColorCyan)

	menu.AddEntity(easy)
	menu.AddEntity(medium)
	menu.AddEntity(hard)

	return menu
}

func (m *menu) Tick(event tl.Event) {
	updateButtons()
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowDown:
			selectedInc()
		case tl.KeyArrowUp:
			selectedDec()
		case tl.KeyEnter:
			difficulty = selected
			StartGame()
		}
	}

}

func updateButtons() {
	switch selected {
	case 3:
		//Update
		hard.SetColor(tl.ColorWhite, tl.ColorRed)
		//Reset
		easy.SetColor(tl.ColorWhite, tl.ColorCyan)
		medium.SetColor(tl.ColorWhite, tl.ColorCyan)
	case 2:
		//Update
		medium.SetColor(tl.ColorWhite, tl.ColorRed)
		//Reset
		easy.SetColor(tl.ColorWhite, tl.ColorCyan)
		hard.SetColor(tl.ColorWhite, tl.ColorCyan)
	case 1:
		//Update
		easy.SetColor(tl.ColorWhite, tl.ColorRed)
		//Reset
		medium.SetColor(tl.ColorWhite, tl.ColorCyan)
		hard.SetColor(tl.ColorWhite, tl.ColorCyan)
	}
}

func selectedInc() {
	if selected == 3 {
		selected = 1
	} else {
		selected++
	}

}

func selectedDec() {
	if selected == 1 {
		selected = 3
	} else {
		selected--
	}
}
