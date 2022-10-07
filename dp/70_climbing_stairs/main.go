package _0_climbing_stairs

func climbStairs(n int) int {
	a := 1
	b := 1
	var c int
	for i := 0; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return a
}
