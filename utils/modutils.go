package utils

type modusLib struct {
	MODUS int
}

type ModusLib interface {
	GetModus() int
	Add(a, b int) int
	Sub(a, b int) int
	Mul(a, b int) int
	Inverse(a int) int
	Exp(a, b int) int
	Fact(n int) int
	Facts(n int) []int
	Fibo(n int) []int
	Nck(n, k int) int
}

func NewModusLib(modus int) ModusLib {
	return &modusLib{MODUS: modus}
}

func (m *modusLib) GetModus() int {
	return m.MODUS
}

func (m *modusLib) Add(a, b int) int {
	return (a + b) % m.MODUS
}

func (m *modusLib) Sub(a, b int) int {
	result := (a - b) % m.MODUS
	if result < 0 {
		result = result + m.MODUS
	}
	return result
}

func (m *modusLib) Mul(a, b int) int {
	a = a % m.MODUS
	b = b % m.MODUS
	result := (a * b) % m.MODUS
	return result
}

func (m *modusLib) Inverse(a int) int {
	divisor := a
	quotient := m.MODUS / a
	reminder := m.MODUS % a
	s1, s2, t1, t2 := 1, 0, 0, 1
	s3, t3 := s1-quotient*s2, t1-quotient*t2
	for reminder != 0 {
		//fmt.Println(divisor, quotient, reminder, s2, s3, t2, t3)
		quotient, divisor, reminder = divisor/reminder, reminder, divisor%reminder
		s1, s2, s3 = s2, s3, s2-quotient*s3
		t1, t2, t3 = t2, t3, t2-quotient*t3
	}
	// fmt.Println(divisor, quotient, reminder, s2, s3, t2, t3)
	result := t2 % m.MODUS
	if result < 0 {
		return result + m.MODUS
	} else {
		return result
	}
}

// Exp computes a raised to power b
func (m *modusLib) Exp(a, b int) int {
	result := 1
	for b > 0 {
		if b%2 == 1 {
			result = (result * a) % m.MODUS
		}
		b = b >> 1
		a = (a * a) % m.MODUS
	}
	return result
}

// Fact computes factorial of n
func (m *modusLib) Fact(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result = (result * i) % m.MODUS
	}
	return result
}

// Facts computes factorials (series) up to n.
// assumes n >= 0
func (m *modusLib) Facts(n int) []int {
	fact := make([]int, n+1)
	fact[1] = 1
	for i := 2; i <= n; i++ {
		fact[i] = (fact[i-1] * i) % m.MODUS
	}
	return fact
}

// Fibo computes fibonacci series up to n.
// assumes n >= 2
func (m *modusLib) Fibo(n int) []int {
	fib := make([]int, n+1)
	if n <= 0 {
		return fib
	}
	fib[1] = 1
	if n == 1 {
		return fib
	}
	fib[2] = 1
	if n == 2 {
		return fib
	}
	for i := 3; i <= n; i++ {
		fib[i] = (fib[i-2] + fib[i-1]) % m.MODUS
	}
	return fib
}

// Nck compute combinations for n choose k
func (m *modusLib) Nck(n, k int) int {
	result := 1
	for i := n; i > n-k; i-- {
		result = (result * i) % m.MODUS
	}
	result = m.Mul(result, m.Inverse(m.Fact(k)))
	return result
}
