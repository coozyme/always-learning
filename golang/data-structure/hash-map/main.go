package main

import "fmt"

func main() {
	var employees = map[string][]string{
		"0822212242": []string{"Andi", "10/1/2010"},
		"0822212222": []string{"Kua", "10/1/2010"},
		"0822321222": []string{"Loli", "10/1/2010"},
		"0822241222": []string{"Jii", "10/1/2010"},
	}

	key := make([]string, 0, len(employees))
	value := make([][]string, 0, len(employees))

	for keys, val := range employees {
		key = append(key, keys)
		value = append(value, val)
	}

	fmt.Println("k", key)
	fmt.Println("V", value)

	fmt.Println("0822212242", employees["0822212242"])

	// Add New
	employees["0822212473"] = []string{"Dian", "28/10/2025"}
	// Edit Value
	employees["0822321222"] = []string{"Inu", "28/10/2025"}
	// Delete
	delete(employees, "0822212222")

	fmt.Println("new", employees)

}
