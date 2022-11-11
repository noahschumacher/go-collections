# go-collections
Golang collection types using Generics

## Collection Types
1. [Set](#set)
2. [Counter](#counter)

### `Set`
A unique set of `Comparable` values. `Set`'s are useful for quickly getting the unique values of a slice and performing set operations such as
* `Difference`
* `Union`
* `Intersection`
* `Complement`

There are also utility functions for using set operations with slices. These helpers follow the syntax `SliceUnion`, `SliceIntersection`, etc.

### `Counter`
A map that contains the counts of `Comparable` items. Counters can be created from
* Maps where the value is an int.
* Slice where the number of times an item is repeated is the count.
* String where each character is parsed as an item.

Some useful methods of counters include
* `MostCommon` - get the most common element
* `MostCommonN` - get the N most common elements
* `AddCounter`/`SubtractCounter` - add/subtract one counter from another
