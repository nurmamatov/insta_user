package utils

import "strings"

func ParseFilter(s string) []string {
	return strings.Split(s, ",")
}

func StringSliceToInterfaceSlice(s []string) []interface{} {
	is := make([]interface{}, len(s))
	for i, v := range s {
		is[i] = v
	}
	return is
}

