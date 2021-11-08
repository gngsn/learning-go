package arrays

/*
	param으로 arrays가 아닌 가변형인 slices를 설정해주어야 함.
	이유는 basic -> arrays, slices 참고
*/
func Sum(numbers []int) int {
    sum := 0

	/*
		for i := 0; i < 5; i++ {
			sum += numbers[i]
		}

	 	-> Refactor
	*/

	for _, number := range numbers {
        sum += number
    }
    return sum
}

func SumAll(numbersToSum ...[]int) []int {
	/* Refactor 1 - arrays -> slice
	
		lengthOfNumbers := len(numbersToSum)
		sums := make([]int, lengthOfNumbers)

		for i, numbers := range numbersToSum {
			sums[i] = Sum(numbers)
		}
	*/

	/* Refactor 2 - slice: append
		for i, numbers := range numbersToSum {
			sums[i] = Sum(numbers)
		}
	*/

	var sums []int

	for _, numbers := range numbersToSum {
        sums = append(sums, Sum(numbers))
    }

    return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
    var sums []int
	
    for _, numbers := range numbersToSum {
		// 에러 방지 -> panic: runtime error: slice bounds out of range
        if len(numbers) == 0 {
            sums = append(sums, 0)
        } else {
            tail := numbers[1:]
            sums = append(sums, Sum(tail))
        }
    }

    return sums
}