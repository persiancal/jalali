package jalali

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLeap(t *testing.T) {
	// This data extracted from the official pdf sources here : https://calendar.ut.ac.ir/Fa/Tyear/PastCal.asp
	table := map[int]bool{
		1371: false,
		1372: false,
		1373: false,
		1374: false,
		1375: true,
		1376: false,
		1377: false,
		1378: false,
		1379: true,
		1380: false,
		1381: false,
		1382: false,
		1383: true,
		1384: false,
		1385: false,
		1386: false,
		1387: true,
		1388: false,
		1389: false,
		1390: false,
		1391: true,
		1392: false,
		1393: false,
		1394: false,
		1395: true,
		1396: false,
		1397: false,
		1398: false,
		1399: true,
	}

	for year, leap := range table {
		require.Equal(t, leap, IsLeap(year))
	}
}
