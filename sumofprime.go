package piscine

func SumOfPrime(n int) int {

	if n < 2 {
		return 0
	}

	sum := 2
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum = sum + i
		}
	}
	return sum
}

func isPrime(p int) bool {
	if p < 2 {
		return false
	}

	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}
