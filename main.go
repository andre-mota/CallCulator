package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"
)

// PayDur is a struct to use when handling call duration and its price
type PayDur struct {
	TotalDuration float64
	TotalPay      int
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// FileParser takes the data file and parses its data to a map
func FileParser(filepath string) map[string][]PayDur {
	callerCounter := make(map[string][]PayDur)

	// Load the file
	f, _ := os.Open(filepath)

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'
	for {
		r, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		// !!todo
		tsStart, err := time.Parse("15:04:05", r[0])
		handleErr(err)

		tsEnd, err := time.Parse("15:04:05", r[1])
		handleErr(err)

		dur := tsEnd.Sub(tsStart)
		price := 0

		if dur.Seconds() >= 300 {
			price = 25 + int(math.Trunc((dur.Seconds()-300)/60))*2

			if math.Remainder(dur.Seconds()-300, 60) != 0 {
				price += 2
			}

		} else {
			price = int(math.Trunc(dur.Seconds()/60)) * 5

			if math.Remainder(dur.Seconds(), 60) != 0 {
				price += 5
			}
		}

		callerCounter[r[2]] = append(callerCounter[r[2]], PayDur{TotalDuration: dur.Seconds(), TotalPay: price})
	}
	return callerCounter
}

// SumCalls takes a map with a colection of calls from every user and returns a new map with a total per caller
func SumCalls(callerCounter map[string][]PayDur) map[string]PayDur {
	totalCalls := make(map[string]PayDur)

	for c := range callerCounter {
		// Initialize vars to sum values per caller
		totalDur := 0.0
		totalPay := 0
		for k := range callerCounter[c] {
			// Sum vales per caller
			totalDur += callerCounter[c][k].TotalDuration
			totalPay += callerCounter[c][k].TotalPay
		}
		totalCalls[c] = PayDur{TotalDuration: totalDur, TotalPay: totalPay}
	}
	return totalCalls
}

// PayFormat takes a map with total daily data per caller and returns the total amout to pay
func PayFormat(price int64) string {
	return ""
}

func main() {
	// Create maps to facilitate calcs
	callerCounter := FileParser(os.Args[1])
	totalCalls := SumCalls(callerCounter)

	fmt.Println(callerCounter, totalCalls)

}
