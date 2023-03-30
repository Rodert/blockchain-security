package utils

import "math"

func DistinctBy[T any, K comparable](list []T, selector func(T) K) []T {
	result := []T{}
	set := map[K]struct{}{}
	for i := range list {
		key := selector(list[i])
		if _, ok := set[key]; !ok {
			set[key] = struct{}{}
			result = append(result, list[i])
		}
	}
	return result
}

func Map[T any, R any](list []T, transform func(T) R) []R {
	result := make([]R, 0, len(list))
	for i := range list {
		result = append(result, transform(list[i]))
	}
	return result
}

func MapIndexed[T any, R any](list []T, transform func(int, T) R) []R {
	result := make([]R, 0, len(list))
	for i := range list {
		result = append(result, transform(i, list[i]))
	}
	return result
}

func MapHasErr[T any, R any](list []T, transform func(T) (R, error)) ([]R, error) {
	result := make([]R, 0, len(list))
	for i := range list {
		ret, err := transform(list[i])
		if err != nil {
			return nil, err
		}
		result = append(result, ret)
	}
	return result, nil
}

func MapNotNil[T any, R any](list []T, transform func(T) *R) []R {
	result := make([]R, 0, len(list))
	for i := range list {
		ret := transform(list[i])
		if ret != nil {
			result = append(result, *ret)
		}
	}
	return result
}

func MapNotNilHasErr[T any, R any](list []T, transform func(T) (*R, error)) ([]R, error) {
	result := make([]R, 0, len(list))
	for i := range list {
		ret, err := transform(list[i])
		if err != nil {
			return nil, err
		}
		if ret != nil {
			result = append(result, *ret)
		}
	}
	return result, nil
}

func AssociateBy[T any, K comparable](list []T, keySelector func(T) K) map[K]T {
	m := map[K]T{}
	for i := range list {
		m[keySelector(list[i])] = list[i]
	}
	return m
}

func Associate[T any, K comparable, V any](list []T, transform func(t T) (K, V)) map[K]V {
	m := map[K]V{}
	for i := range list {
		k, v := transform(list[i])
		m[k] = v
	}
	return m
}

func SliceContains[T comparable](list []T, t T) bool {
	for i := range list {
		if list[i] == t {
			return true
		}
	}
	return false
}

func FlatMap[T any, R any](list []T, transform func(T) ([]R, error)) ([]R, error) {
	result := []R{}
	for i := range list {
		ret, err := transform(list[i])
		if err != nil {
			return nil, err
		}
		result = append(result, ret...)
	}
	return result, nil
}

func FlatMapPtr[T any, R any](list []T, transform func(T) ([]*R, error)) ([]*R, error) {
	result := []*R{}
	for i := range list {
		ret, err := transform(list[i])
		if err != nil {
			return nil, err
		}
		result = append(result, ret...)
	}
	return result, nil
}

// 不包含 string 的 + 运算符类型约束。
type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~complex64 | ~complex128
}

func Reduce[S Addable, T any](list []T, operation func(S, T) S) S {
	var accumulator = S(0)
	for i := range list {
		accumulator = operation(accumulator, list[i])
	}
	return accumulator
}

func GroupBy[T any, K comparable](list []T, keySelector func(T) K) map[K][]T {
	destination := make(map[K][]T, len(list))
	for i := range list {
		key := keySelector(list[i])
		if _, ok := destination[key]; ok {
			destination[key] = append(destination[key], list[i])
		} else {
			destination[key] = []T{list[i]}
		}
	}
	return destination
}

func GroupByTransform[T any, K comparable, V any](
	list []T,
	keySelector func(T) K,
	valueTransform func(T) V,
) map[K][]V {
	destination := make(map[K][]V, len(list))
	for i := range list {
		key := keySelector(list[i])
		if _, ok := destination[key]; ok {
			destination[key] = append(destination[key], valueTransform(list[i]))
		} else {
			destination[key] = []V{valueTransform(list[i])}
		}
	}
	return destination
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func SliceRange[T Integer](rs ...T) []T {
	destination := []T{}
	if len(rs) == 1 {
		if rs[0] < 0 {
			return destination
		}
		for i := T(0); i <= rs[0]; i++ {
			destination = append(destination, i)
		}
	} else if len(rs) == 2 {
		if rs[0] < 0 {
			return destination
		}
		if rs[1] < 0 {
			return destination
		}
		for i := rs[0]; i <= rs[1]; i++ {
			destination = append(destination, i)
		}
	}

	return destination
}

// retrieve num elements from position fromIdx in array rs, rs is required when run
func SliceSubRange[T any](list []T, fromIdx int, num int) []T {
	if fromIdx < 0 || num < 0 || fromIdx >= len(list) {
		return []T{}
	}
	return list[fromIdx:int(math.Min(float64(fromIdx+num), float64(len(list))))]
}

func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			nn = append(nn, v)
		}
	}
	return nn
}
