package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

// PayDur is a struct to use when handling call duration and its price
type PayDur struct {
	TotalDuration float64
	TotalPay      Money
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// Money is a type to handle money format
type Money int

// FileParser takes the data file and parses its data to a map
func FileParser(filepath string) map[string][]PayDur {
	callerCounter := make(map[string][]PayDur)

	// Load the file
	f, err := os.Open(filepath)
	handleErr(err)

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
		var price Money

		if dur.Seconds() >= 300 {
			price = 25 + Money(math.Trunc((dur.Seconds()-300)/60))*2

			if math.Remainder(dur.Seconds()-300, 60) != 0 {
				price += 2
			}

		} else {
			price = Money(math.Trunc(dur.Seconds()/60)) * 5

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
		var totalDur float64
		var totalPay Money
		for k := range callerCounter[c] {
			// Sum vales per caller
			totalDur += callerCounter[c][k].TotalDuration
			totalPay += callerCounter[c][k].TotalPay
		}
		totalCalls[c] = PayDur{TotalDuration: totalDur, TotalPay: totalPay}
	}
	return totalCalls
}

// TopCaller finds the top caller and its amount to pay
func TopCaller(totalCalls map[string]PayDur) string {
	var topCaller string
	maxDur := 0.0
	for k := range totalCalls {
		if totalCalls[k].TotalDuration > maxDur {
			maxDur = totalCalls[k].TotalDuration
			topCaller = k
		}
	}
	return topCaller
}

// TotalDayPay takes a call colection, top caller and returns the correct amount to pay.
func TotalDayPay(totalCalls map[string]PayDur, topCaller string) Money {
	var p Money
	for k := range totalCalls {
		if k == topCaller {
			continue
		}
		p += totalCalls[k].TotalPay
	}
	return p
}

func (money Money) String() string {
	a := money / 100
	b := money % 100

	if b < 0 {
		b *= -1
	}

	return fmt.Sprintf("%d.%02d", a, b)
}

func exec(filepath string) Money {
	// Create maps to facilitate calcs
	callerCounter := FileParser(filepath)
	totalCalls := SumCalls(callerCounter)
	topCaller := TopCaller(totalCalls)
	totalDayPay := TotalDayPay(totalCalls, topCaller)
	return totalDayPay
}
func main() {
	fmt.Println(exec(os.Args[1]))

}
