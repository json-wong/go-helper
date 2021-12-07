package helper

import (
	"math"
	"reflect"
	"testing"
)

// TestListIntersect
func TestListIntersect(t *testing.T) {
	n := []int{1, 2, 2, 3, 4, 5}
	m := []int{2, 2, 2, 3, 4, 5, 6, 7}

	t.Log(ListIntersect(n, m))
}

// TestMapKeys
func TestArrayKeys(t *testing.T) {
	input := map[interface{}]interface{}{"foo": 123, "bar": "abc"}
	expected := []interface{}{"foo", "bar"}
	output := ArrayKeys(input)
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("output: %v, expected: %v", output, expected)
	}

	t.Log(output)
}

func TestListMerge(t *testing.T) {
	t.Log(math.Ceil(232343.23232*100) / 100)
}
