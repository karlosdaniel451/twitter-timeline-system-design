package utils

import "fmt"

func MapStringer[T fmt.Stringer](values []T) []string {
	mappedValues := make([]string, 0, len(values))
	for _, value := range values {
		mappedValues = append(mappedValues, value.String())
	}

	return mappedValues
}

func Map[T any](values []T, mapper func(T) T) []T {
	mappedValues := make([]T, 0, len(values))
	for _, value := range values {
		mappedValues = append(mappedValues, mapper(value))
	}

	return mappedValues
}

func ValueToPointer[T any](t T) *T {
	return &t
}
