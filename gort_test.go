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

func TestSortForIntArray(t *testing.T) {
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

func TestSortForIntSlice(t *testing.T) {
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

type PeoplesSlice []People
type PeoplesArray [4]People

func TestSortForStructArray(t *testing.T) {
	getPeoplesArrayData := func() PeoplesArray {
		return PeoplesArray{
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
		array    PeoplesArray
		order    Order
		len      int
		expected PeoplesArray
	}{
		{
			array: getPeoplesArrayData(),
			order: Desc,
			len:   4,
			expected: PeoplesArray{
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
			array: getPeoplesArrayData(),
			order: Desc,
			len:   2,
			expected: PeoplesArray{
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
			array: getPeoplesArrayData(),
			order: Asc,
			len:   4,
			expected: PeoplesArray{
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
			array: getPeoplesArrayData(),
			order: Asc,
			len:   2,
			expected: PeoplesArray{
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
func TestSortForStructSlice(t *testing.T) {
	getPeoplesSliceData := func() PeoplesSlice {
		return PeoplesSlice{
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
		slice    PeoplesSlice
		order    Order
		len      int
		expected PeoplesSlice
	}{
		{
			slice: getPeoplesSliceData(),
			order: Desc,
			len:   4,
			expected: PeoplesSlice{
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
			slice: getPeoplesSliceData(),
			order: Desc,
			len:   2,
			expected: PeoplesSlice{
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
			slice: getPeoplesSliceData(),
			order: Asc,
			len:   4,
			expected: PeoplesSlice{
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
			slice: getPeoplesSliceData(),
			order: Asc,
			len:   2,
			expected: PeoplesSlice{
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
