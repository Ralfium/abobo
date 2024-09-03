package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = ';'

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	numConvA, _ := strconv.ParseFloat(records[0][0], 64)

	numConvB, _ := strconv.ParseFloat(records[0][1], 64)
	res := numConvA / numConvB
	fmt.Printf("%.4f", res)
}
