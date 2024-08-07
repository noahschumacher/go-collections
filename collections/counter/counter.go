package counter

import (
	"sort"
	"strings"
)

// A Counter is a map containing an int count of the item. The count cannot be
// negative and all non-included elements have a count of 0.
type Counter[T comparable] map[T]int

func New[T comparable](size int) Counter[T] {
	return make(Counter[T], size)
}

// Increment a value by 1
func (c Counter[T]) Inc(k T) {
	c[k]++
}

// Increment a value by 1
func (c Counter[T]) Dec(k T) {
	c[k]--
	c.del0(k)
}

// Delete entries with <= 0 count
func (c Counter[T]) del0(k T) {
	if c[k] <= 0 {
		delete(c, k)
	}
}

// Add v counts of a element. If the element does not exist it will be added
// with count v.
func (c Counter[T]) Add(k T, v int) {
	c[k] += v
}

// Subtract v counts of a element. If the subtraction leaves the element with <=
// 0 it will be deleted
func (c Counter[T]) Subtract(k T, v int) {
	c[k] -= v
	c.del0(k)
}

// Get the count of an element.
func (c Counter[T]) Get(k T) int {
	return c[k]
}

// Return the length of map.
func (c Counter[T]) Size() int {
	return len(c)
}

// Get a slice of the elements in the Counter.
func (c Counter[T]) Elements() []T {
	s := make([]T, 0, c.Size())
	for key := range c {
		s = append(s, key)
	}
	return s
}

// Returns the most common element. If multiple elements have the same maximum
// count, only one will returned with no guarentees.
func (c Counter[T]) MostCommon() T {
	var mostCommon T
	var maxValue int

	for key, val := range c {
		if val > maxValue {
			mostCommon = key
			maxValue = val
		}
	}
	return mostCommon
}

type item[T comparable] struct {
	key T
	c   int
}

// Return the N most comment elements ordered by their counts. If elements have
// the same count their order is not guarenteed. A max of n elements is ever
// returned.
func (c Counter[T]) MostCommonN(n int) []T {
	if n == 0 || len(c) == 0 {
		return []T{}
	}

	// Don't allow n to be larger than the counter list
	if n > len(c) {
		n = len(c)
	}

	// Convert the counter into a slice of items
	mc := make([]item[T], 0, c.Size())
	for key, val := range c {
		mc = append(mc, item[T]{key, val})
	}

	// Sort the item slice by the count of the items
	sort.Slice(mc, func(i, j int) bool {
		return mc[i].c > mc[j].c
	})

	mostCommon := make([]T, n)
	for i := 0; i < n; i++ {
		mostCommon[i] = mc[i].key
	}

	return mostCommon
}

// Get the total count of all elements.
func (c Counter[T]) Total() int {
	total := 0
	for _, val := range c {
		total += val
	}
	return total
}

// Add a counter to the current counter.
func (c Counter[T]) AddCounter(c2 Counter[T]) {
	for key, val := range c2 {
		c[key] += val
	}
}

// Subtract a counter from the current counter. All <= 0
// elements will be deleted.
func (c Counter[T]) SubtractCounter(c2 Counter[T]) {
	for key, val := range c2 {
		c[key] -= val
		c.del0(key)
	}
}

// -----------------------------------------------------------------------------
// Creation helpers

// Create a counter from a map. The int value in the map is the count. If the
// int value is <= 0 the element will be ignored from the count.
func CounterFromMap[T comparable](m map[T]int) Counter[T] {
	c := New[T](len(m))
	for key, val := range m {
		if val <= 0 {
			continue
		}
		c.Add(key, val)
	}
	return c
}

// Create a counter from a slice. The count of an element will be the number of
// times it appears in the slice.
func CounterFromSlice[T comparable](s []T) Counter[T] {
	c := New[T](len(s))
	for _, v := range s {
		c.Inc(v)
	}
	return c
}

// Breaks a string into each element of the string and counts the characters.
func CounterFromString(s string) Counter[string] {
	characters := strings.Split(s, "")
	return CounterFromSlice(characters)
}
