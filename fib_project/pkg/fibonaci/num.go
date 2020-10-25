package fibonaci

// Num - fibonaci counter
func Num(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Num(n-1) + Num(n-2)
}
