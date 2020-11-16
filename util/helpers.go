package util

import (
	"text/template"
)

func Sequence(n int) []int {
	ret := make([]int, 0)
	for i := 1; i <= n; i++ {
		ret = append(ret, i)
	}
	return ret
}

var HelperFuncs = template.FuncMap{
	"seq":    Sequence,
}