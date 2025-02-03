package services

import (
	"math"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsArmstrong(n int) bool {
	sum, temp := 0, n
	digits := int(math.Log10(float64(n)) + 1)

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	return sum == n
}

func DigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func GetNumberProperties(num int) []string {
	properties := []string{}

	if IsArmstrong(num) {
		properties = append(properties, "armstrong")
	}

	if num%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	return properties
}
