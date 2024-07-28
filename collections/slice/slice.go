package slice

import "github.com/noahschumacher/go-collections/collections/set"

// Unique returns a new slice with duplicates removed. This function does not
// modify the input slice. The order of the unique elements is preserved.
func Unique[T comparable](s []T) []T {
	seen := set.New[T](0)

	new := make([]T, 0, len(s))
	for _, v := range s {
		if !seen.Contains(v) {
			new = append(new, v)
			seen.Add(v)
		}
	}

	return new
}

// Filter returns a new slice with elements removed as specified by the provided
// filter function. This function modifies the input slice.
func Filter[T any](s []T, filter func(T) bool) []T {
	i := 0
	for _, v := range s {
		if filter(v) {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

// Contains returns true if the slice contains the element.
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the element in the slice. If the element is not
// in the slice, -1 is returned.
func IndexOf[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

// Map applies the function f to each element in the slice and returns a new
// slice with the results.
func Map[T, U any](s []T, f func(T) U) []U {
	new := make([]U, len(s))
	for i, v := range s {
		new[i] = f(v)
	}
	return new
}
