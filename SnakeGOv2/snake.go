package main

import (
	tl "github.com/JoelOtter/termloop"
)

type direction int

const (
	right direction = iota
	left
	up
	down
)

type Snake struct {
	*tl.Entity
	body      []Coord
	bodyLen   int
	direction direction
}

func NewSnake() *Snake {
	s := new(Snake)
	s.Entity = tl.NewEntity(5, 5, 1, 1)
	s.body = []Coord{
		{3, 5},
		{4, 5},
		{5, 5},
	}

	s.bodyLen = len(s.body)
	s.direction = right
	return s
}

func (s *Snake) head() *Coord {
	return &s.body[len(s.body)-1]
}

func (s *Snake) grow(amount int) {
	s.bodyLen += amount
}

func (s *Snake) isGrowing() bool {
	return s.bodyLen > len(s.body)
}

func (s *Snake) isCollidingWithSelf() bool {
	for i := 0; i < len(s.body)-1; i++ {
		if *s.head() == s.body[i] {
			return true
		}
	}

	return false
}

func (s *Snake) isCollidingWithBorder() bool {
	return border.Contains(*s.head())
}

func (s *Snake) Draw(screen *tl.Screen) {
	newHead := *s.head()
	switch s.direction {
	case right:
		newHead.x++
	case left:
		newHead.x--
	case up:
		newHead.y--
	case down:
		newHead.y++
	}

	if s.isGrowing() {
		s.body = append(s.body, newHead)
	} else {
		s.body = append(s.body[1:], newHead)
	}

	s.SetPosition(newHead.x, newHead.y)

	if s.isCollidingWithSelf() || s.isCollidingWithBorder() {
		EndGame()
	}

	for _, c := range s.body {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: 'o',
		})
	}
}

func (s *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if s.direction != left {
				s.direction = right
			}
		case tl.KeyArrowLeft:
			if s.direction != right {
				s.direction = left
			}
		case tl.KeyArrowUp:
			if s.direction != down {
				s.direction = up
			}
		case tl.KeyArrowDown:
			if s.direction != up {
				s.direction = down
			}

		case 0:
			switch event.Ch {
			case 'h', 'H':
				if s.direction != right {
					s.direction = left
				}
			case 'j', 'J':
				if s.direction != up {
					s.direction = left
				}
			case 'k', 'K':
				if s.direction != down {
					s.direction = up
				}
			case 'l', 'L':
				if s.direction != left {
					s.direction = right
				}
			}
		}
	}
}

func (s *Snake) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Food:
		s.handleFoodCollision()
	case *Border:
		s.handleBorderCollision()
	}
}

func (s *Snake) handleFoodCollision() {
	s.grow(5)
}

func (s *Snake) handleBorderCollision() {
	EndGame()
}
