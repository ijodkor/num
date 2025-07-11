// support -> num

package main

import (
	"strconv"
)

func ToInt(s string) int {
	// string to int
	v, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return v
}

func ToInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}
