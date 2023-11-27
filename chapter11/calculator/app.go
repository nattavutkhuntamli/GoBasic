package calculator

func Summation(numbers ...int) int {
	total :=0
	for _, value := range numbers {
		total+=value
	}
	return total
}