/*
	Copyright 2019 Brian Bauer

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

/*
Package dateslice creates slices containing time.Time elements
*/
package dateslice

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

/*
Sometimes you need a slice of dates.  Here are some functions that make that
	a little easier.
*/

/*
Today returns a slice containing a single element - the current day
*/
func Today() []time.Time {
	return []time.Time{time.Now()}
}

/*
Yesterday returns a slice containing a single element - yesterday
*/
func Yesterday() []time.Time {
	return []time.Time{time.Now().AddDate(0, 0, -1)}
}

/*
Tomorrow returns a slice containing a single element - tomorrow
*/
func Tomorrow() []time.Time {
	return []time.Time{time.Now().AddDate(0, 0, 1)}
}

func aWeek(baseDate time.Time) []time.Time {
	ds := make([]time.Time, 7)

	dow := baseDate.Weekday()

	// Reset the base date to Sunday
	baseDate = baseDate.AddDate(0, 0, 0-int(dow))

	for i := range ds {
		ds[i] = baseDate.AddDate(0, 0, i)
	}

	return ds
}

/*
WeekOf returns a slice containing all dates that occur during this specific week
	(Sunday is the first day of the week in Go!)
*/
func WeekOf(date time.Time) []time.Time {
	return aWeek(date)
}

/*
ThisWeek returns a slice containing all dates that occur this week
	(Sunday is the first day of the week in Go!)
*/
func ThisWeek() []time.Time {
	return aWeek(time.Now())
}

/*
LastWeek returns a slice containing all dates that occured last week
	(Sunday is the first day of the week in Go!)
*/
func LastWeek() []time.Time {
	return aWeek(time.Now().AddDate(0, 0, -7))
}

/*
NextWeek returns a slice containing all dates that will occur next week
	(Sunday is the first day of the week in Go!)
*/
func NextWeek() []time.Time {
	return aWeek(time.Now().AddDate(0, 0, 7))
}

func aMonth(baseDate time.Time) []time.Time {
	// This is used for subtraction, so the first day of the month needs to be a 0 instead of a 1
	dom := baseDate.Day() - 1

	//reset the base date to the 1st of the month
	baseDate = baseDate.AddDate(0, 0, 0-int(dom))

	firstOfNextMonth := baseDate.AddDate(0, 1, 0)
	daysInThisMonth := firstOfNextMonth.Sub(baseDate).Hours() / 24.0
	fmt.Printf("%f days in the month\n", math.Ceil(daysInThisMonth))

	ds := make([]time.Time, int(math.Ceil(daysInThisMonth)))

	for i := range ds {
		ds[i] = baseDate.AddDate(0, 0, i)
	}

	return ds
}

/*
ThisMonth returns a slice containing all dates that occur this month
*/
func ThisMonth() []time.Time {
	return aMonth(time.Now())
}

/*
LastMonth returns a slice containing all dates that occured last month
*/
func LastMonth() []time.Time {
	return aMonth(time.Now().AddDate(0, -1, 0))
}

/*
NextMonth returns a slice containing all dates that will occur next month
*/
func NextMonth() []time.Time {
	return aMonth(time.Now().AddDate(0, 1, 0))
}

/*
MonthOf returns a slice containing all dates that occur in the specific month
*/
func MonthOf(date time.Time) []time.Time {
	return aMonth(date)
}

func aYear(baseDate time.Time) []time.Time {
	// This is used for subtraction, so the first day of the month needs to be a 0 instead of a 1
	dom := baseDate.YearDay() - 1

	//reset the base date to the 1st of the month
	baseDate = baseDate.AddDate(0, 0, 0-int(dom))

	firstOfNextYear := baseDate.AddDate(1, 0, 0)
	daysInThisYear := firstOfNextYear.Sub(baseDate).Hours() / 24.0
	fmt.Printf("%f days in the year\n", math.Ceil(daysInThisYear))

	ds := make([]time.Time, int(math.Ceil(daysInThisYear)))

	for i := range ds {
		ds[i] = baseDate.AddDate(0, 0, i)
	}

	return ds
}

/*
Range returns a slices of dates specified in the range
*/
func Range(beg, end time.Time) []time.Time {
	daysInRange := end.Sub(beg).Hours()/24.0 + 1.0

	ds := make([]time.Time, int(math.Ceil(daysInRange)))

	for i := range ds {
		ds[i] = beg.AddDate(0, 0, i)
	}

	return ds
}

/*
RangeString transforms a beginning and ending date from strings into dates and then returns
	the results of the Range function
*/
func RangeString(begDt, endDt string) []time.Time {
	if len(begDt) == 6 {
		//Treat this as the beginning of a month
		begDt += "01"
	}
	if len(begDt) == 4 {
		//Treat this as the beginning of a year
		begDt += "0101"
	}
	if len(endDt) == 6 {
		//Treat this as the end of a month
		endDt += "01"

		temp, err := time.Parse("20060102", endDt)
		if err != nil {
			fmt.Println(err)
		}
		temp = temp.AddDate(0, 1, 0)
		temp = temp.AddDate(0, 0, -1)
		endDt = temp.Format("20060102")
	}
	if len(endDt) == 4 {
		//Treat this as the end of a month
		endDt += "0101"

		temp, err := time.Parse("20060102", endDt)
		if err != nil {
			fmt.Println(err)
		}
		temp = temp.AddDate(1, 0, 0)
		temp = temp.AddDate(0, 0, -1)
		endDt = temp.Format("20060102")
	}
	beg, err := time.Parse("20060102", begDt)
	if err != nil {
		fmt.Println(err)
	}
	end, err := time.Parse("20060102", endDt)
	if err != nil {
		fmt.Println(err)
	}
	return Range(beg, end)
}

/*
DateStringToSlice returns a slice of dates corresponding to the text in a string
*/
func DateStringToSlice(dateText string) []time.Time {
	var ds []time.Time
	if strings.EqualFold(dateText, "today") {
		ds = Today()
	} else if strings.EqualFold(dateText, "yesterday") {
		ds = Yesterday()
	} else if strings.EqualFold(dateText, "thisweek") {
		ds = ThisWeek()
	} else if strings.EqualFold(dateText, "lastweek") {
		ds = LastWeek()
	} else if strings.EqualFold(dateText, "thismonth") {
		ds = ThisMonth()
	} else if strings.EqualFold(dateText, "lastmonth") {
		ds = LastMonth()
	}
	return ds
}

/*
DateObjectsToSlice returns a slice of dates based upon the contents of 3 flags
*/
func DateObjectsToSlice(dateString, begDt, endDt string) (ds []time.Time) {
	if len(dateString) != 0 {
		ds = DateStringToSlice(dateString)
	}

	if len(begDt) != 0 {
		if len(endDt) != 0 {
			ds = RangeString(begDt, endDt)
		} else {
			ds = RangeString(begDt, begDt)
		}
	}
	return
}

/*
ParseDate tries to figure out what is being asked for
*/
func ParseDate(dateString, rangeType string) (ds []time.Time) {
	switch rangeType {
	case "daily":
		if strings.EqualFold(dateString, "today") {
			ds = Today()
		} else if strings.EqualFold(dateString, "yesterday") {
			ds = Yesterday()
		} else if _, err := strconv.Atoi(dateString); err == nil {
			ds = RangeString(dateString, dateString)
		}
	case "weekly":
		if strings.EqualFold(dateString, "current") {
			ds = ThisWeek()
		} else if strings.EqualFold(dateString, "last") {
			ds = LastWeek()
		} else if _, err := strconv.Atoi(dateString); err == nil {

		}
	case "monthly":
	}
	return
}
