package utils

import "strconv"

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}
