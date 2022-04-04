package main

import "fmt"

type AliasForInt int
type AliasForInt64 int64
type AliasForString string

func main() {
	ints := []int{1, 2, 3, 4, 5}
	chars := []string{"a", "b", "c", "d", "e"}

	fmt.Println("** Add results **")

	fmt.Println(add("hello, ", "world!"))
	fmt.Println(add(2, 7))

	fmt.Println(add(AliasForInt(1), AliasForInt(9)))
	fmt.Println(add(AliasForInt64(22), AliasForInt64(8)))
	fmt.Println(add(AliasForString("Hello "), AliasForString("Alias")))

	fmt.Println("** Compare results **")

	fmt.Println(compare("a", "b"))
	fmt.Println(compare(1, 1))

	fmt.Println("** Find results **")

	fmt.Println(find(ints, 3))
	fmt.Println(find(chars, "d"))

	fmt.Println("** For Each results **")

	// Running some logic for each element
	allEven := forEach([]int{2, 3, 4, 5, 6}, func(v int) bool { return v%2 == 0 })
	if !allEven {
		fmt.Println("Some odd")
	}

	fmt.Println("** Filter results **")

	// Filtering slices
	filteredInts := filter(ints, func(i int) bool { return i > 2 })

	fmt.Println(filteredInts)

	filteredChars := filter(chars, func(i string) bool { return i > "b" })

	fmt.Println(filteredChars)

	fmt.Println("** Map results **")

	// Run some logic/transformation on each element in a slice
	taggedInts := mapFunc(ints, func(v int) string { return "<" + fmt.Sprint(v) + ">" })
	fmt.Println(taggedInts)

	multInts := mapFunc(ints, func(v int) int { return v * 2 })
	fmt.Println(multInts)

	dblChars := mapFunc(chars, func(s string) string { return s + s })
	fmt.Println(dblChars)

	fmt.Println("** Pointer of results **")

	// Replace pointers package
	f := pointerOf("foo")
	fmt.Println(*f)

	o := pointerOf(1)
	fmt.Println(*o)

	fmt.Println("** List results **")

	l := NewList([]int{1, 2, 3})
	l.Push(4, 5, 6)
	fmt.Println(l.Slice())
	l.Remove(3)
	fmt.Println(l.Slice())
	c := l.Clone()
	fmt.Println(c.Equals(l))

	fmt.Println("** Generator results **")

	a := generator(10, 2)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println("** Min/Max results **")

	fmt.Println(max(ints))
	fmt.Println(min(ints))
}
