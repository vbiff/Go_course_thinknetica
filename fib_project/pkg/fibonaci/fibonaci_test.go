package fibonaci

import "testing"

func TestNum(t *testing.T) {
	got := Num(7)
	want := 13
	if got != want {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}
