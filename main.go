package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Time struct {
	hour int
	min  int
}

func (t Time) fromstr(str1 string, str2 string) Time {
	hour, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("error reading hour value")
	}
	min, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("error reading hour value")
	}
	return Time{hour, min}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Measurement struct {
	time  Time
	speed int
}

func (m Measurement) fromstr(strs []string) Measurement {
	var time Time
	time = Time.fromstr(time, strs[0], strs[1])
	speed, err := strconv.Atoi(strs[2])
	if err != nil {
		fmt.Println("error converting string to measurement")
	}
	return Measurement{time, speed}
}

func (m Measurement) toString() string {
	return fmt.Sprintf("time: %02d:%02d, speed: %d", m.time.hour, m.time.min, m.speed)
}

type Car struct {
	lp   string
	data []Measurement
}

func (c Car) print() {
	fmt.Printf("Car: %s\n", c.lp)
	for _, msrmnt := range c.data {
		fmt.Println(msrmnt.toString())
	}
}

func parseLine(line string, cars *[]Car) {
	parts := strings.Split(line, "\t")
	lp := parts[0]
	var measurement Measurement
	measurement = Measurement.fromstr(measurement, parts[1:4])

	for i := range *cars {
		if (*cars)[i].lp == lp {
			(*cars)[i].data = append((*cars)[i].data, measurement)
			return
		}
	}

	*cars = append(*cars, Car{lp, []Measurement{measurement}})

}

func main() {

	fmt.Println("beolvasás")

	f, err := os.Open("jeladas.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	fscanner := bufio.NewScanner(f)
	fscanner.Split(bufio.ScanLines)

	var flines []string

	for fscanner.Scan() {
		flines = append(flines, fscanner.Text())
	}

	cars := []Car{}

	for _, line := range flines {
		parseLine(line, &cars)
	}

	for _, car := range cars {
		car.print()
	}

	fmt.Println("1. feladat")
}
