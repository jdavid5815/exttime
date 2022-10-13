package exttime

import (
	"fmt"
	"strings"
	"time"
)

func MonthToInteger(month string) int {

	switch strings.ToLower(month) {
	case "jan", "january":
		return 1
	case "feb", "february":
		return 2
	case "mar", "march":
		return 3
	case "apr", "april":
		return 4
	case "may":
		return 5
	case "jun", "june":
		return 6
	case "jul", "july":
		return 7
	case "aug", "august":
		return 8
	case "sep", "september":
		return 9
	case "oct", "october":
		return 10
	case "nov", "november":
		return 11
	case "dec", "december":
		return 12
	}
	return 0
}

func Leapyear(year int) bool {

	if year%4 != 0 {
		// years not divisible by 4 are not leap years
		return false
	} else {
		if year%100 != 0 {
			// years divisible by 4 but not by 100 are leap years
			return true
		} else {
			// years divisible by 100 aren't leap years ..
			if year%400 != 0 {
				return false
			} else {
				// unless they are also divisible by 400
				return true
			}
		}
	}
}

/*
 * Convert a Gregorian date to a Julian Day Number.
 * Function is only valid for the year 1582 and later.
 * If an invalid year is given, zero will be returned.
 */
func JulianDayNumber(gregorian Date) float32 {

	var a, b, c, e, f int

	if gregorian.Year < 1582 {
		return 0.0
	}
	if gregorian.Month == 1 || gregorian.Month == 2 {
		gregorian.Year = gregorian.Year - 1
		gregorian.Month = gregorian.Month + 12
	}
	a = gregorian.Year / 100
	b = a / 4
	c = 2 - a + b
	e = int(365.25 * float32(gregorian.Year+4716))
	f = int(30.6001 * float32(gregorian.Month+1))
	return float32(c+gregorian.Day+e+f) - 1524.5
}

func StartOfEuropeanDST(year int) Date {

	var startofdst Date

	/* In the EU, DST starts the last Sunday of March at 1:00 AM UTC.
	 * So basically, we know all the required variables, except for
	 * the day. Start by looking at the 25th of March, which is
	 * exactly one week before the end of March, so there must be
	 * one and exactly one Sunday in that last week.
	 */
	startofdst.Year = year
	startofdst.Month = 3
	startofdst.Day = 25
	startofdst.Hour = 1
	startofdst.Minutes = 0
	t := time.Date(year, time.March, 25, 0, 0, 0, 0, time.UTC)
	switch t.Weekday() {
	case time.Monday:
		startofdst.Day += 6
	case time.Tuesday:
		startofdst.Day += 5
	case time.Wednesday:
		startofdst.Day += 4
	case time.Thursday:
		startofdst.Day += 3
	case time.Friday:
		startofdst.Day += 2
	case time.Saturday:
		startofdst.Day++
	}
	return startofdst
}

func EndOfEuropeanDST(year int) Date {

	var endofdst Date

	/* In the EU, DST ends the last Sunday of October at 1:00 AM UTC.
	 * So basically, we know all the required variables, except for
	 * the day. Start by looking at the 25th of October, which is
	 * exactly one week before the end of October, so there must be
	 * one and exactly one Sunday in that last week.
	 */
	endofdst.Year = year
	endofdst.Month = 10
	endofdst.Day = 25
	endofdst.Hour = 1
	endofdst.Minutes = 0
	t := time.Date(year, time.October, 25, 0, 0, 0, 0, time.UTC)
	switch t.Weekday() {
	case time.Monday:
		endofdst.Day += 6
	case time.Tuesday:
		endofdst.Day += 5
	case time.Wednesday:
		endofdst.Day += 4
	case time.Thursday:
		endofdst.Day += 3
	case time.Friday:
		endofdst.Day += 2
	case time.Saturday:
		endofdst.Day++
	}
	return endofdst
}

func EuropeanSummerTime(today Date) bool {

	var start, stop Date

	start = StartOfEuropeanDST(today.Year)
	stop = EndOfEuropeanDST(today.Year)
	if today.Month >= start.Month && today.Month <= stop.Month {
		switch today.Month {
		case start.Month:
			if today.Day >= start.Day {
				return today.Hour >= start.Hour
			}
			return false
		case stop.Month:
			fmt.Printf("Peekaboo!\n")
			if today.Day <= stop.Day {
				return today.Hour <= stop.Hour
			}
			return false
		default:
			return true
		}
	}
	return false
}
