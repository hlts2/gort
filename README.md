# gort
gort is a simple library for sorting slices and arrays.

## Requirement
Go 1.8+

## Installation

```go
go get github.com/hlts2/gort
```

## Example

### sorting int array or slice

```go

//when sorting int array
array := [4]int{-20, 2, 1, -10}

Sort(&array, len(array), func (i, j int) bool {
    return array[i] > array[j]
})

fmt.Println(array) //[2, 1, -10, -20]


//when sorting int slice
slice := []int{-20, 2, 1, -10}

Sort(&slice, len(slice), func (i, j int) bool {
    return slice[i] > slice[j]
})

fmt.Println(slice) //[2, 1, -10, -20]

```

### sorting struct array or slice

```go

//when sorting struct array
array := [4]People{
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

Sort(&array, len(array), func (i, j int) bool {
    return array[i].Age > array[j].Age
})

fmt.Println(array) //[People{Name: "b", Age:  30], People{Name: "d", Age:  3], People{Name: "a", Age:  2], People{Name: "c", Age:  1]]

//when sorting struct slice
slice := []People{
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

Sort(&slice, len(slice), func (i, j int) bool {
    return slice[i].Age > slice[j].Age
})

fmt.Println(slice) //[People{Name: "b", Age:  30], People{Name: "d", Age:  3], People{Name: "a", Age:  2], People{Name: "c", Age:  1]]

```

### sorting part of array or slice

```go

//when sorting int array
array := [4]int{-20, 2, 1, -10}

Sort(&array, 2, func (i, j int) bool {
    return array[i] > array[j]
})

fmt.Println(array) //[2, -20, 1, -10]


//when sorting int slice
slice := []int{-20, 2, 1, -10}

Sort(&slice, 3, func (i, j int) bool {
    return slice[i] > slice[j]
})

fmt.Println(slice) //[2, 1, -20, -10]
```
