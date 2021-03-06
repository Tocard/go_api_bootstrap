package data

import (
	"fmt"
	"strconv"
)

func GetFees() (float64, string) {
	return 0.001, "PPLNS"
}

func GetName() string {
	return "French Farmers"
}

func GetLogo() string {
	return API_HOST + ":" + API_PORT + "/logo.svg"
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

const (
	TB = 1000000000000
	GB = 1000000000
	MB = 1000000
	KB = 1000
)

func lenReadable(length int, decimals int, displayUnit bool) (out string) {
	var unit string
	var i int
	var remainder int

	// Get whole number, and the remainder for decimals
	if length > TB {
		unit = "TB"
		i = length / TB
		remainder = length - (i * TB)
	} else if length > GB {
		unit = "GB"
		i = length / GB
		remainder = length - (i * GB)
	} else if length > MB {
		unit = "MB"
		i = length / MB
		remainder = length - (i * MB)
	} else if length > KB {
		unit = "KB"
		i = length / KB
		remainder = length - (i * KB)
	} else {
		return strconv.Itoa(length) + " B"
	}

	if decimals == 0 {
		return strconv.Itoa(i) + " " + unit
	}

	// This is to calculate missing leading zeroes
	width := 0
	if remainder > GB {
		width = 12
	} else if remainder > MB {
		width = 9
	} else if remainder > KB {
		width = 6
	} else {
		width = 3
	}

	// Insert missing leading zeroes
	remainderString := strconv.Itoa(remainder)
	for iter := len(remainderString); iter < width; iter++ {
		remainderString = "0" + remainderString
	}
	if decimals > len(remainderString) {
		decimals = len(remainderString)
	}
	if displayUnit == true {
		return fmt.Sprintf("%d.%s %s", i, remainderString[:decimals], unit)
	}
	return fmt.Sprintf("%d.%s", i, remainderString[:decimals])
}
