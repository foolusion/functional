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

	c := Comp(Taking(3), Filtering(even), Mapping(inc))
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

func Range(n int) []int {
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, i)
	}
	return out
}

var range1000 = Range(1000)

func BenchmarkStuff(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(Mapping(inc)(add), 0, range1000)
	}
}
