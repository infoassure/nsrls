package goutil

import (
	"strconv"
	"strings"
	"time"
)

// ##### Constants ###########################################################

const MIN_TIME = "2006-01-02T15:04:05Z07:00"

// ##### Methods #############################################################

// Converts an Int64 to a string
func ConvertInt64ToString(data int64) string {
	return strconv.FormatInt(data, 10 )
}

// Converts an Int64 to a string
func ConvertIntToString(data int) string {
	return strconv.FormatInt(int64(data), 10 )
}

// Converts an Int64 to a string
func ConvertUInt16ToString(data uint16) string {
	return strconv.FormatInt(int64(data), 10 )
}

// Converts a string to an int64
func ConvertStringToInt64(data string) int64 {
	ret, err := strconv.ParseInt(data, 0, 64)
	if err != nil {
		return -1
	}
	return ret
}

//
func ParseBoolean(data string, match string, retVal bool) bool {
	data = RemoveQuotes(data)
	if strings.ToLower(data) == strings.ToLower(match) {
		return retVal
	} else {
		if retVal == true {
			return false
		} else {
			return true
		}
	}
}

//
func ParseInt(data string, intType string) int {
	if len(data) == 0 {
		return 0
	}

	tmpInt, err := strconv.Atoi(data)
	if err != nil {
		return 0
	}

	return tmpInt
}

//
func ParseFloat(data string, floatType string) float64 {
	if len(data) == 0 {
		return 0
	}

	tmpFloat, err := strconv.ParseFloat(data, 32)
	if err != nil {
		return 0
	}

	return tmpFloat
}

// Parses a date/time string and formats as a PostGres compatible
// string, returns a default golang value if blank or enable to parse
func ParseTimestamp(layout, data string) time.Time {
	if len(data) == 0 {
		parsedTimestamp, _ := time.Parse(time.RFC3339, MIN_TIME)
		return parsedTimestamp
	}

	parsedTimestamp, err := time.Parse(layout, strings.TrimSpace(data))
	if err != nil {
		parsedTimestamp, _ := time.Parse(time.RFC3339, MIN_TIME)
		return parsedTimestamp
	}

	return parsedTimestamp
}
