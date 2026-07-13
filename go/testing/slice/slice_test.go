package main

import (
	"fmt"
	"testing"
)

func appendSlice(s []int, n int) {
	s = append(s, n)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
func TestSliceAppend(t *testing.T) {
	s1 := make([]int, 5)
	appendSlice(s1, 10)
	if len(s1) != 5 || cap(s1) != 5 {
		t.Errorf("Expected len=5, cap=5; got len=%d, cap=%d", len(s1), cap(s1))
	}
}
