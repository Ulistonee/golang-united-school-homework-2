package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type (
	intCustomType int
)

const (
	SidesCircle   intCustomType = 0
	SidesTriangle intCustomType = 3
	SidesSquare   intCustomType = 4
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch {
	case sidesNum == SidesTriangle:
		return sideLen * sideLen * math.Sqrt(3) / 4
	case sidesNum == SidesSquare:
		return sideLen * sideLen
	case sidesNum == SidesCircle:
		return math.Pi * (sideLen * sideLen)
	default:
		return 0
	}
}
