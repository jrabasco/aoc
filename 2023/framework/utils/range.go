package utils

type Range struct {
	start int
	end   int
}

var EmptyRange = Range{0, 0}

func (r Range) Start() int {
	return r.start
}

func (r Range) End() int {
	return r.end
}

func (r Range) Empty() bool {
	return r.end <= r.start
}

func NewRange(start int, end int) Range {
	return Range{start, end}
}

func (r Range) FullyContains(other Range) bool {
	return r.start <= other.start && other.end <= r.end
}

func (r Range) Contains(num int) bool {
	return r.start <= num && num <= r.end
}

func (r Range) Overlaps(o Range) bool {
	return r.Contains(o.start) || r.Contains(o.end) || o.FullyContains(r)
}

// returns the range that's in r and an array of ranges outside of it
func (r Range) Split(other Range) (Range, []Range) {
	res := []Range{}
	// other is fully contained
	if r.start <= other.start && other.end <= r.end {
		return other, res
	}

	// other's beginning is contained
	if r.start <= other.start && r.end < other.end {
		res = append(res, Range{r.end + 1, other.end})
		return Range{other.start, r.end}, res
	}

	// other's end is contained
	if other.start < r.start && other.end >= r.start && other.end <= r.end {
		res = append(res, Range{other.start, r.start - 1})
		return Range{r.start, other.end}, res
	}

	// other contains r
	if other.start < r.start && r.end < other.end {
		res = append(res, Range{other.start, r.end - 1})
		res = append(res, Range{r.end + 1, other.end})
		return Range{r.start, r.end}, res
	}

	// disjoint
	res = append(res, other)
	return EmptyRange, res
}

// inclusive ranges
func (r Range) Len() int {
	return r.end - r.start + 1
}

func (r *Range) MoveStart(start int) {
	r.start = start
}

func (r *Range) MoveEnd(end int) {
	r.end = end
}

func (r *Range) Shift(delta int) {
	r.start += delta
	r.end += delta
}
