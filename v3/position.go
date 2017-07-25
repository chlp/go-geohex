package geohex

import (
	"fmt"
	"math"
	"strconv"
)

// Position implements a grid tile position
type Position struct {
	X, Y  int
	Level uint8
}

// Decode decodes a string code into a Position.
func Decode(code string) (Position, error) {
	lnc := len(code)
	pos := Position{Level: uint8(lnc - 2)}

	if pos.Level < 0 || pos.Level > MaxLevel {
		return pos, ErrLevelInvalid
	}

	var n1, n2 int
	var ok bool

	if n1, ok = hIndex[code[0]]; !ok {
		return pos, ErrCodeInvalid
	} else if n2, ok = hIndex[code[1]]; !ok {
		return pos, ErrCodeInvalid
	}

	base := n1*30 + n2
	if base < 100 {
		code = "0" + strconv.Itoa(base) + code[2:]
	} else {
		code = strconv.Itoa(base) + code[2:]
	}

	for i, digit := range code {
		n := uint8(digit - '0')
		if n < 0 || n > 9 {
			return pos, fmt.Errorf("expected a digit, got %q", string(digit))
		}

		pow := pow3[lnc-i]
		switch n / 3 {
		case 0:
			pos.X -= pow
		case 2:
			pos.X += pow
		}
		switch n % 3 {
		case 0:
			pos.Y -= pow
		case 2:
			pos.Y += pow
		}
	}

	// normalise/adjust X	and Y
	dA := pos.X - pos.Y
	if dA < 0 {
		dA = -dA
	}
	if dM := pow3[lnc]; dM == dA && pos.X > pos.Y {
		pos.X, pos.Y = pos.Y, pos.X
	} else if dS := dA - dM; dS > 0 {
		dX := dS / 2
		dY := dS - dX
		if pos.X > pos.Y {
			pos.X, pos.Y = pos.Y+dX+dY, pos.X-dX-dY
		} else if pos.Y > pos.X {
			pos.X, pos.Y = pos.Y-dY-dX, pos.X+dX+dY
		}
	}

	return pos, nil
}

// Encode encodes a lat/lon/level into a Position
func Encode(lat, lon float64, level uint8) (Position, error) {
	return NewLL(lat, lon).Position(level)
}

// Centroid returns the centroid point of the tile
func (p Position) Centroid() Point {
	z := zooms[p.Level]
	x := float64(p.X)
	y := float64(p.Y)
	n := (hK*x*z.w + y*z.h) / 2
	e := (n - y*z.h) / hK
	return Point{E: e, N: n}
}

// LL converts the position into a LL
func (p Position) LL() LL {
	c := p.Centroid()
	z := zooms[p.Level]

	exp := math.Exp(c.N / hBase * 180 * hD2R)
	lat := 180.0 / math.Pi * (2*math.Atan(exp) - math.Pi/2)
	lon := -180.0

	if math.Abs(-hBase-c.E) > z.size/2 {
		lon = c.E / hBase * 180
	}

	return NewLL(lat, lon)
}

// Code returns string Code of this position
func (p Position) Code() string {
	x, y := p.X, p.Y

	var code [22]byte
	var bx, by [3]uint8
	var c3x, c3y uint8

	for i := uint8(0); i < p.Level+3; i++ {
		n := int(p.Level + 2 - i)
		pow := pow3[n]
		p2c := halfPow3[n]

		if x >= p2c {
			x -= pow
			c3x = 2
		} else if x <= -p2c {
			x += pow
			c3x = 0
		} else {
			c3x = 1
		}

		if y >= p2c {
			y -= pow
			c3y = 2
		} else if y <= -p2c {
			y += pow
			c3y = 0
		} else {
			c3y = 1
		}

		if i >= 3 {
			code[i-1] = '0' + uint8(3*c3x+c3y)
		} else {
			bx[i] = c3x
			by[i] = c3y
		}
	}

	ll := p.LL()
	// Magic time. Unoptimized so far.
	if ll.Lon == -180 || ll.Lon >= 0 {
		if bx[1] == by[1] && bx[2] == by[2] {
			if bx[0] == 2 && by[0] == 1 {
				bx[0], by[0] = 1, 2
			} else if bx[0] == 1 && by[0] == 0 {
				bx[0], by[0] = 0, 1
			}
		}
	}

	base := 3*(100*int(bx[0])+10*int(bx[1])+int(bx[2])) + (100*int(by[0]) + 10*int(by[1]) + int(by[2]))
	code[0] = hChars[base/30]
	code[1] = hChars[base%30]
	return string(code[:p.Level+2])
}
