package main

func compare[T comparable](lhs, rhs T) bool {
	return lhs == rhs
}

func find[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}
