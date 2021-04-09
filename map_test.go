package functional

import (
	"fmt"
	"testing"
)

func inc(i int) int {
	return i + 1
}

func add(a, b int) int {
	return a + b
}

func even(i int) bool {
	return i%2 == 0
}

func TestStuff(t *testing.T) {
	v := Reduce(Mapping(inc)(Filtering(even)(add)), 0, []int{1, 1, 1, 2})
	if v != 6 {
		t.Errorf("Reduce(): got %v, want %v", v, 6)
	}

	c := Comp(Mapping(inc), Filtering(even), Taking(3))
	v = Reduce(c(add), 0, []int{2, 1, 1, 1, 1, 1, 1, 2})
	if v != 6 {
		t.Errorf("Reduce(): got %v, want %v", v, 6)
	}
	plus100 := func(i int) int { return i + 100 }
	myPrint := func(i int) string { return fmt.Sprintf("%#x", i) }
	myAppend := func(ss []string, s string) []string { return append(ss, s) }
	v = Reduce(Mapping(plus100)(Mapping(myPrint)(myAppend)), []string{}, []int{1, 2, 3, 4, 5})
	t.Log(v)
}
