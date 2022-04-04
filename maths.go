package main

import "golang.org/x/exp/constraints"

type Addable interface {
	~int | ~int64 | ~string
}

func max[T constraints.Ordered](s []T) T {
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func min[T constraints.Ordered](s []T) T {
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func generator[T Addable](a T, v T) func() T {
	return func() T {
		r := a
		a = a + v
		return r
	}
}

func add[T Addable](a, b T) T {
	return a + b
}
