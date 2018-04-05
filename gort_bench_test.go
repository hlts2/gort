package gort

import (
	"sort"
	"testing"
)

const LengthOfArrayAndSliceForBenchmarkTest = 100

func BenchmarkGortIntArray(b *testing.B) {
	b.StopTimer()

	var unsorted [LengthOfArrayAndSliceForBenchmarkTest]int
	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = i ^ 0xcccc
	}
	var data [LengthOfArrayAndSliceForBenchmarkTest]int

	for i := 0; i < b.N; i++ {
		for j, v := range unsorted {
			data[j] = v
		}

		b.StartTimer()

		Sort(&data, len(data), func(i, j int) bool {
			return data[i] > data[j]
		})

		b.StopTimer()
	}
}

func BenchmarkGortIntSlice(b *testing.B) {
	b.StopTimer()

	unsorted := make([]int, LengthOfArrayAndSliceForBenchmarkTest)
	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = i ^ 0xcccc
	}

	data := make([]int, LengthOfArrayAndSliceForBenchmarkTest)

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)

		b.StartTimer()

		Sort(&data, len(data), func(i, j int) bool {
			return data[i] > data[j]
		})

		b.StopTimer()
	}
}

func BenchmarkGortStructArray(b *testing.B) {
	b.StopTimer()

	type PeopleArray [LengthOfArrayAndSliceForBenchmarkTest]People

	var unsorted PeopleArray
	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = People{
			Name: "",
			Age:  i ^ 0xcccc,
		}
	}

	var data PeopleArray

	for i := 0; i < b.N; i++ {
		for j, v := range unsorted {
			data[j] = v
		}

		b.StartTimer()

		Sort(&data, len(data), func(i, j int) bool {
			return data[i].Age > data[j].Age
		})
		b.StopTimer()
	}
}

func BenchmarkGortStructSlice(b *testing.B) {
	b.StopTimer()

	type PeopleSlice []People

	unsorted := make(PeopleSlice, LengthOfArrayAndSliceForBenchmarkTest)
	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = People{
			Name: "",
			Age:  i ^ 0xcccc,
		}
	}

	data := make(PeopleSlice, LengthOfArrayAndSliceForBenchmarkTest)

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)

		b.StartTimer()

		Sort(&data, len(data), func(i, j int) bool {
			return data[i].Age > data[j].Age
		})

		b.StopTimer()
	}
}

func BenchmarkDefaultSortIntSlice(b *testing.B) {
	b.StopTimer()

	unsorted := make([]int, LengthOfArrayAndSliceForBenchmarkTest)
	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = i ^ 0xcccc
	}

	data := make([]int, LengthOfArrayAndSliceForBenchmarkTest)

	for i := 0; i < b.N; i++ {
		for j, v := range unsorted {
			data[j] = v
		}

		b.StartTimer()

		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})

		b.StopTimer()
	}
}

func BenchmarkDefaultSortStructSlice(b *testing.B) {
	b.StopTimer()

	type PeopleSlice []People

	unsorted := make(PeopleSlice, LengthOfArrayAndSliceForBenchmarkTest)
	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = People{
			Name: "",
			Age:  i ^ 0xcccc,
		}
	}

	data := make(PeopleSlice, LengthOfArrayAndSliceForBenchmarkTest)

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)

		b.StartTimer()

		sort.Slice(data, func(i, j int) bool {
			return data[i].Age > data[j].Age
		})

		b.StopTimer()
	}
}
