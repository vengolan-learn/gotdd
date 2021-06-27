package main

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lenOfNumbers := len(numbersToSum)
	sums := make([]int, lenOfNumbers)

	for i, numbers := range numbersToSum {

		sums[i] = Sum(numbers)
	}
	return sums
}

func main() {

}
