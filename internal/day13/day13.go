package day13

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"sort"
)

type Track int

const (
	LeftUp Track = iota
	Vertical
	LeftDown
	Horizontal
	Intersection
)

func (t Track) String() string {
	switch t {
	case LeftUp:
		return "/"
	case LeftDown:
		return "\\"
	case Vertical:
		return "|"
	case Horizontal:
		return "-"
	case Intersection:
		return "+"
	}

	panic("No track representation")
}

type NextIntersectionAction int

const (
	Left     NextIntersectionAction = iota
	Straight NextIntersectionAction = iota
	Right    NextIntersectionAction = iota
)

func (n NextIntersectionAction) String() string {
	switch n {
	case Left:
		return "L"
	case Right:
		return "R"
	case Straight:
		return "S"
	}

	panic("No representation")
}

func (c Cart) newDirection(t Track) Direction {
	switch t {
	case Intersection:
		switch c.NextAction {
		case Left:
			return c.Dir.TurnLeft()
		case Straight:
			return c.Dir
		case Right:
			return c.Dir.TurnRight()
		}
	case Horizontal:
		return c.Dir
	case LeftUp:
		if c.Dir == LEFT || c.Dir == RIGHT {
			return c.Dir.TurnLeft()
		} else {
			return c.Dir.TurnRight()
		}
	case LeftDown:
		if c.Dir == LEFT || c.Dir == RIGHT {
			return c.Dir.TurnRight()
		} else {
			return c.Dir.TurnLeft()
		}
	case Vertical:
		return c.Dir
	}

	panic(fmt.Sprintf("No next dir found! %v %v", t, c))

}

type Cart struct {
	ID         string
	Pos        Coordinate
	Dir        Direction
	Crashed    bool
	NextAction NextIntersectionAction
}

func (c Cart) String() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v", c.ID, c.Pos, c.Dir, c.NextAction, c.Crashed)
}

func (c Cart) nextAction() NextIntersectionAction {
	switch c.NextAction {
	case Right:
		return Left
	case Left:
		return Straight
	case Straight:
		return Right
	}

	panic("No next action!")
}

func Solve(lines []string, partOne bool) string {

	txtMap := make(map[Coordinate]rune)
	carts := make([]*Cart, 0)

	cartId := 'A' - 1
	cartFactory := func(x int, y int, dir Direction) *Cart {
		cartId++
		return &Cart{string(cartId), Coordinate{X: x, Y: y}, dir, false, Left}
	}

	items := 0
	for y, l := range lines {
		for x, c := range l {
			r := rune(c)
			if r == ' ' {
				continue
			} else if r == '<' {
				carts = append(carts, cartFactory(x, y, LEFT))
			} else if r == '>' {
				carts = append(carts, cartFactory(x, y, RIGHT))
			} else if r == '^' {
				carts = append(carts, cartFactory(x, y, UP))
			} else if r == 'v' {
				carts = append(carts, cartFactory(x, y, DOWN))
			}
			txtMap[Coordinate{X: x, Y: y}] = r
			items++
			//fmt.Printf("(%v,%v) %v\n", x, y, string(r))
		}
	}

	trackMap := make(map[Coordinate]Track)
	for c, tr := range txtMap {
		var trackType Track = -1

		switch tr {
		case '/':
			trackType = LeftUp
		case '\\':
			trackType = LeftDown
		case '|':
			trackType = Vertical
		case '-':
			trackType = Horizontal
		case '<':
			trackType = Horizontal
		case '>':
			trackType = Horizontal
		case '^':
			trackType = Vertical
		case 'v':
			trackType = Vertical
		case '+':
			trackType = Intersection
		}

		if trackType == -1 {
			panic("Wut? No track!")
		}

		trackMap[c] = trackType

		//fmt.Printf("%v\t%v\n", c, trackType)
	}

	cartSorter := func(i, j int) bool {
		c1 := carts[i]
		c2 := carts[j]
		if c1.Pos.Y < c2.Pos.Y {
			return true
		} else if c1.Pos.Y > c2.Pos.Y {
			return false
		} else {
			return c1.Pos.X < c2.Pos.X
		}
	}

	loop := 0

	cartsAliveMap := make(map[string]*Cart)
	for _, c := range carts {
		cartsAliveMap[c.ID] = c
		fmt.Println(c)
	}

	oneTickMore := false
	for {
		if loop == 810 {
			return ""
		}
		alive := 0
		aliveC := carts[0]
		for _, c := range carts {
			if !c.Crashed {
				alive++
				aliveC = c
				//fmt.Printf("%v is alive\n", c.ID)
			}
		}

		if alive == 1 {
			fmt.Println(aliveC.ID)
			fmt.Println(aliveC.Pos)
			fmt.Println(aliveC)
			if !oneTickMore {
				oneTickMore = true
			} else {
				return fmt.Sprint(aliveC.Pos)
			}
		} else if loop%100 == 0 {
			//fmt.Printf("Alive: %v\n", alive)
			//fmt.Printf("Carts: %v\n", len(carts))
		}

		sort.Slice(carts, cartSorter)
		cartMap := make(map[Coordinate]*Cart)
		for _, c := range carts {

			if !c.Crashed {
				cartMap[c.Pos] = c
			}
			// fmt.Printf("%v: %v,%v\t", c.ID, c.Pos.X, c.Pos.Y)
		}
		// fmt.Println("")

		for i, c := range carts {
			if c.Crashed {
				continue
			}

			_, isAlive := cartsAliveMap[c.ID]

			if !isAlive {
				continue
			}

			newPos := c.Pos.Move(c.Dir)
			otherCart, found := cartMap[newPos]
			if found {
				c.Crashed = true
				otherCart.Crashed = true
				fmt.Printf("%v: Crashed %v %v %v\n", loop, newPos, c.ID, otherCart.ID)
				delete(cartMap, c.Pos)
				delete(cartMap, otherCart.Pos)
				delete(cartsAliveMap, c.ID)
				delete(cartsAliveMap, otherCart.ID)

				for _, c := range carts {
					fmt.Println(c)
				}

				if partOne {
					return fmt.Sprint(newPos)
				} else {
					continue
				}
			}

			piece, onTrack := trackMap[newPos]
			if !onTrack {
				fmt.Println(loop)
				panic("Derailed!")
			}

			newDir := c.newDirection(piece)
			c.Dir = newDir
			c.Pos = newPos
			if piece == Intersection {
				c.NextAction = c.nextAction()
			}

			cartMap[newPos] = c
			carts[i] = c
		}

		loop++
	}

	return ""
}

func printWorld(carts map[Coordinate]*Cart, tracks map[Coordinate]Track) {
	str := ""
	for y := 0; y < 151; y++ {
		for x := 0; x < 151; x++ {
			c := Coordinate{X: x, Y: y}
			cart, cf := carts[c]
			track, tf := tracks[c]
			if cf {
				str += fmt.Sprint(cart.Dir)
			} else if tf {
				str += fmt.Sprint(track)
			} else {
				str += " "
			}
		}
		str += "\n"
	}

	fmt.Println(str)
}
