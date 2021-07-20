package utils

import (
	"fmt"
	"strconv"
)

func FloatToString(misc float64) string {
	return fmt.Sprintf("%f", misc)
}

func StringToFloat(misc string) float64 {
	tmpFloat, _ := strconv.ParseFloat(misc, 64)
	return tmpFloat
}

func IntToString(misc int) string {
	return fmt.Sprintf("%d", misc)
}
