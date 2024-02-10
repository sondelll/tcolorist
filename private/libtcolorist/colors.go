package libtcolorist

import (
	"fmt"
	"strconv"
)

type RGB8 struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func removeHash(color string) string {
	if color[0] == '#' {
		return color[1:]
	}
	return color
}

func hexToUint8(hex string) uint8 {
	value, err := strconv.ParseUint(hex, 16, 8)
	if err != nil {
		return uint8(0)
	}
	return uint8(value)
}

func ParseColor8Bit(color string) (RGB8, error) {
	col := removeHash(color)
	if len(col) != 6 {
		return RGB8{0, 0, 0}, fmt.Errorf("invalid color format")
	}
	r := hexToUint8(col[0:2])
	g := hexToUint8(col[2:4])
	b := hexToUint8(col[4:6])

	return RGB8{r, g, b}, nil
}
