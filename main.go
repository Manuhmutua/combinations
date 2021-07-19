package main

import (
	"encoding/csv"
	"fmt"
	"log"
	// "math"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("eurocup_2020_results.csv")

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var clean [][]string
	var clean2 [][]string

	for _, record := range records {

		var s []string
		var s2 []string

		var entry string

		if record[7] > record[8] {
			entry = "0"
		} else if record[8] > record[7] {
			entry = "2"
		} else if record[8] == record[7] {
			entry = "1"
		}

		s = append(s, trimString(record[5]), trimString(record[6]), trimString(record[7]), trimString(record[8]), trimString(entry))

		s2 = append(s2, trimString(record[5])+" "+trimString(record[6]), trimString(entry))

		clean = append(clean, s)
		clean2 = append(clean2, s2)
	}

	fmt.Println(clean)
	fmt.Println(clean2)

	writeFile("result.csv", clean)

	writeFile("result2.csv", clean2)

	fmt.Printf("iterable = %s, r = %d", "[]int{1, 2, 3, 4, 5, 6}", 3)
	fmt.Println()

	combinations([]int{1, 2}, 9)

}

func writeFile(s string, w [][]string) {
	file, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range w {
		err := writer.Write(value)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func trimString(s string) string {
	res1 := strings.Trim(s, " ")
	return res1
}

func combinations(iterable []int, r int) {
	pool := iterable
	n := len(pool)

	if r > n {
		return
	}

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]int, r)
	for i, el := range indices {
		result[i] = pool[el]
	}

	fmt.Println(result)

	for {
		i := r - 1
		for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
		}

		if i < 0 {
			return
		}

		indices[i] += 1
		for j := i + 1; j < r; j += 1 {
			indices[j] = indices[j-1] + 1
		}

		for ; i < len(indices); i += 1 {
			result[i] = pool[indices[i]]
		}
		fmt.Println(result)

	}
}