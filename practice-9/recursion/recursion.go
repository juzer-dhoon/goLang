package recursion

func Fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * Fact(num-1)
}

func SumOfDigit(num int) int {
	if num == 0 {
		return 0
	}
	return SumOfDigit(num/10) + (num % 10)
}

func RevString(str string) string {
	if len(str) == 0 {
		return ""
	}
	return RevString(str[1:]) + string(str[0])
}
