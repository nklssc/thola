package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfThenElse(t *testing.T) {
	// Test with bool
	result := IfThenElse(true, true, false)
	assert.Equal(t, true, result, "IfThenElse(true, true, false) should return true")

	result = IfThenElse(false, true, false)
	assert.Equal(t, false, result, "IfThenElse(false, true, false) should return false")

	// Test with int
	resultInt := IfThenElse(true, 1, 2)
	assert.Equal(t, 1, resultInt, "IfThenElse(true, 1, 2) should return 1")

	resultInt = IfThenElse(false, 1, 2)
	assert.Equal(t, 2, resultInt, "IfThenElse(false, 1, 2) should return 2")

	// Test with string
	resultStr := IfThenElse(true, "yes", "no")
	assert.Equal(t, "yes", resultStr, "IfThenElse(true, \"yes\", \"no\") should return \"yes\"")

	resultStr = IfThenElse(false, "yes", "no")
	assert.Equal(t, "no", resultStr, "IfThenElse(false, \"yes\", \"no\") should return \"no\"")
}

func TestSliceUnique(t *testing.T) {
	// Test with int slice
	input := []int{1, 2, 2, 3, 1, 4}
	result := SliceUnique(input)
	expected := []int{1, 2, 3, 4}
	assert.Equal(t, expected, result, "SliceUnique should remove duplicates from int slice")

	// Test with string slice
	inputStr := []string{"a", "b", "b", "c", "a"}
	resultStr := SliceUnique(inputStr)
	expectedStr := []string{"a", "b", "c"}
	assert.Equal(t, expectedStr, resultStr, "SliceUnique should remove duplicates from string slice")

	// Test with empty slice
	empty := []int{}
	resultEmpty := SliceUnique(empty)
	assert.Len(t, resultEmpty, 0, "SliceUnique should return empty slice for empty input")

	// Test with no duplicates
	noDup := []int{1, 2, 3}
	resultNoDup := SliceUnique(noDup)
	assert.Equal(t, noDup, resultNoDup, "SliceUnique should return same slice when no duplicates")
}

func TestSliceContains(t *testing.T) {
	// Test with string slice
	slice := []string{"apple", "banana", "cherry"}
	assert.True(t, SliceContains(slice, "banana"), "SliceContains should return true for existing string element")
	assert.False(t, SliceContains(slice, "grape"), "SliceContains should return false for non-existing string element")

	// Test with int slice
	intSlice := []int{1, 2, 3, 4}
	assert.True(t, SliceContains(intSlice, 3), "SliceContains should return true for existing int element")
	assert.False(t, SliceContains(intSlice, 5), "SliceContains should return false for non-existing int element")

	// Test with empty slice
	empty := []string{}
	assert.False(t, SliceContains(empty, "anything"), "SliceContains should return false for empty slice")
}

func TestSameSlice(t *testing.T) {
	// Test identical string slices
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	assert.True(t, SameSlice(a, b), "SameSlice should return true for identical string slices")

	// Test same string elements, different order
	c := []string{"c", "a", "b"}
	assert.True(t, SameSlice(a, c), "SameSlice should return true for same string elements in different order")

	// Test different string elements
	d := []string{"a", "b", "d"}
	assert.False(t, SameSlice(a, d), "SameSlice should return false for different string elements")

	// Test different lengths
	e := []string{"a", "b"}
	assert.False(t, SameSlice(a, e), "SameSlice should return false for different lengths")

	// Test empty slices
	empty1 := []string{}
	empty2 := []string{}
	assert.True(t, SameSlice(empty1, empty2), "SameSlice should return true for empty slices")

	// Test with int slices
	intA := []int{1, 2, 3}
	intB := []int{3, 1, 2}
	assert.True(t, SameSlice(intA, intB), "SameSlice should return true for same int elements in different order")

	intC := []int{1, 2, 4}
	assert.False(t, SameSlice(intA, intC), "SameSlice should return false for different int elements")
}
