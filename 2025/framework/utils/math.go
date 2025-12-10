package utils

import "math"

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

func IntLen(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

func BitCount(i int) int {
	count := 0
	for i > 0 {
		if i%2 != 0 {
			count += 1
		}
		i /= 2
	}
	return count
}

func IntPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
