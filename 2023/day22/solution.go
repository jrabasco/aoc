package day22

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"strconv"
	"strings"
)

type Point3 struct {
	x int
	y int
	z int
}

func parsePoint3(line string) (Point3, error) {
	parts := strings.Split(line, ",")
	pt := Point3{}
	if len(parts) != 3 {
		return pt, fmt.Errorf("invalid point spec: %s", line)
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return pt, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return pt, err
	}
	z, err := strconv.Atoi(parts[2])
	if err != nil {
		return pt, err
	}
	pt.x = x
	pt.y = y
	pt.z = z
	return pt, nil
}

type Brick struct {
	x         utils.Range
	y         utils.Range
	z         utils.Range
	supported utils.Set[*Brick]
	supports  utils.Set[*Brick]
}

func NewBrick() Brick {
	return Brick{
		x:         utils.EmptyRange,
		y:         utils.EmptyRange,
		z:         utils.EmptyRange,
		supported: utils.NewSet[*Brick](),
		supports:  utils.NewSet[*Brick](),
	}
}

//func (b Brick) points() []Point3 {
//	res := []Point3{}
//	if b.a.x != b.b.x {
//		start := min(b.a.x, b.b.x)
//		end := max(b.a.x, b.b.x)
//		for i := start; i <= end; i++ {
//			res = append(res, Point3{i, b.a.y, b.a.z})
//		}
//	} else if b.a.y != b.b.y {
//		start := min(b.a.y, b.b.y)
//		end := max(b.a.y, b.b.y)
//		for i := start; i <= end; i++ {
//			res = append(res, Point3{b.a.x, i, b.a.z})
//		}
//	} else if b.a.z != b.a.z {
//		start := min(b.a.z, b.b.z)
//		end := max(b.a.z, b.b.z)
//		for i := start; i <= end; i++ {
//			res = append(res, Point3{b.a.x, b.a.y, i})
//		}
//	} else {
//		res = append(res, b.a)
//	}
//	return res
//}

func parseBrick(line string) (Brick, error) {
	abStr := strings.Split(line, "~")
	brick := NewBrick()
	if len(abStr) != 2 {
		return brick, fmt.Errorf("invalid brick spec: %s", line)
	}
	a, err := parsePoint3(abStr[0])
	if err != nil {
		return brick, err
	}
	b, err := parsePoint3(abStr[1])
	if err != nil {
		return brick, err
	}

	brick.x = utils.NewRange(min(a.x, b.x), max(a.x, b.x))
	brick.y = utils.NewRange(min(a.y, b.y), max(a.y, b.y))
	brick.z = utils.NewRange(min(a.z, b.z), max(a.z, b.z))
	return brick, nil
}

//type Grid3 struct {
//	g    map[Point3]*Brick
//	maxZ int
//}
//
//func NewGrid3() Grid3 {
//	return Grid3{map[Point3]*Brick{}, 0}
//}
//
//func (g Grid3) Get(x, y, z int) (*Brick, bool) {
//	b, exists := g.g[Point3{x, y, z}]
//	return b, exists
//}
//
//func (g *Grid3) mark(p Point3, b *Brick) {
//	if p.z > g.maxZ {
//		g.maxZ = p.z
//	}
//	g.g[p] = b
//}
//
//func (g *Grid3) MaxZ() int {
//	return g.maxZ
//}
//
//func (g *Grid3) PutBrick(b *Brick) {
//	for _, p := range b.points() {
//		g.mark(p, b)
//	}
//}

type Bricks []Brick

func (bs *Bricks) DropBricks() {
	for i, b := range *bs {
		dropped := bs.DropBrick(b)
		if dropped.z != b.z {
			(*bs)[i] = dropped
		}
	}
}

func (bs Bricks) DropBrick(drop Brick) Brick {
	var z int
	minZ := drop.z.Start()
	for z = minZ - 1; z > 0; z-- {
		if bs.AnyOverlap(z, drop.x, drop.y) {
			break
		}
	}
	z++ // undo
	drop.z.Shift(z - minZ)
	return drop
}

func (bs Bricks) AnyOverlap(z int, x utils.Range, y utils.Range) bool {
	return len(bs.Overlaps(z, x, y)) != 0
}

func (bs Bricks) Overlaps(z int, x utils.Range, y utils.Range) []*Brick {
	res := []*Brick{}
	for i := range bs {
		b := bs[i]
		if !b.z.Contains(z) {
			continue
		}

		if b.x.Overlaps(x) && b.y.Overlaps(y) {
			res = append(res, &b)
		}
	}
	return res
}

func (bs *Bricks) FillSupport() {
	for i := range *bs {
		b := (*bs)[i]
		fmt.Println(b)
		ontop := bs.Overlaps(b.z.End()+1, b.x, b.y)
		for j := range ontop {
			t := ontop[j]
			b.supports.Add(t)
			t.supported.Add(&b)
		}
	}
}

func (bs *Bricks) Disintegreteables() []*Brick {
	res := []*Brick{}
	for i := range *bs {
		b := (*bs)[i]
		can := true
		for sup := range b.supports {
			if len(sup.supported) == 1 {
				can = false
				break
			}
		}
		if can {
			res = append(res, &b)
		}
	}
	return res
}

func Solution() int {
	parsed, err := parse.GetLinesAs[Brick]("day22/input.txt", parseBrick)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	//g := NewGrid3()
	//for i := range bricks {
	//	g.PutBrick(&bricks[i])
	//}

	bricks := Bricks(parsed)
	fmt.Println(bricks)
	bricks.DropBricks()
	bricks.FillSupport()
	fmt.Println(bricks)
	fmt.Printf("Part 1: %d\n", len(bricks.Disintegreteables()))
	return 0
}
