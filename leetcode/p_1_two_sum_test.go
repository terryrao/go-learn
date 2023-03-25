package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	expect := []int{2, 7}
	sum := 9
	result := twoSum(nums, sum)
	if b := reflect.DeepEqual(expect, result); !b {
		t.Fatalf("not equal")
	}
}
func TestTwoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	expect := []int{0, 1}
	sum := 9
	result := twoSum2(nums, sum)
	if b := reflect.DeepEqual(expect, result); !b {
		t.Fatalf("not equal")
	}
}
