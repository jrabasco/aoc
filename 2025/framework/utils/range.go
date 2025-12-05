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

func (r Range) FullyContains(o Range) bool {
	return r.start <= o.start && o.end <= r.end
}

func (r Range) Contains(num int) bool {
	return r.start <= num && num <= r.end
}

func (r Range) Overlaps(o Range) bool {
	return r.Contains(o.start) || r.Contains(o.end) || o.FullyContains(r)
}

func (r Range) Merge(o Range) Range {
	if !r.Overlaps(o) {
		panic("trying to merge non-overlapping ranges")
	}
	st := r.Start()
	if r.Start() > o.Start() {
		st = o.Start()
	}

	end := r.End()
	if o.End() > r.End() {
		end = o.End()
	}
	return NewRange(st, end)
}

// returns the range that's in r and an array of ranges outside of it
func (r Range) Split(o Range) (Range, []Range) {
	res := []Range{}
	// o is fully contained
	if r.start <= o.start && o.end <= r.end {
		return o, res
	}

	// o's beginning is contained
	if r.start <= o.start && r.end < o.end {
		res = append(res, Range{r.end + 1, o.end})
		return Range{o.start, r.end}, res
	}

	// o's end is contained
	if o.start < r.start && o.end >= r.start && o.end <= r.end {
		res = append(res, Range{o.start, r.start - 1})
		return Range{r.start, o.end}, res
	}

	// o contains r
	if o.start < r.start && r.end < o.end {
		res = append(res, Range{o.start, r.end - 1})
		res = append(res, Range{r.end + 1, o.end})
		return Range{r.start, r.end}, res
	}

	// disjoint
	res = append(res, o)
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
