package utils

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

var member struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

func SetFromSlice[T comparable](slice []T) Set[T] {
	set := Set[T]{}
	set.AddSlice(slice)
	return set
}

func (s *Set[T]) Add(elms ...T) {
	for _, elm := range elms {
		(*s)[elm] = member
	}
}

func (s *Set[T]) AddSlice(slice []T) {
	for _, elm := range slice {
		s.Add(elm)
	}
}

func (s *Set[T]) AddSet(set Set[T]) {
	for elm := range set {
		s.Add(elm)
	}
}

func (s *Set[T]) Remove(elm T) {
	delete(*s, elm)
}

func (s Set[T]) Contains(elm T) bool {
	_, exists := s[elm]
	return exists
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	res := make(Set[T])
	for elm := range s {
		if other.Contains(elm) {
			res.Add(elm)
		}
	}
	return res
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	res := make(Set[T])
	res.AddSet(s)
	res.AddSet(other)
	return res
}

func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for elm := range s {
		if !other.Contains(elm) {
			return false
		}
	}
	return true
}

func (s Set[T]) Peek() T {
	for elm := range s {
		return elm
	}
	var zero T
	return zero
}

func (s Set[T]) String() string {
	strElms := []string{}
	for elm := range s {
		strElms = append(strElms, fmt.Sprintf("%v", elm))
	}
	return fmt.Sprintf("{%v}", strings.Join(strElms, " "))
}
