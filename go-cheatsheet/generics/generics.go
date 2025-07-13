package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// MapKeys is a generic function that works on any map type.
// K must be `comparable` (can be used as a map key). V can be `any` type.
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Number is a custom constraint that includes integers and floats.
// This is an interface that acts as a type set.
type Number interface {
	constraints.Integer | constraints.Float
}

// SumNumbers is a generic function that uses the Number constraint.
func SumNumbers[T Number](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

// DemonstrateGenerics shows generic functions in action.
func DemonstrateGenerics() {
	fmt.Println("\n[Generics]")

	intMap := map[int]string{1: "a", 2: "b"}
	stringMap := map[string]int{"a": 1, "b": 2}

	fmt.Println("Keys from int map:", MapKeys(intMap))
	fmt.Println("Keys from string map:", MapKeys(stringMap))

	ints := []int{1, 2, 3}
	floats := []float64{1.1, 2.2, 3.3}

	fmt.Println("Sum of ints:", SumNumbers(ints))
	fmt.Println("Sum of floats:", SumNumbers(floats))
}
