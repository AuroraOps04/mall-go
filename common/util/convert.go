package util

import (
	"strconv"
	"strings"
)


func Str2IntSlice (str *string, sep string) []int{
	strs := strings.Split(*str, sep) 
	var ints []int
	for _,v := range strs{
		i,err := strconv.Atoi(strings.TrimSpace(v))	
		if err == nil{
			ints = append(ints, i)
		}

	}
	return ints
}
