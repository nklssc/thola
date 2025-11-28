package utility

// IfThenElse is a generic wrapper for the if condition.
func IfThenElse[T any](condition bool, t, e T) T {
	if condition {
		return t
	}
	return e
}

// SliceUnique removes duplicates from a slice of comparable types.
func SliceUnique[T comparable](a []T) []T {
	l := len(a)
	seen := make(map[T]struct{}, l)
	k := 0

	for i := 0; i < l; i++ {
		v := a[i]
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		a[k] = v
		k++
	}

	return a[0:k]
}

func SliceContains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func SameSlice[T comparable](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of T -> int
	diff := make(map[T]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the value
		diff[_x]++
	}
	for _, _y := range y {
		// If the value _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}
