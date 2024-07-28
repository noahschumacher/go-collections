package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterFromMap(t *testing.T) {
	c := CounterFromMap(map[string]int{"a": 4, "b": 0, "c": -20, "d": 2})

	assert.Equal(t, 4, c.Get("a"), "should be 4")
	assert.Equal(t, 0, c.Get("c"), "c should be 0 aswell")
	assert.ElementsMatch(t, []string{"a", "d"}, c.Elements(), "elements mismatch")
}

func TestCounterFromSlice(t *testing.T) {
	c := CounterFromSlice([]int{0, -1, 2, 2, 2, 2, 2, 2, -100, 23})

	assert.Equal(t, 1, c.Get(0))
	assert.Equal(t, 1, c.Get(-100))
	assert.Equal(t, 6, c.Get(2))
	assert.ElementsMatch(t, []int{-100, -1, 0, 2, 23}, c.Elements(), "elements mismatch")
}

func TestCounterFromString(t *testing.T) {
	c := CounterFromString("Noah Schumacher")

	assert.Equal(t, 2, c.Get("a"))
	assert.Equal(t, 1, c.Get(" "))
	assert.Equal(t, 1, c.Get("S"))
	assert.ElementsMatch(
		t,
		[]string{"N", "S", "o", "a", "c", "h", "u", "m", "e", "r", " "},
		c.Elements(),
		"elements mismatch",
	)
}

func TestMostCommont(t *testing.T) {

	a := CounterFromSlice([]string{"a", "b", "b", "b", "z", "z"})
	assert.Equal(t, "b", a.MostCommon(), "most common should be b")

	b := CounterFromSlice([]int{0, 1, 0, 2, 3, 4, 3, 3, 3, 10000, -23})
	assert.Equal(t, 3, b.MostCommon(), "most common should be b")

	c := CounterFromSlice([]int{})
	assert.Equal(t, 0, c.MostCommon(), "most common should be 0")
}

func TestMostCommontN(t *testing.T) {

	a := CounterFromSlice([]string{"a", "b", "b", "b", "z", "z"})
	assert.Equal(t, []string{"b", "z"}, a.MostCommonN(2))

	b := CounterFromSlice([]int{-2, 0, 0, 1, 2, 3, 3, 3, 3})
	mcb := b.MostCommonN(3)
	assert.Equal(t, []int{3, 0}, mcb[0:2])
	ok := false
	for _, v := range []int{-2, 1, 2} {
		if v == mcb[2] {
			ok = true
			break
		}
	}
	if !ok {
		assert.Fail(t, "did not find the third value", mcb[2])
	}

	c := CounterFromSlice([]int{})
	assert.Equal(t, []int{}, c.MostCommonN(1))
}

func TestAddCounter(t *testing.T) {
	a := CounterFromSlice([]int{0, 1, 2, 2, 3, 3, 3})
	b := CounterFromSlice([]int{-1, 0, 1, 2, 3, 4})
	a.AddCounter(b)

	exp := Counter[int]{
		-1: 1, 0: 2, 1: 2, 2: 3, 3: 4, 4: 1,
	}

	for key, val := range exp {
		assert.Equal(t, val, a.Get(key), "key: %v", key)
		a.Subtract(key, a.Get(key))
	}

	assert.Equal(t, 0, a.Size())
}

func TestSubtractCounter(t *testing.T) {
	a := CounterFromSlice([]int{0, 1, 2, 2, 3, 3, 3, 4})
	b := CounterFromSlice([]int{-1, 0, 1, 2, 3, 4, 4, 4})
	b.SubtractCounter(a)

	exp := Counter[int]{-1: 1, 4: 2}
	for key, val := range exp {
		assert.Equal(t, val, b.Get(key), "key: %v", key)
		b.Subtract(key, b.Get(key))
	}

	assert.Equal(t, 0, b.Size())
}
