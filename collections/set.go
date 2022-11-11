package collections

// A collection of unique comparable items. Uses a map with only true values
// to accomplish set functionality.
type Set[T comparable] map[T]bool

// Create a new empty set with the specified initial size.
func NewSet[T comparable](size int) Set[T] {
	return make(Set[T], size)
}

// Add a new key to the set
func (s Set[T]) Add(key T) {
	s[key] = true
}

// Remove a key from the set. If the key is not in the set then noop
func (s Set[T]) Remove(key T) {
	delete(s, key)
}

// Check if Set s contains key
func (s Set[T]) Contains(key T) bool {
	return s[key]
}

// A union B
func (a Set[T]) Union(b Set[T]) Set[T] {
	small, large := smallLarge(a, b)

	for key := range small {
		large.Add(key)
	}
	return large
}

// A intersect B
func (a Set[T]) Intersection(b Set[T]) Set[T] {
	small, large := smallLarge(a, b)

	resultSet := NewSet[T](0)
	for key := range small {
		if large.Contains(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

// A compliment
func (a Set[T]) Complement(b Set[T]) Set[T] {
	resultSet := NewSet[T](0)
	for key := range b {
		if !a.Contains(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

// A difference B | NOTE: A-B != B-A
func (a Set[T]) Difference(b Set[T]) Set[T] {
	resultSet := NewSet[T](0)
	for key := range a {
		if !b.Contains(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

// A == B (all elements of A are in B and vice versa)
func (a Set[T]) Equals(b Set[T]) bool {
	return len(a.Difference(b)) == 0 && len(b.Difference(a)) == 0
}

// Turn a Set into a slice
func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}

	return slice
}

// returns the small and large sets according to their len
func smallLarge[T comparable](a, b Set[T]) (Set[T], Set[T]) {
	small, large := b, a
	if len(b) > len(a) {
		small, large = a, b
	}

	return small, large
}

// -------------------------------------------------
// SLICE HELPERS

// Create a Set from a slice.
func SliceToSet[T comparable](s []T) Set[T] {
	set := NewSet[T](len(s))
	for _, item := range s {
		set.Add(item)
	}
	return set
}

// Map a slice to a set using a function f
func MapSliceToSet[S any, T comparable](s []S, f func(s S) T) Set[T] {
	set := NewSet[T](len(s))
	for _, item := range s {
		set.Add(f(item))
	}
	return set
}

// Union two slices. The provided slices do not need to be unique. Order not guaranteed.
func SliceUnion[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	union := aSet.Union(bSet)
	return union.ToSlice()
}

// Intersection of two slices. The provided slices do not need to be unique. Order not guaranteed.
func SliceIntersection[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	intersection := aSet.Intersection(bSet)
	return intersection.ToSlice()
}

// Complement of A with regards to B. Slices do not need to be unique. Order not guaranteed.
func SliceComplement[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	complement := aSet.Complement(bSet)
	return complement.ToSlice()
}

// Difference of A-B. Slices do not need to be unique. Order not guaranteed.
func SliceDifference[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	difference := aSet.Difference(bSet)
	return difference.ToSlice()
}
