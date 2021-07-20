package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
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

	combinations(6)

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

func combinations(r int) {
	totalSlips := math.Pow(3, float64(r))

	fmt.Println(int(totalSlips) / r)

	var results [][]string

	var gameOne []string
	var gameTwo []string

	for i := 0; i < r; i++ {
		if i == 0 {
			b := []string{"1", "X", "2"}
			for i := 0; i < int(totalSlips)/3; i++ {
				gameOne = append(gameOne, b...)
			}
		}
		if i == 1 {
			b := []string{"1", "X", "2", "X", "1", "1", "2", "2", "X"}
			for i := 0; i < int(totalSlips)/9; i++ {
				gameTwo = append(gameTwo, b...)
			}
		}

		var games []string

		if i >= 2 {

			var pow = []string{"1", "X", "2"}

			for _, v := range pow {
				position := (i + 1)
				// 	// fmt.Println(position)
				s := math.Pow(3, float64(position))
				iterations := (s / 3)
				// 	fmt.Println(iterations)

				var game []string
				for i := 0; i < int(iterations); i++ {
					// fmt.Println("**********", v)
					game = append(game, v)
				}
				// fmt.Println(game)
				games = append(games, game...)
			}
			

			if int(totalSlips) > len(games) {
				// fmt.Println(int(totalSlips) / len(games))
				var g []string
				for i := 0; i < int(totalSlips)/len(games); i++ {
					g = append(g, games...)
				}
				// fmt.Println(g)
				results = append(results, g)
			}

		}
	}

	// fmt.Println(gameOne)
	results = append(results, gameOne)
	// fmt.Println(gameTwo)
	results = append(results, gameTwo)

	fmt.Println(results)

	writeFile("combinations.csv", results)
}
