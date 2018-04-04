package gort

import (
	"reflect"
	"testing"
)

type Order int

const (
	Desc Order = iota
	Asc
)

func TestSortIntArray(t *testing.T) {
	getArrayTestData := func() [4]int {
		return [4]int{-20, 2, 1, -10}
	}

	tests := []struct {
		array    [4]int
		order    Order
		len      int
		expected [4]int
	}{
		{
			array:    getArrayTestData(),
			order:    Desc,
			len:      4,
			expected: [4]int{2, 1, -10, -20},
		},
		{
			array:    getArrayTestData(),
			order:    Desc,
			len:      2,
			expected: [4]int{2, -20, 1, -10},
		},
		{
			array:    getArrayTestData(),
			order:    Asc,
			len:      4,
			expected: [4]int{-20, -10, 1, 2},
		},
		{
			array:    getArrayTestData(),
			order:    Asc,
			len:      2,
			expected: [4]int{-20, 2, 1, -10},
		},
	}

	for _, test := range tests {
		var order func(i, j int) bool
		if test.order == Desc {
			order = func(i, j int) bool {
				return test.array[i] > test.array[j]
			}
		} else {
			order = func(i, j int) bool {
				return test.array[i] < test.array[j]
			}
		}

		Sort(&test.array, test.len, order)

		if !reflect.DeepEqual(test.array, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, test.array)
		}
	}
}

func TestSortIntSlice(t *testing.T) {
	getTestSliceData := func() []int {
		return []int{-20, 2, 1, -10}
	}

	tests := []struct {
		slice    []int
		order    Order
		len      int
		expected []int
	}{
		{
			slice:    getTestSliceData(),
			order:    Desc,
			len:      4,
			expected: []int{2, 1, -10, -20},
		},
		{
			slice:    getTestSliceData(),
			order:    Desc,
			len:      2,
			expected: []int{2, -20, 1, -10},
		},
		{
			slice:    getTestSliceData(),
			order:    Asc,
			len:      4,
			expected: []int{-20, -10, 1, 2},
		},
		{
			slice:    getTestSliceData(),
			order:    Asc,
			len:      2,
			expected: []int{-20, 2, 1, -10},
		},
	}

	for _, test := range tests {
		var order func(i, j int) bool
		if test.order == Desc {
			order = func(i, j int) bool {
				return test.slice[i] > test.slice[j]
			}
		} else {
			order = func(i, j int) bool {
				return test.slice[i] < test.slice[j]
			}
		}

		Sort(&test.slice, test.len, order)

		if !reflect.DeepEqual(test.slice, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, test.slice)
		}
	}
}

type People struct {
	Name string
	Age  int
}

type PeopleSlice []People
type PeopleArray [4]People

func TestSortStructArray(t *testing.T) {
	getPeopleArrayData := func() PeopleArray {
		return PeopleArray{
			People{
				Name: "a",
				Age:  2,
			},
			People{
				Name: "b",
				Age:  30,
			},
			People{
				Name: "c",
				Age:  1,
			},
			People{
				Name: "d",
				Age:  3,
			},
		}
	}

	tests := []struct {
		array    PeopleArray
		order    Order
		len      int
		expected PeopleArray
	}{
		{
			array: getPeopleArrayData(),
			order: Desc,
			len:   4,
			expected: PeopleArray{
				People{
					Name: "b",
					Age:  30,
				},
				People{
					Name: "d",
					Age:  3,
				},
				People{
					Name: "a",
					Age:  2,
				},
				People{
					Name: "c",
					Age:  1,
				},
			},
		},
		{
			array: getPeopleArrayData(),
			order: Desc,
			len:   2,
			expected: PeopleArray{
				People{
					Name: "b",
					Age:  30,
				},
				People{
					Name: "a",
					Age:  2,
				},
				People{
					Name: "c",
					Age:  1,
				},
				People{
					Name: "d",
					Age:  3,
				},
			},
		},
		{
			array: getPeopleArrayData(),
			order: Asc,
			len:   4,
			expected: PeopleArray{
				People{
					Name: "c",
					Age:  1,
				},
				People{
					Name: "a",
					Age:  2,
				},
				People{
					Name: "d",
					Age:  3,
				},
				People{
					Name: "b",
					Age:  30,
				},
			},
		},
		{
			array: getPeopleArrayData(),
			order: Asc,
			len:   2,
			expected: PeopleArray{
				People{
					Name: "a",
					Age:  2,
				},
				People{
					Name: "b",
					Age:  30,
				},
				People{
					Name: "c",
					Age:  1,
				},
				People{
					Name: "d",
					Age:  3,
				},
			},
		},
	}

	for _, test := range tests {
		var order func(i, j int) bool
		if test.order == Desc {
			order = func(i, j int) bool {
				return test.array[i].Age > test.array[j].Age
			}
		} else {
			order = func(i, j int) bool {
				return test.array[i].Age < test.array[j].Age
			}
		}

		Sort(&test.array, test.len, order)

		if !reflect.DeepEqual(test.array, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, test.array)
		}
	}
}
func TestSortStructSlice(t *testing.T) {
	getPeopleSliceData := func() PeopleSlice {
		return PeopleSlice{
			People{
				Name: "a",
				Age:  2,
			},
			People{
				Name: "b",
				Age:  30,
			},
			People{
				Name: "c",
				Age:  1,
			},
			People{
				Name: "d",
				Age:  3,
			},
		}
	}

	tests := []struct {
		slice    PeopleSlice
		order    Order
		len      int
		expected PeopleSlice
	}{
		{
			slice: getPeopleSliceData(),
			order: Desc,
			len:   4,
			expected: PeopleSlice{
				People{
					Name: "b",
					Age:  30,
				},
				People{
					Name: "d",
					Age:  3,
				},
				People{
					Name: "a",
					Age:  2,
				},
				People{
					Name: "c",
					Age:  1,
				},
			},
		},
		{
			slice: getPeopleSliceData(),
			order: Desc,
			len:   2,
			expected: PeopleSlice{
				People{
					Name: "b",
					Age:  30,
				},
				People{
					Name: "a",
					Age:  2,
				},
				People{
					Name: "c",
					Age:  1,
				},
				People{
					Name: "d",
					Age:  3,
				},
			},
		},
		{
			slice: getPeopleSliceData(),
			order: Asc,
			len:   4,
			expected: PeopleSlice{
				People{
					Name: "c",
					Age:  1,
				},
				People{
					Name: "a",
					Age:  2,
				},
				People{
					Name: "d",
					Age:  3,
				},
				People{
					Name: "b",
					Age:  30,
				},
			},
		},
		{
			slice: getPeopleSliceData(),
			order: Asc,
			len:   2,
			expected: PeopleSlice{
				People{
					Name: "a",
					Age:  2,
				},
				People{
					Name: "b",
					Age:  30,
				},
				People{
					Name: "c",
					Age:  1,
				},
				People{
					Name: "d",
					Age:  3,
				},
			},
		},
	}

	for _, test := range tests {
		var order func(i, j int) bool
		if test.order == Desc {
			order = func(i, j int) bool {
				return test.slice[i].Age > test.slice[j].Age
			}
		} else {
			order = func(i, j int) bool {
				return test.slice[i].Age < test.slice[j].Age
			}
		}

		Sort(&test.slice, test.len, order)

		if !reflect.DeepEqual(test.slice, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, test.slice)
		}
	}
}
