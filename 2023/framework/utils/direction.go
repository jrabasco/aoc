package utils

type Direction int

const (
	RIGHT Direction = iota
	DOWN
	LEFT
	UP
	NODIR
)

func (d Direction) Reverse() Direction {
	switch d {
	case RIGHT:
		return LEFT
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	case UP:
		return DOWN
	default:
		return NODIR
	}
}

func (d Direction) Turn(to Direction) Direction {
	if d == UP {
		return to
	}

	// assuming "turning to UP" means carry on straight
	if to == UP {
		return d
	}

	if d == DOWN {
		return to.Reverse()
	}

	// assuming "turning to DOWN" means reverse
	if to == DOWN {
		return d.Reverse()
	}

	switch d {
	case RIGHT:
		switch to {
		case LEFT:
			return UP
		case RIGHT:
			return DOWN
		}
	case LEFT:
		switch to {
		case LEFT:
			return DOWN
		case RIGHT:
			return UP
		}
	}
	return NODIR
}

func (d Direction) String() string {
	switch d {
	case RIGHT:
		return "right"
	case DOWN:
		return "down"
	case LEFT:
		return "left"
	case UP:
		return "up"
	default:
		return "meh"
	}
}
