package example

import (
	"fmt"
	"sort"
)

type Ordered interface {
	~int | ~float64 | ~string
}

// Group of functions that ensure that an OrderedSlice can be sorted
type OrderedSlice[T Ordered] []T // T must implement < and >

func (s OrderedSlice[T]) Len() int {
	return len(s)
}

func (s OrderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s OrderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Group of functions that ensure that SortType can be sorted
type SortType[T any] struct {
	slice   []T
	compare func(T, T) bool
}

func (s SortType[T]) Len() int {
	return len(s.slice)
}

func (s SortType[T]) Less(i, j int) bool {
	return s.compare(s.slice[i], s.slice[j])
}

func (s SortType[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func PerformSort[T any](slice []T, compare func(T, T) bool) {
	sort.Sort(SortType[T]{slice, compare})
}

type Student struct {
	Name string
	ID   int
	Age  float64
}

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

func TestOrderd() {
	students1 := []int{}
	result1 := addStudent(students1, 78)
	result1 = addStudent(result1, 64)
	result1 = addStudent(result1, 45)
	sort.Sort(OrderedSlice[int](result1))
	fmt.Println(result1)

	students2 := []Student{}
	result2 := addStudent(students2, Student{"John", 213, 17.5})
	result2 = addStudent(result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})
	// PerformSort(result2, func(s1, s2 Student) bool {
	// 	return s1.Age < s2.Age // comparing two Student values
	// })
	sort.Sort(SortType[Student]{students2, func(s1, s2 Student) bool {
		return s1.Age < s2.Age
	}})
	fmt.Println(result2)
}
