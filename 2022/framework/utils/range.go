package utils

type Range struct {
	start int
	end int
}

var EmptyRange = Range{0,0}

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
