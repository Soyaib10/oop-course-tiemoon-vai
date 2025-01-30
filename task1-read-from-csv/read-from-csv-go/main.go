package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// read the csv
	file, err := os.Open("salarysheet.csv")
	if err != nil {
		fmt.Println("error in file opening: ", err)
		return
	}
	defer file.Close()

	// read the file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll() // 2d array of string
	if err != nil {
		log.Fatal("failed reading csv: ", err)
	}

	// all
	var (
		maxSalary, minSalary, totalSalary int
		minSalaryPerson, maxSalaryPerson  string
		cnt                             int
	)

	minSalary = math.MaxInt32
	for _, x := range records {
		name := x[0]
		salary, err := strconv.Atoi(x[1])
		if err != nil {
			log.Fatal("Failed to parse salary: ", err)
		}

		// find min salary
		if salary < minSalary {
			minSalary = salary
			minSalaryPerson = name
		}

		// find max salary
		if salary > maxSalary {
			maxSalary = salary
			maxSalaryPerson = name
		}

		// find total and count the number of person
		totalSalary += salary
		cnt++
	}

	avgSalary := totalSalary / cnt
	fmt.Printf("Maximum Salary: %s - %d\n", maxSalaryPerson, maxSalary)
	fmt.Printf("Minimum Salary: %s - %d\n", minSalaryPerson, minSalary)
	fmt.Printf("Average Salary: %d\n", avgSalary)
}
