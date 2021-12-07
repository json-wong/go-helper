package helper

import (
	"math"
	"reflect"
)

// InArray in_array()
// haystack supported types: slice, array or map
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type must be slice, array or map")
	}

	return false
}

// ArrayKeys array_keys()
func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key := range elements {
		keys[i] = key
		i++
	}
	return keys
}

// ArrayValues array_values()
func ArrayValues(elements map[interface{}]interface{}) []interface{} {
	i, vals := 0, make([]interface{}, len(elements))
	for _, val := range elements {
		vals[i] = val
		i++
	}
	return vals
}

// ArrayKeyExists array_key_exists()
func ArrayKeyExists(key interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[key]
	return ok
}

// ListIntersect — Computes the intersection of lists
func ListIntersect(nums1, nums2 []int) []int {

	res := make([]int, 0, len(nums1))
	nc := make(map[int]int)

	for _, n := range nums1 {
		nc[n]++
	}

	for _, n := range nums2 {
		if nc[n] > 0 {
			res = append(res, n)
			nc[n]--
		}
	}

	return res
}

// ListMerge — Merge one or more lists
func ListMerge(arr ...[]interface{}) []interface{} {
	n := 0
	for _, v := range arr {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range arr {
		s = append(s, v...)
	}
	return s
}

// ListReverse - Return a list with elements in reverse order
func ListReverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// ListChunk array_chunk()
func ListChunk(s []interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("size: cannot be less than 1")
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}

// ListSlice array_slice()
func ListSlice(s []interface{}, offset, length uint) []interface{} {
	if offset > uint(len(s)) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < uint(len(s)) {
		return s[offset:end]
	}
	return s[offset:]
}

// ListPop array_pop()
// Pop the element off the end of slice
func ListPop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// ListUnshift array_unshift()
// Prepend one or more elements to the beginning of a slice
func ListUnshift(s *[]interface{}, elements ...interface{}) int {
	*s = append(elements, *s...)
	return len(*s)
}

// ListShift array_shift()
// Shift an element off the beginning of slice
func ListShift(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}
