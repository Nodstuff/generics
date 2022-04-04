# Go Generics Examples

### Map
```go
func mapFunc[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}
```
### Filter
```go
func filter[T any](a []T, f func(T) bool) []T {
	var n []T
	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
```
### For Each
```go
func forEach[T any](a []T, f func(T) bool) bool {
	for _, e := range a {
		if !f(e) {
			return false
		}
	}

	return true
}
```
### Find
```go
func find[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}
```
### Compare
```go
func compare[T comparable](lhs, rhs T) bool {
	return lhs == rhs
}
```
### Min/Max
```go
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
```
### Add
```go
type Addable interface {
    ~int | ~int64 | ~string
}

func add[T Addable](a, b T) T {
	return a + b
}
```
### Pointers
```go
func pointerOf[T any](v T) *T {
	return &v
}
```