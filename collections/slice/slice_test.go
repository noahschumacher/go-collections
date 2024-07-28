package slice

import "testing"

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestUnique(t *testing.T) {
	tcs := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "no duplicates",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "duplicates",
			input: []int{1, 1, 2, 2, 3, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "mixed",
			input: []int{1, 2, 1, 3, 2, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "all duplicates",
			input: []int{1, 1, 1, 1},
			want:  []int{1},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := Unique(tc.input)
			if !slicesEqual(got, tc.want) {
				t.Errorf("Unique(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tcs := []struct {
		name   string
		input  []int
		filter func(int) bool
		want   []int
	}{
		{
			name:  "no filter",
			input: []int{1, 2, 3},
			filter: func(int) bool {
				return true
			},
			want: []int{1, 2, 3},
		},
		{
			name:  "filter even",
			input: []int{1, 2, 3, 4},
			filter: func(i int) bool {
				return i%2 == 0
			},
			want: []int{2, 4},
		},
		{
			name:  "filter odd",
			input: []int{1, 2, 3, 4},
			filter: func(i int) bool {
				return i%2 != 0
			},
			want: []int{1, 3},
		},
		{
			name:  "empty",
			input: []int{},
			filter: func(int) bool {
				return true
			},
			want: []int{},
		},
		{
			name:  "all filtered",
			input: []int{1, 2, 3, 4},
			filter: func(int) bool {
				return false
			},
			want: []int{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := Filter(tc.input, tc.filter)
			if !slicesEqual(got, tc.want) {
				t.Errorf("Filter(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tcs := []struct {
		name  string
		input []int
		e     int
		want  bool
	}{
		{
			name:  "contains",
			input: []int{1, 2, 3},
			e:     2,
			want:  true,
		},
		{
			name:  "does not contain",
			input: []int{1, 2, 3},
			e:     4,
			want:  false,
		},
		{
			name:  "empty",
			input: []int{},
			e:     1,
			want:  false,
		},
		{
			name:  "all duplicates",
			input: []int{1, 1, 1, 1},
			e:     1,
			want:  true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := Contains(tc.input, tc.e)
			if got != tc.want {
				t.Errorf("Contains(%v, %v) = %v; want %v", tc.input, tc.e, got, tc.want)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	tcs := []struct {
		name  string
		input []int
		e     int
		want  int
	}{
		{
			name:  "contains",
			input: []int{1, 2, 3},
			e:     2,
			want:  1,
		},
		{
			name:  "does not contain",
			input: []int{1, 2, 3},
			e:     4,
			want:  -1,
		},
		{
			name:  "empty",
			input: []int{},
			e:     1,
			want:  -1,
		},
		{
			name:  "all duplicates",
			input: []int{1, 1, 1, 1},
			e:     1,
			want:  0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := IndexOf(tc.input, tc.e)
			if got != tc.want {
				t.Errorf("IndexOf(%v, %v) = %v; want %v", tc.input, tc.e, got, tc.want)
			}
		})
	}
}
