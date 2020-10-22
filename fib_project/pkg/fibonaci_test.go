package fibonaci

import "testing"

func TestFibo(t *testing.T) {
	got := Fibo(7)
	want := 13
	if got != want {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}
