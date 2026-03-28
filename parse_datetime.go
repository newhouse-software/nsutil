package nsutil

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ParsePSJsonDateTime(jsonDate string) (time.Time, error) {
	// parse a JSON date string returned by PowerShell / WMI / CIM cmdlets
	//  jsonDate something like  /Date(1697059200000)/
	re := regexp.MustCompile(`/Date\((\d+)\)/`)
	matches := re.FindStringSubmatch(jsonDate)
	if len(matches) < 2 {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	ms, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing milliseconds: %w", err)
	}

	t := time.UnixMilli(ms)
	return t, nil
}

func ParseWMIDateTime(s string) (time.Time, error) {
	if len(s) < 25 {
		return time.Time{}, fmt.Errorf("invalid WMI datetime")
	}

	// Split into datetime and minute offset
	base := s[:21]         // yyyymmddHHMMSS.mmmmmm
	offsetMinStr := s[21:] // ±MMM

	offsetMin, err := strconv.Atoi(offsetMinStr)
	if err != nil {
		return time.Time{}, err
	}

	sign := "+"
	if offsetMin < 0 {
		sign = "-"
		offsetMin = -offsetMin
	}

	hours := offsetMin / 60
	mins := offsetMin % 60

	// Convert to ±HHMM format required by time.Parse
	offset := fmt.Sprintf("%s%02d%02d", sign, hours, mins)
	final := base + offset
	return time.Parse("20060102150405.000000-0700", final)
}

func ParseDMTFDateTime(dmtf string) (time.Time, error) {
	if len(dmtf) < 14 {
		return time.Time{}, fmt.Errorf("invalid DMTF date string")
	}
	layout := "20060102150405"
	parsedTime, err := time.Parse(layout, dmtf[:14])
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func ParseDateString(dateStr string) (time.Time, error) {
	layout := "1/2/2006"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
