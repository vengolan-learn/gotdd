package main

//Getthe sum of all values in the int array
func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

//get an array that stores the sum of each individual array
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

//Tail means  of all items excluding the first element of array
//This function returns the array of all tail elements of each array
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		var s int
		if len(numbers) > 0 {
			s = Sum(numbers[1:])
		}
		sums = append(sums, s)
	}
	return sums
}

func main() {

}
