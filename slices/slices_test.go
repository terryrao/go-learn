package slices

import "testing"

func TestCopy(t *testing.T) {
	src := []string{"a", "b", "c"}
	dest := Copy(src)

}
