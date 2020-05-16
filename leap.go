package jalali

const (
	firstCycle = 2346
	extraPart  = 24219858156
	round      = 100000000000

	leapCycle   = 2820
	leapInCycle = 683
)

// leapStatus checks if the Jalali year is leap or not and how many full leap has
// passed since the current cycle start
func leapStatus(year int) (int, bool) {
	extraDays := (year + firstCycle) * extraPart
	fullLeap := extraDays / round
	thisYearExtra := extraDays - (fullLeap * round)
	// Are we in the second leap?
	// TODO: Just for fun, add the third cycle
	if year > leapCycle-firstCycle {
		fullLeap -= leapInCycle
	}
	return fullLeap, thisYearExtra < extraPart
}

// IsLeap reaturns true if the year is leap.
// This function is based on the Musa Akrami paper: The development of Iranian calendar: historical and astronomical
// foundations.
// Based on this paper, Each year has an extra 0.24219858156 day in the end. So the first year, which is the year
// -2346, starts precisely at noon, and the next year at 12:00+1*0.24219858156*(24h), the next, is
// 12:00+2*0.24219858156*(24h) and so on.
// In the 5th year, the extra part is 5*0.24219858156*(24h), which is 1.21099290780 (A day and a little extra)
// So in the 4th year, it does not leap, but the 5th year leaps one day, and it is a leap year.
// The only way to detect this is to find the jumps. If the part after the floating point is less then the
// 0.24219858156, this is the leap number.
// For example, the year 1386:
// Extra days are (1386+2346)*0.24219858156, exactly 903.88510638192 days, so it's not a leap year
// (0.88510638192 > 0.24219858156) It means 903 full days and 0.88510638192 of a day.
// For the next year (1387), this is 904.12730496348 days, which means the year 1387 is a leap year
// (0.12730496348 < 0.24219858156)
// XXX (f0ruD): To prevent using floating-point numbers, (which can be tricky in some languages) I decided to
// use int, I can use float64 instead, not sure if this has any effect
func IsLeap(year int) bool {
	_, leap := leapStatus(year)
	return leap
}
