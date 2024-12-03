package utils

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {

	if len(integers) < 2 {
		return 1
	}

	a := integers[0]
	b := integers[1]

	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Abs(x int) int {
	return AbsDiff(x, 0)
}

func AbsDiff(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}
