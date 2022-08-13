package codility

import "fmt"

func CyrcilcRotation(A []int, K int) []int {
	length := len(A)
	if length > 0 {
		if K > length {
			K = K % length
			fmt.Println("k", K)
		}
		fmt.Println("ds", length-K)
		fmt.Println("a", A[length-K:length])
		fmt.Println("a2", A[0:length-K])
		A = append(A[length-K:length], A[0:length-K]...)
	}
	return A
}
