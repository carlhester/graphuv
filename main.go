package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	chart "github.com/wcharczuk/go-chart"
)

func drawChart(times []time.Time, vals []float64) int {

	graph := chart.Chart{
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeMinuteValueFormatter,
		},

		Series: []chart.Series{
			chart.TimeSeries{
				XValues: times,
				YValues: vals,
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
	return 0
}

func main() {
	lines := readLinesFromFile()
	times, vals := linesToTimeValues(lines)
	result := drawChart(times, vals)
	os.Exit(result)
}

func readLinesFromFile() []string {

	f, err := os.Open("./testfile.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func linesToTimeValues(lines []string) ([]time.Time, []float64) {

	var times []time.Time
	var measures []float64

	for _, v := range lines {
		vals := strings.Split(v, ",")
		times = append(times, convertTimeFormat(string(vals[0])))
		floatval, _ := strconv.ParseFloat(vals[1], 64)
		measures = append(measures, floatval)
	}

	return times, measures
}

func convertTimeFormat(value string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, value)
	if err != nil {
		fmt.Println(err)
	}
	return t
}
