package jalali

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonthDayToYear(t *testing.T) {
	for y := 0; y < 366; y++ {
		m, d, err := dayToMonth(y)
		require.NoError(t, err)

		y2, err := monthDayToYear(m, d)
		require.NoError(t, err)
		assert.Equal(t, y, y2)
	}
}

func TestWeekDay(t *testing.T) {
	base := time.Time{}
	for i := -1000; i < 1000; i++ {
		day := base.Add(time.Duration(i*24) * time.Hour)
		since := int(day.Unix() / secInDay)

		wd := dayToWeekday(since)
		require.Equal(t, day.Weekday(), wd)
	}
}

func TestDayToJTime(t *testing.T) {
	// This test is not reliable. it depends on the same code and the leap year
	// function which can be wrong
	month := jEpocMonth
	day := jEpocDay
	for i := 0; i < 100; i++ {
		if IsLeap(i + jEpocYear - 1) {
			day--
			if day <= 0 {
				month--
				day = jMonthLen[month-1]
			}
		}
		y, m, d, err := dayToJTime(i * 365)
		require.NoError(t, err)
		assert.Equal(t, jEpocYear+i, y)
		assert.Equal(t, month, m, fmt.Sprintf("%d - %d - %d", i, month, d))
		assert.Equal(t, day, d)
	}

	current := jEpocDiff
	for i := 0; i < 300; i++ {
		if IsLeap(jEpocYear - i + 1) {
			current++
		}
		y, d := dayToYear(-i * 365)
		assert.Equal(t, jEpocYear-i, y)
		assert.Equal(t, current, d)

	}
}
