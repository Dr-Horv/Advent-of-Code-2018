package day13

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"sort"
)

type Action int

const (
	Left     Action = 1
	Straight        = 2
	Right           = 3
)

type CartDirection Coordinate

type Cart struct {
	ID               string
	Pos              Coordinate
	Crashed          bool
	NextIntersection Action
	Dir              CartDirection
}

var CartLeft = CartDirection(Coordinate{X: -1})
var CartRight = CartDirection(Coordinate{X: 1})
var CartUp = CartDirection(Coordinate{Y: -1})
var CartDown = CartDirection(Coordinate{Y: 1})

func (c *Cart) turnLeft() {
	switch c.Dir {
	case CartLeft:
		c.Dir = CartDown
	case CartDown:
		c.Dir = CartRight
	case CartRight:
		c.Dir = CartUp
	case CartUp:
		c.Dir = CartLeft
	}
}

func (c *Cart) turnRight() {
	switch c.Dir {
	case CartLeft:
		c.Dir = CartUp
	case CartUp:
		c.Dir = CartRight
	case CartRight:
		c.Dir = CartDown
	case CartDown:
		c.Dir = CartLeft
	}
}

func Solve(lines []string, partOne bool) string {
	tracks := make(map[Coordinate]rune)
	id := 'A' - 1
	cartMap := make(map[Coordinate]*Cart, 0)
	cartFactory := func(x int, y int, cd CartDirection) *Cart {
		id++
		return &Cart{string(id), Coordinate{X: x, Y: y}, false, Left, cd}
	}

	for y, l := range lines {
		for x, r := range l {
			c := Coordinate{X: x, Y: y}
			if r == '^' {
				tracks[c] = '|'
				cartMap[c] = cartFactory(x, y, CartUp)
			} else if r == '>' {
				tracks[c] = '-'
				cartMap[c] = cartFactory(x, y, CartRight)
			} else if r == '<' {
				tracks[c] = '-'
				cartMap[c] = cartFactory(x, y, CartLeft)
			} else if r == 'v' {
				tracks[c] = '|'
				cartMap[c] = cartFactory(x, y, CartDown)
			} else {
				tracks[c] = r
			}
		}
	}

	tick := 1
	for {
		cartsAlive := len(cartMap)
		if cartsAlive == 1 {
			for _, c := range cartMap {
				return fmt.Sprint(c.Pos)
			}
		}

		carts := make([]*Cart, 0)
		for _, c := range cartMap {
			carts = append(carts, c)
		}

		sort.Slice(carts, func(i, j int) bool {
			c1 := carts[i]
			c2 := carts[j]

			if c1.Pos.Y < c2.Pos.Y {
				return true
			} else if c1.Pos.Y == c2.Pos.Y && c1.Pos.X < c2.Pos.X {
				return true
			}

			return false
		})

		for _, c := range carts {
			if c.Crashed {
				continue
			}

			newPos := c.Pos.Plus(Coordinate(c.Dir))
			otherCart, hasCart := cartMap[newPos]

			if hasCart {
				c.Crashed = true
				otherCart.Crashed = true
				delete(cartMap, otherCart.Pos)
				fmt.Printf("tick: %v, Crash at %v between %v %v\n", tick, newPos, c.ID, otherCart.ID)
				if partOne {
					return fmt.Sprint(newPos)
				}
			}

			delete(cartMap, c.Pos)

			if c.Crashed {
				continue
			}

			c.Pos = newPos
			trackPiece, trackFound := tracks[newPos]

			if !trackFound {
				fmt.Printf("Derailed %v at %v", c.ID, c.Pos)
				return ""
			}

			switch trackPiece {
			case '+':
				switch c.NextIntersection {
				case Left:
					c.turnLeft()
					c.NextIntersection = Straight
				case Straight:
					c.NextIntersection = Right
				case Right:
					c.turnRight()
					c.NextIntersection = Left
				}
			case '/':
				switch c.Dir {
				case CartDown:
					c.turnRight()
				case CartRight:
					c.turnLeft()
				case CartLeft:
					c.turnLeft()
				case CartUp:
					c.turnRight()
				}
			case '\\':
				switch c.Dir {
				case CartDown:
					c.turnLeft()
				case CartRight:
					c.turnRight()
				case CartLeft:
					c.turnRight()
				case CartUp:
					c.turnLeft()
				}
			}

			cartMap[c.Pos] = c

		}
		tick++
	}
}
