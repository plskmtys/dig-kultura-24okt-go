package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Time struct {
	hour int
	min  int
}

func (t Time) cmp(other Time) int {
	return cmp.Compare(t.hour*60+t.min, other.hour*60+other.min)
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

func (t Time) toString() string {
	return fmt.Sprintf("%02d:%02d", t.hour, t.min)
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
	return fmt.Sprintf("time: %s, speed: %d", m.time.toString(), m.speed)
}

type Car struct {
	idx  int
	lp   string
	data []Measurement
}

func (c Car) print() {
	fmt.Printf("Car: %s, index: %d, Measurements: %d\n", c.lp, c.idx, len(c.data))

	for _, msrmnt := range c.data {
		fmt.Println(msrmnt.toString())
	}
}

func (c Car) toString() string {
	return fmt.Sprintf("Car: %s, Measurements: %d, Number: %d\n", c.lp, len(c.data), c.idx)
}

func parseLine(idx *int, line string, cars *[]Car) {
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

	*cars = append(*cars, Car{*idx, lp, []Measurement{measurement}})
	*idx++

}

func main() {

	fmt.Println("\n1. feladat")
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
	measurement_num := 0
	car_num := 0

	for _, line := range flines {
		parseLine(&car_num, line, &cars)
		measurement_num++
	}

	slices.SortFunc(cars, func(a, b Car) int { return cmp.Compare(a.lp, b.lp) })

	for _, car := range cars {
		fmt.Println(car.toString())
	}

	fmt.Printf("number of measurements: %d\n", measurement_num)
	fmt.Printf("number of cars: %d\n", len(cars))

	fmt.Println("\n2. feladat")

	max_time := Time{0, 0}
	max_time_lp := ""
	for _, car := range cars {
		for _, meas := range car.data {
			if meas.time.cmp(max_time) > 0 {
				max_time = meas.time
				max_time_lp = car.lp
			}
		}
	}

	fmt.Printf("legutolsó jeladás: %s, %s", max_time.toString(), max_time_lp)

	fmt.Println("\n3. feladat")
	for _, car := range cars {
		if car.idx == 0 {
			car.print()
		}
	}

	fmt.Println("\n4. feladat")
	//TODO user input...

	fmt.Println("\n5. feladat")
	max_speed := 0
	lps := []string{}
	for _, car := range cars {
		for _, meas := range car.data {
			if meas.speed > max_speed {
				max_speed = meas.speed
				lps = []string{}
				lps = append(lps, car.lp)
			} else if meas.speed == max_speed {
				lps = append(lps, car.lp)
			}
		}
	}

	fmt.Printf("max_speed: %v\n", max_speed)
	for _, str := range lps {
		fmt.Printf("%v ", str)
	}

	fmt.Println("\n6. feladat")
	//TODO user input...

	fmt.Println("\n7. feladat")
	//TODO file write...
}
