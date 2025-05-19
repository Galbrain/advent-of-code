package guard

import (
	"errors"
	"fmt"
)

type Direction string

const (
	North Direction = "North"
	East  Direction = "East"
	South Direction = "South"
	West  Direction = "West"
)

func (d Direction) Validate() error {
	switch d {
	case North, East, South, West:
		return nil
	default:
		return errors.New(fmt.Sprintf("invalid value: must be one of the following: %v", []Direction{North, East, South, West}))
	}
}
