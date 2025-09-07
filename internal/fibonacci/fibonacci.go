package fibonacci

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	anterior, atual := 0, 1

	for i := 2; i <= n; i++ {
		anterior, atual = atual, atual+anterior
	}

	return atual
}
