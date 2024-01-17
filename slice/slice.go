package sliceUtils

import (
	"math/rand"
	"strings"
)

func TakeRandomElement[T any](arr []T) (e T) {
	if len(arr) == 0 {
		return
	}
	return arr[rand.Intn(len(arr))]
}

func Contain[T comparable](a T, array []T) bool {
	for _, b := range array {
		if b == a {
			return true
		}
	}
	return false
}

func StrContainAnyElement(str string, array []string) bool {
	for _, b := range array {
		if strings.Contains(str, b) {
			return true
		}
	}
	return false
}

func StrContainAllElement(str string, array []string) bool {
	for _, b := range array {
		if !strings.Contains(str, b) {
			return false
		}
	}
	return true
}

func AnyElementContainStr(str string, array []string) bool {
	for _, b := range array {
		if strings.Contains(b, str) {
			return true
		}
	}
	return false
}

func AllElementContainStr(str string, array []string) bool {
	for _, b := range array {
		if !strings.Contains(b, str) {
			return false
		}
	}
	return true
}

func AllEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Distinct[T comparable](array []T) []T {
	keys := make(map[T]bool)
	newArr := []T{}
	for _, entry := range array {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			newArr = append(newArr, entry)
		}
	}
	return newArr
}

func SuffleNewSlice[T comparable](s []T) []T {
	var copy []T
	copy = append(copy, s...)
	rand.Shuffle(len(copy), func(i, j int) { copy[i], copy[j] = copy[j], copy[i] })
	return copy
}

func ShuffleSlice[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

func RandomIndexs(len int) []int {
	indexes := make([]int, len)
	for i := 0; i < len; i++ {
		indexes[i] = i
	}

	ShuffleSlice(indexes)
	return indexes
}

func IsSubset[T comparable](subset []T, superset []T) bool {
	checkset := make(map[T]bool)
	for _, element := range superset {
		checkset[element] = true
	}
	for _, value := range subset {
		if !checkset[value] {
			return false
		}
	}
	return true
}

func IsSubsetStr(subset []string, superset []string) bool {
	checkset := make(map[string]bool)
	for _, element := range superset {
		checkset[strings.ToLower(element)] = true
	}
	for _, value := range subset {
		if !checkset[strings.ToLower(value)] {
			return false
		}
	}
	return true
}

func IsSubsetStrContains(subset []string, superset []string) bool {

	for i, v := range subset {
		subset[i] = strings.ToLower(v)
	}

	for i, v := range superset {
		superset[i] = strings.ToLower(v)
	}

loop:
	for _, k := range subset {
		for _, s := range superset {
			if strings.Contains(s, k) {
				continue loop
			}
		}
		return false
	}
	return true
}

func Chunk[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func EqualFoldStr(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		if !Contain(v, b) {
			return false
		}
	}
	return true
}
