package main

type fn func(string)
type fs func(string) string
type fb func(string) bool

type StringSlice []string

// Invokes passed func on every element. Comparable to underscore’s each.
func Each(f fn, slice []string) {
	for _, s := range slice {
		f(s)
	}
}

func Filter(f fb, slice []string) []string {
	out := make([]string, 0)

	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}

	return out
}

// Returns first element which returns true for passed func. Comparable to Linq’s First or underscore’s find.
// Returns error if no elements satisfy the func.
func First(f fb, slice []string) string {
	return Head(Filter(f, slice))
}

// Returns the first element of the Slice, if slice is empty, panic
func Head(slice []string) string {
	return slice[0]
}

// Create a new slice, of same length of original slice, by mapping each value through
// the function func. Comparable to underscore’s map.
func Map(f fs, slice []string) []string {
	out := make([]string, len(slice))
	for i, s := range slice {
		out[i] = f(s)
	}
	return out
}

// Returns the a new slice with the first element removed, if slice is empty, panic
func Tail(slice []string) []string {
	return slice[1:]
}
