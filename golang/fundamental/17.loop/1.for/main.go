package main

import "fmt"

func main() {
	var counter int = 1

	for counter <= 10 {
		fmt.Println("loop ke", counter)

		counter++
		/*
			output:
			jika dicomment counter++ nya outputnya akanloop ke 1 dan terjadi infinite loop, dikarenakan tidak ada kondisi membuat loop itu berhenti
			jika tidak dicomment maka counter akan terus bertambah 1 nilai sampe kondisi loop berhenti karena kondisi tersebut kita set <= 10
		*/
	}
}
