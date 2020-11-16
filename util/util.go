package util

import "fmt"

func Quotify(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}