package utils

import "strconv"

func IntToStringSlice(slice []int32) []string {
	newSlice := make([]string, 0, len(slice))
	for _, el := range slice {
		newSlice = append(newSlice, strconv.FormatInt(int64(el), 10))
	}
	return newSlice
}
