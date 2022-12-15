package day15

import (
	"fmt"
	"github.com/jrabasco/aoc/2022/framework/parse"
	"strconv"
	"strings"
)

type Point interface {
	X() int
	Y() int
}

type Beacon struct {
	x int
	y int
}

func NewBeacon(x, y int) Beacon {
	return Beacon{x, y}
}

func (b Beacon) X() int {
	return b.x
}

func (b Beacon) Y() int {
	return b.y
}

type Sensor struct {
	x    int
	y    int
	dist int
}

func NewSensor(x, y int) Sensor {
	return Sensor{x, y, 0}
}

func (s Sensor) X() int {
	return s.x
}

func (s Sensor) Y() int {
	return s.y
}

func (s Sensor) Dist() int {
	return s.dist
}

func (s *Sensor) MarkBeacon(b Beacon) {
	s.dist = dist(s, b)
}

type Map struct {
	beacons []Beacon
	sensors []Sensor
	minX    int
	maxX    int
}

func NewMap() Map {
	return Map{[]Beacon{}, []Sensor{}, 9223372036854775807, -9223372036854775808}
}

func (m *Map) AddBeacon(b Beacon) {
	for _, mb := range m.beacons {
		if b.X() == mb.X() && b.Y() == mb.Y() {
			return
		}
	}
	m.beacons = append(m.beacons, b)
}

func (m *Map) AddSensor(s Sensor) {
	pos := len(m.sensors)
	for i, cs := range m.sensors {
		if s.X() < cs.X() {
			pos = i
			break
		}
	}

	if pos == len(m.sensors) {
		m.sensors = append(m.sensors, s)
	} else {
		m.sensors = append(m.sensors[:pos+1], m.sensors[pos:]...)
		m.sensors[pos] = s
	}

	if s.X()-s.Dist() < m.minX {
		m.minX = s.X() - s.Dist()
	}

	if s.X()+s.Dist() > m.maxX {
		m.maxX = s.X() + s.Dist()
	}
}

func (m Map) MinX() int {
	return m.minX
}

func (m Map) MaxX() int {
	return m.maxX
}

func (m Map) BeaconsOnLine(y int) []Beacon {
	res := []Beacon{}
	for _, b := range m.beacons {
		if y == b.Y() {
			res = append(res, b)
		}
	}
	return res
}

func parseSensor(m *Map, line string) error {
	sb := strings.Split(line, ": ")
	if len(sb) != 2 {
		return fmt.Errorf("invalid sensor line: %s")
	}

	if len(sb[0]) < 10 || len(sb[1]) < 21 {
		return fmt.Errorf("invalid sensor line: %s")
	}

	// skip "closest beacon is at "
	beaconS := sb[1][21:]
	b, err := parsePoint[Beacon](beaconS, NewBeacon)
	if err != nil {
		return err
	}

	m.AddBeacon(b)

	// skip "Sensor at "
	sensorS := sb[0][10:]
	s, err := parsePoint[Sensor](sensorS, NewSensor)
	if err != nil {
		return err
	}
	s.MarkBeacon(b)
	m.AddSensor(s)

	return nil
}

func abs(e int) int {
	if e < 0 {
		return -e
	}
	return e
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dist(p1, p2 Point) int {
	return abs(p1.X()-p2.X()) + abs(p1.Y()-p2.Y())
}

func distInts(p1 Point, x, y int) int {
	return abs(p1.X()-x) + abs(p1.Y()-y)
}

func parsePoint[T Point](pointS string, newPoint func(int, int) T) (T, error) {
	parts := strings.Split(pointS, ", ")
	if len(parts) < 2 {
		return newPoint(0, 0), fmt.Errorf("invalid point string: %s", pointS)
	}

	partsX := strings.Split(parts[0], "=")
	if len(partsX) < 2 {
		return newPoint(0, 0), fmt.Errorf("invalid point string: %s", pointS)
	}

	x, err := strconv.Atoi(partsX[1])
	if err != nil {
		return newPoint(0, 0), fmt.Errorf("could not parse x: %v", err)
	}

	partsY := strings.Split(parts[1], "=")
	if len(partsY) < 2 {
		return newPoint(0, 0), fmt.Errorf("invalid point string: %s", pointS)
	}

	y, err := strconv.Atoi(partsY[1])
	if err != nil {
		return newPoint(0, 0), fmt.Errorf("could not parse y: %v", err)
	}

	return newPoint(x, y), nil
}

type Range struct {
	from int
	to   int
}

func (r Range) Len() int {
	return r.to - r.from + 1
}

func genRanges(m *Map, y, minX, maxX int) []Range {
	ranges := []Range{}
	for _, s := range m.sensors {
		dy := abs(y - s.Y())
		if dy > s.Dist() {
			continue
		}
		dx := s.Dist() - dy
		from := max(s.X()-dx, minX)
		to := s.X() + dx

		if to < minX {
			continue
		}

		if from > maxX {
			break
		}

		curRange := Range{from, to}
		inserted := false
		for i := 0; i < len(ranges); i++ {
			r := ranges[i]
			if curRange.from > r.to+1 {
				continue
			}

			if curRange.to < r.from-1 {
				ranges = append(ranges[:i+1], ranges[i:]...)
				ranges[i] = curRange
				inserted = true
				break
			}

			// r is fully within curRange
			if curRange.from <= r.from && r.to <= curRange.to {
				ranges = append(ranges[:i], ranges[i+1:]...)
				continue
			}

			// curRange is fully within r
			if r.from <= curRange.from && curRange.to <= r.to {
				inserted = true
				break
			}

			// r overlaps on the left of curRange
			if r.from <= curRange.from && r.to < curRange.to {
				// r is the last of the ranges
				if i >= len(ranges)-1 {
					// extend on the right of r and break
					ranges[i].to = curRange.to
					inserted = true
					break
				}

				next := ranges[i+1]

				// curRange does not overlap with the next range
				if curRange.to < next.from-1 {
					// extend on the right of r and break
					ranges[i].to = curRange.to
					inserted = true
					break
				}

				// curRange spans r and next
				// merge r into curRange and then drop it from ranges
				if curRange.to >= next.from-1 {
					curRange.from = r.from
					// remove element i
					ranges = append(ranges[:i], ranges[i+1:]...)
					i--
					continue
				}
			}

			// r overlaps on the right of curRange
			// we know there is no overlap with the previous one
			if curRange.from <= r.from && curRange.to < r.to {
				ranges[i].from = curRange.from
				inserted = true
				break
			}
		}

		if !inserted {
			ranges = append(ranges, curRange)
		}

	}
	return ranges
}

func countUnavailable(m *Map, y int) int {
	ranges := genRanges(m, y, -9223372036854775808, 9223372036854775807)

	res := 0
	beacons := m.BeaconsOnLine(y)
	for _, r := range ranges {
		res += r.Len()
		for _, b := range beacons {
			if r.from <= b.X() && b.X() <= r.to {
				res -= 1
			}
		}
	}
	return res
}

func findBeacon(m *Map, maxCoord int) Beacon {
	for y := 0; y <= maxCoord; y++ {
		ranges := genRanges(m, y, 0, maxCoord)

		if len(ranges) == 2 {
			return NewBeacon(ranges[0].to+1, y)
		}
	}
	return NewBeacon(0, 0)
}

func solvePart(part string) int {
	test := ""
	p1y := 2000000
	maxCoord := 4000000
	if part == "test" {
		p1y = 10
		maxCoord = 20
		test = "_test"
	}
	fileName := fmt.Sprintf("day15/input%s.txt", test)
	parsed, err := parse.GetLines(fileName)
	if err != nil {
		fmt.Printf("Failed to parse input : %v\n", err)
		return 1
	}

	m := NewMap()

	for _, line := range parsed {
		parseSensor(&m, line)
	}

	if part == "1" {
		fmt.Printf("Part %s: %v\n", part, countUnavailable(&m, p1y))
	} else if part == "2" {
		b := findBeacon(&m, maxCoord)
		fmt.Printf("Part %s: %v\n", part, b.X()*4000000+b.Y())
	} else {
		fmt.Println(m)
		fmt.Printf("Part 1 %s: %v\n", part, countUnavailable(&m, p1y))
		b := findBeacon(&m, maxCoord)
		fmt.Printf("Part 2 %s: %v\n", part, b.X()*4000000+b.Y())
	}
	return 0
}

func Solution(part string) int {
	if part != "1" && part != "2" && part != "test" {
		p1 := solvePart("1")
		if p1 != 0 {
			return p1
		}
		return solvePart("2")
	} else {
		return solvePart(part)
	}
}
