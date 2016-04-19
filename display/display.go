package display

import "github.com/nsf/termbox-go"

const (
	BigCharWidth  = 3
	BigCharHeight = 5
)

var bigChars = map[rune][BigCharHeight]string{
	0: {
		"X X",
		" X ",
		"X X",
		" X ",
		"X X",
	},
	' ': {
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	},
	'0': {
		"XXX",
		"X X",
		"X X",
		"X X",
		"XXX",
	},
	'1': {
		" X ",
		"XX ",
		" X ",
		" X ",
		"XXX",
	},
	'2': {
		"XXX",
		"  X",
		"XXX",
		"X  ",
		"XXX",
	},
	'3': {
		"XXX",
		"  X",
		"XXX",
		"  X",
		"XXX",
	},
	'4': {
		"X X",
		"X X",
		"XXX",
		"  X",
		"  X",
	},
	'5': {
		"XXX",
		"X  ",
		"XXX",
		"  X",
		"XXX",
	},
	'6': {
		"XXX",
		"X  ",
		"XXX",
		"X X",
		"XXX",
	},
	'7': {
		"XXX",
		"  X",
		"  X",
		" X ",
		" X ",
	},
	'8': {
		"XXX",
		"X X",
		"XXX",
		"X X",
		"XXX",
	},
	'9': {
		"XXX",
		"X X",
		"XXX",
		"  X",
		"  X",
	},
	':': {
		"   ",
		" X ",
		"   ",
		" X ",
		"   ",
	},
	'.': {
		"   ",
		"   ",
		"   ",
		" XX",
		" XX",
	},
}

type Point struct {
	X, Y   int
	Fg, Bg termbox.Attribute
}

func (p Point) Char(ch rune) {
	termbox.SetCell(p.X, p.Y, ch, p.Fg, p.Bg)
}

func (p Point) Str(s string) {
	for _, c := range s {
		p.Char(c)
		p.X++
	}
}

func (p Point) Pattern(pattern [BigCharHeight]string) {
	for y, line := range pattern {
		for x, c := range line {
			q := p
			q.X += x
			q.Y += y
			if c != ' ' {
				q.Fg, q.Bg = p.Bg, p.Fg
			}
			q.Char(' ')
		}
	}
}

func (p Point) BigChar(ch rune) {
	pattern, ok := bigChars[ch]
	if !ok {
		pattern = bigChars[0]
	}
	p.Pattern(pattern)
}

func (p Point) BigStr(s string) {
	xOffset := p.X
	for i, c := range s {
		p.X = xOffset + i*(BigCharWidth+1)
		p.BigChar(c)
	}
}

func (p Point) ProgressBar(length, cur, total int) {
	divider := (length * cur) / total
	for x := 0; x < length; x++ {
		ch := ' '
		q := p
		q.X += x
		if x == divider {
			ch = '░'
		}
		if x < divider {
			q.Fg, q.Bg = p.Bg, p.Fg
		}
		q.Char(ch)
	}
}
