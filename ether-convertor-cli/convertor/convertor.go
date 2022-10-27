package convertor

import (
	"math"
	"strconv"
)

//ConvertorReceiver : write
func ConvertorReceiver(to, from, value *string) float64 {
	var num = parseFloat(*value)

	chart := make(map[string]int)
	chart["wei"] = 0
	chart["kwei"] = 3
	chart["mwei"] = 6
	chart["gwei"] = 9
	chart["microether"] = 12
	chart["milliether"] = 15
	chart["ether"] = 18

	if chart[*from] <= chart[*to] {
		return (num/math.Pow(float64(10),float64(chart[*to])))*math.Pow(float64(10),float64(chart[*from]))
	} else {
		return num*math.Pow(float64(10),float64(chart[*to]))
	}
}

func parseFloat(value string) float64 {
	val,err := strconv.ParseFloat(value,64)
	if err != nil {
		panic("Something went wrong")
	}
	return val
}
