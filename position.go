package prompt

import (
	"io"
	"strings"

	istrings "github.com/elk-language/go-prompt/strings"
	"github.com/mattn/go-runewidth"
)

// Position stores the coordinates
// of a p.
//
// (0, 0) represents the top-left corner of the prompt,
// while (n, n) the bottom-right corner.
type Position struct {
	X istrings.Width
	Y int
}

// Join two positions and return a new position.
func (p Position) Join(other Position) Position {
	if other.Y == 0 {
		p.X += other.X
	} else {
		p.X = other.X
		p.Y += other.Y
	}
	return p
}

// Add two positions and return a new position.
func (p Position) Add(other Position) Position {
	return Position{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

// Subtract two positions and return a new position.
func (p Position) Subtract(other Position) Position {
	return Position{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

// positionAtEndOfString calculates the position of the
// p at the end of the given string.
func positionAtEndOfStringLine(str string, columns istrings.Width, line int) Position {
	pos := positionAtEndOfReaderLine(strings.NewReader(str), columns, line)
	return pos
}

// positionAtEndOfString calculates the position
// at the end of the given string.
func positionAtEndOfString(str string, columns istrings.Width) Position {
	pos := positionAtEndOfReader(strings.NewReader(str), columns)
	return pos
}

// positionAtEndOfReader calculates the position
// at the end of the given io.Reader.
func positionAtEndOfReader(reader io.RuneReader, columns istrings.Width) Position {
	var down int
	var right istrings.Width

charLoop:
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break charLoop
		}

		switch char {
		case '\r':
			char, _, err := reader.ReadRune()
			if err != nil {
				break charLoop
			}

			if char == '\n' {
				down++
				right = 0
			}
		case '\n':
			down++
			right = 0
		default:
			right += istrings.Width(runewidth.RuneWidth(char))
			if right == columns {
				right = 0
				down++
			}
		}
	}

	return Position{
		X: right,
		Y: down,
	}
}

// positionAtEndOfReaderLine calculates the position
// at the given line of the given io.Reader.
func positionAtEndOfReaderLine(reader io.RuneReader, columns istrings.Width, line int) Position {
	var down int
	var right istrings.Width

charLoop:
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break charLoop
		}

		switch char {
		case '\r':
			char, _, err := reader.ReadRune()
			if err != nil {
				break charLoop
			}

			if char == '\n' {
				if down == line {
					break charLoop
				}
				down++
				right = 0
			}
		case '\n':
			if down == line {
				break charLoop
			}
			down++
			right = 0
		default:
			right += istrings.Width(runewidth.RuneWidth(char))
			if right > columns {
				if down == line {
					right = columns - 1
					break charLoop
				}
				right = istrings.Width(runewidth.RuneWidth(char))
				down++
			}
		}
	}

	return Position{
		X: right,
		Y: down,
	}
}
