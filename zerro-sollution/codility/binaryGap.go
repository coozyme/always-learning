package codility

func BinaryGap(N int64) int {
	// var result int
	// var result_temp int
	// // var calc bool

	// for N > 0 {
	// 	fmt.Println("aa1", N)
	// 	fmt.Println("aa", N%2)
	// 	if N%2 == 1 {
	// 		// 	if !calc {
	// 		// 		calc = true
	// 		// 		fmt.Println("a", calc)
	// 		// 	} else {
	// 		if result_temp > result {
	// 			result = result_temp
	// 			fmt.Println("b", result)
	// 		}
	// 		result_temp = 0
	// 		fmt.Println("c", result_temp)
	// 		// }
	// 		// fmt.Println("if", result_temp)
	// 	} else {
	// 		// if calc {
	// 		result_temp++
	// 		fmt.Println("d", result_temp)
	// 		// }
	// 	}
	// 	result_temp++

	// 	fmt.Println("e", result_temp)
	// 	N /= 2
	// 	fmt.Println("n/2", N)
	// }

	// return result
	var result int
	var result_temp int
	var calc bool

	for N > 0 {
		if N%2 == 1 {
			if !calc {
				calc = true
			} else {
				if result_temp > result {
					result = result_temp
				}
				result_temp = 0
			}
		} else {
			if calc {
				result_temp++
			}
		}
		N = N / 2
	}

	return result
}
