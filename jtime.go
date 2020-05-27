package jalali

import (
	"fmt"
	"time"
)

const (
	secInMin  = 60
	minInHour = 60
	secInHour = minInHour * secInMin
	hourInDay = 24
	minInDay  = hourInDay * minInHour
	secInDay  = secInHour * hourInDay
)

const (
	jEpocDay        = 10 // Zero based
	jEpocMonth      = 9  // Dey, zero based
	jEpocYear       = 1348
	jEpocDiff       = 286 // Epoch day is the 286th day in jalali year, (zero based)
	jEpocWDay       = time.Thursday
	jWeekLen        = 7
	jEpocLeapBefore = 211 // The year 1348 was not a leap year

	jYearLen = 365
)

var (
	jMonthLen    = [12]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	jAccuonthLen = [12]int{0, 31, 62, 93, 124, 155, 186, 216, 246, 276, 306, 336}
)

// secToJTime converts the sec (positive only) to day, hour minute and seconds.
// TODO : handle negative values for old times?
func secToDays(sec int) (day, hour, minute, second int) {
	day = sec / secInDay
	sec %= secInDay
	hour = sec / secInHour
	sec %= secInHour
	minute = sec / secInMin
	second = sec % secInMin

	return
}

// TODO: check if instead of error we need to panic, i.e, we generate the day and is
// not from user input
func dayToMonth(days int) (int, int, error) {
	if days < 0 || days > 365 {
		return 0, 0, fmt.Errorf("the days in the year is out of range [0-365] but is %d", days)
	}

	var month int
	for _, ml := range jMonthLen {
		if days > ml {
			days -= ml
			month++
		} else {
			break
		}
	}

	return month, days, nil
}

func monthDayToYear(month, days int) (int, error) {
	if month < 0 || month > 11 {
		return 0, fmt.Errorf("month %d is out of range, it should be [0-11]", month)
	}

	if days < 0 || days > jMonthLen[month] {
		return 0, fmt.Errorf("for month %d day %d is out of range", month, days)
	}

	ydays := jAccuonthLen[month]
	ydays += days

	return ydays, nil
}

func dayToWeekday(days int) time.Weekday {
	p := (days + int(jEpocWDay)) % jWeekLen
	if p < 0 {
		p += jWeekLen
	}

	return time.Weekday(p)
}

func dayToYear(days int) (int, int) {
	direction := 1
	current := jEpocYear

	days += jEpocDiff

	if days < 0 {
		direction = -1
	}

	// TODO: Predict and jump to the nearest year, this loop is shit.
	for ; ; current += direction {
		ln := jYearLen
		if IsLeap(current) {
			ln++
		}

		if days >= 0 && days < ln {
			break
		}
		days -= (direction * ln)
	}

	return current, days
}

func dayToJTime(days int) (year, month, day int, err error) {
	var remains int
	year, remains = dayToYear(days)
	month, day, err = dayToMonth(remains)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("converting date failed: %w", err)
	}

	return
}
