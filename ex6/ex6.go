// calendar utility
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// globals
// ----------------------------------------------------------------------------
const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1

const version = "0.1"

var oneMonth, threeMonths bool
var months int
var sunday, monday bool
var blocks int

var want_version bool

// functions
// ----------------------------------------------------------------------------

// init module
//
// setup the flag environment for the on-line help
func init() {

	// command line arguments for parsing the number of months to show
	flag.BoolVar(&oneMonth, "1", true, "Display single month output")
	flag.BoolVar(&threeMonths, "3", false, "Display three months spanning the date")
	flag.IntVar(&months, "months", 0, "Display number of months, starting from the month containing the date")

	// command line arguments for parsing the first day of the week
	flag.BoolVar(&sunday, "sunday", false, "Display Sunday as the first day of the week")
	flag.BoolVar(&monday, "monday", true, "Display Monday as the irst day of the week")

	// command line argument for determining the number of months per block
	flag.IntVar(&blocks, "blocks", 3, "Set number of calendar sheet blocks")

	// also, create an additional flag for showing the version
	flag.BoolVar(&want_version, "version", false, "shows version info and exits")
}

// showVersion
//
// show the current version of this program and exits with the given signal
func showVersion(signal int) {

	fmt.Printf(" %v %v\n", os.Args[0], version)
	os.Exit(signal)
}

// min
//
// return the min of two ints
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// centerText
//
// returns a string which centers the given text within a field of the specified
// width
func centerText(text string, width int) string {

	// create a format string to specify the available space for writing the
	// text
	lftmargin := fmt.Sprintf("%%-%ds", width)
	rgtmargin := fmt.Sprintf("%%%ds", len(text)+(width-len(text))/2)
	return fmt.Sprintf(lftmargin,
		fmt.Sprintf(rgtmargin, text))
}

// span
//
// return the following values: prev, refdate, post with the following meaning
//
// prev: number of months to display previous to the start date
// refdate: reference date
// post: number of months to display from the reference date
//
// For example (-1, 01 01 2020, 2) shows three months starting in december 2019
// until march 2020 which is not shown
//
// the input parameters are:
//
// onemonth: display a single month
// threemonths: display three months spanning the reference date
// months: number of months to display
// ref: reference date given as a slice of strings in the format "dd mm yyyy"
//
// if the reference date is empty, then the current date is taken as the
// reference date. If onemonth and threemonth are true then three months are
// displayed. If months is given then the maximum of any combination of months
// is returned
func span(onemonth, threemonths bool, months int, ref []string) (int, time.Time, int) {

	var prev, post int

	// compute the number of months to display before the reference date and the
	// number of months shown from the reference date
	if onemonth {
		prev = 0
		post = 1
	}
	if threemonths {
		prev = 1
		post = 2
	}
	if months > 0 {

		if (threemonths && months > 2) || !threemonths {
			prev = 0
			post = months
		}
	}

	// now, compute the reference date. If neither the month nor the day are
	// given, they are assumed by default to be equal to the first one
	var day, year = 1, 1
	var month time.Month = 1
	if len(ref) == 0 {
		day = time.Now().Day()
		month = time.Now().Month()
		year = time.Now().Year()
	}
	if len(ref) == 1 {
		fmt.Sscanf(ref[0], "%d", &year)
	}
	if len(ref) == 2 {

		// note here the usage of strings.Join along with Sscanf to parse the
		// two ints altogether
		fmt.Sscanf(strings.Join(ref, " "), "%d %d", &month, &year)
	}
	if len(ref) == 3 {
		fmt.Sscanf(strings.Join(ref, " "), "%d %d %d", &day, &month, &year)
	}

	return prev, time.Date(year, month, day, 0, 0, 0, 0, time.UTC), post
}

// formatMonth
//
// writes in a matrix of strings the days of all dates comprising the interval
// [from, to). Those in the interval [ref0, from) are shown a little fainted, as
// much as those in the interval (to, ref1]. Those in the range [from, to) are
// shown in normal font. The current date is highlighted
func formatMonth(ref0, from, to, ref1 time.Time) []string {

	// get the current day, month and year
	now := time.Now()
	curr := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	output := make([]string, 0)

	// start with the initial date until the whole month has been exhausted
	for ref := ref0; ref.Before(ref1) || ref.Equal(ref1); ref = ref.AddDate(0, 0, 1) {

		var cell string
		if ref.Before(from) || ref.After(to) || ref.Equal(to) {
			if ref.Weekday() == time.Sunday {
				cell = fmt.Sprintf("\033[38;2;120;40;40m%2d\033[0m", ref.Day())
			} else if ref.Equal(curr) {
				cell = fmt.Sprintf("\033[38;2;100;100;10;1m%2d\033[0m", ref.Day())
			} else {
				cell = fmt.Sprintf("\033[38;2;90;90;90m%2d\033[0m", ref.Day())
			}
		} else if ref.Equal(curr) {
			cell = fmt.Sprintf("\033[38;2;210;210;10;1m%2d\033[0m", ref.Day())
		} else {

			// the only exception here to take care of is that sundays shall be
			// highlighted
			if ref.Weekday() == time.Sunday {
				cell = fmt.Sprintf("\033[38;2;180;10;10;1m%2d\033[0m", ref.Day())
			} else {
				cell = fmt.Sprintf("%2d", ref.Day())
			}
		}
		output = append(output, cell)
	}

	return output
}

// getDates
//
// return a slice of strings with all dates spanning the given month which is
// given as a full date starting in day 1. Notice the string spans over a whole
// block, i.e., from the first day before or equal to the day 1 of the specified
// month until the last day of the month or the subsequent last day of the next
// month
//
// To compute the dates within the block of a specific month, the first day of
// the week to show is required. If sunday is true then weeks start on sunday;
// otherwise they start on mondays.
func getDates(start time.Time, sunday bool) []string {

	// it is neccessary to know the zero date, which is the date in dd/mm/yyy
	// format of the first cell in the output calendar for the given month. For
	// this, we substract the number of days from the weekday of the 'start' date
	// to 'start' from sunday and add one to start from monday
	zero := start
	if sunday {
		if start.Weekday() != time.Sunday {
			zero = start.AddDate(0, 0, -int(start.Weekday()))
		}
	} else {
		if start.Weekday() != time.Sunday {
			zero = start.AddDate(0, 0, 1-int(start.Weekday()))
		} else {
			zero = start.AddDate(0, 0, -6)
		}
	}

	// finally, we also need to compute the last day of the block which shows
	// the current month. Again, this is done by adding days to the end date
	end := start.AddDate(0, 1, 0)
	horizon := end.AddDate(0, 0, -1)
	if sunday {
		if horizon.Weekday() != time.Saturday {
			horizon = horizon.AddDate(0, 0, 6-int(horizon.Weekday()))
		}
	} else {
		if horizon.Weekday() != time.Sunday {
			horizon = horizon.AddDate(0, 0, 7-int(horizon.Weekday()))
		}
	}

	// now that we know all the necessary dates (first day of the block, first
	// day of the month, last day of the month and last day of the block), we
	// return the string comprising all those dates in a single block
	return formatMonth(zero, start, end, horizon)
}

// getAllDates
//
// computes a matrix of strings, where each slices contains the dates of each
// month, in the period [start, end).
//
// To compute the dates of all blocks in the specified range, the first day of
// the week to show is required. If sunday is true then weeks start on sunday;
// otherwise they start on mondays.
func getAllDates(start, end time.Time, sunday bool) [][]string {

	// -- initialization
	output := make([][]string, 0)

	// and now just simply add the dates of each month from 'start' until the
	// 'end'
	for ref := start; ref.Before(end); ref = ref.AddDate(0, 1, 0) {
		output = append(output, getDates(ref, sunday))
	}

	// and return the requested dates
	return output
}

// displayMonths
//
// show all months on the standard output in blocks of the given width starting
// in the given reference date. If fullHeader takes the value true then both the
// month and the year are shown on top of each sheet
func displayMonths(ref time.Time, months [][]string, width int, fullHeader bool) {

	fmt.Println()

	// if the fullheader is not requested, then the year should be shown on top
	// of the calendar sheet. This function actually assumes that months of the
	// same year are to be displayed
	// if !fullHeader {
	// 	stryear := fmt.Sprintf("%d", ref.Year())
	// 	lftmargin := fmt.Sprintf("%%%d", 21*width+3*(width-1))
	// 	header := fmt.Sprintf()
	// 	fmt.Println()
	// }

	for index := 0; index < len(months); index += width {

		// show the headers for all months to be printed out in this iteration

		// -- month and, optionally, year
		var header string
		for idmonth := 0; idmonth < min(width, len(months)-index); idmonth++ {

			// retrieve the name of this month and be aware that the width of
			// each month is exactly equal to 21 characters ---when leaving one
			// blank space between successive days of the same month
			strmonth := fmt.Sprintf("%s", ref.AddDate(0, index+idmonth, 0).Month())
			if fullHeader {

				// if the full header has been requested, then add the year of
				// the current month
				strmonth += fmt.Sprintf(" %d", ref.AddDate(0, index+idmonth, 0).Year())
			}
			header += centerText(strmonth, 21) + "  "
			// header += fmt.Sprintf(
			// 	"%-21s",
			// 	fmt.Sprintf(fmt.Sprintf("%%%ds", len(strmonth)+(21-len(strmonth))/2),
			// 		strmonth)) + "  "
		}
		fmt.Printf("\033[38;2;10;160;120m%s\033[0m\n", header)

		// --weekdays

		// Some months might span over less weeks than others but they should
		// nevertheless be shown uniformly. Thus, compute the number of rows
		// required to show the next 'width' months.
		nbweeks := 0
		for _, month := range months[index:min(index+width, len(months))] {
			if len(month)/7 > nbweeks {
				nbweeks = len(month) / 7
			}
		}

		// second, all those months with less rows than the maximum computed
		// above, add a new one with blank spaces
		for idmonth := index; idmonth < min(index+width, len(months)); idmonth++ {
			for len(months[idmonth])/7 < nbweeks {

				// if so, add a new week with no data to show, until it equals
				// the number of weeks of the other months
				for day := 0; day < 7; day++ {
					months[idmonth] = append(months[idmonth], "  ")
				}
			}
		}

		// and now show the next 'width' months altogether
		for week := 0; week < nbweeks; week++ {
			line := make([]string, 0)

			// first, join all dates of each month with one space in between
			for idmonth := 0; idmonth < min(width, len(months)-index); idmonth++ {
				line = append(line, strings.Join(months[index+idmonth][7*week:7*(1+week)], " "))
			}

			// now, join the dates of all months with three blanks in between
			fmt.Println(strings.Join(line, "   "))
		}
		if index+width < len(months) {
			fmt.Println()
		}
	}
}

// main function
//
// given a number decide whether it is divisible by 7 or not
func main() {

	// first things first, parse the flags
	flag.Parse()

	// if the current version is requested, then show it on the standard output
	// and exit
	if want_version {
		showVersion(EXIT_SUCCESS)
	}

	// verify that the number of months, if given, is not a negative value, and
	// also that no more than three numbers are given without arguments to
	// specify a legal date
	if months < 0 {
		log.Fatal("Months out of range")
	}
	if len(flag.Args()) > 3 {
		log.Fatal("Incorrect date. Use --help for more information")
	}

	// now, get the number of months to show before a reference date as selected
	// by the user, and the number of monts to show after it
	prev, refdate, post := span(oneMonth, threeMonths, months, flag.Args())

	// compute the start and end dates of the period to show
	start := refdate.AddDate(0, -prev, 0)
	end := refdate.AddDate(0, post, 0)

	// but make sure to force those dates to start in day 1
	start = time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.UTC)
	end = time.Date(end.Year(), end.Month(), 1, 0, 0, 0, 0, time.UTC)

	// decide whether the header on top of each month should show the year or
	// not. The rule is to show it always unless a whole year is been shown
	fullHeader := true
	if months == 12 && start.Month() == time.January {
		fullHeader = false
	}

	// get all dates in the interval [start, end) and display them on the
	// standard output
	displayMonths(start, getAllDates(start, end, sunday), blocks, fullHeader)
	fmt.Println()
}
