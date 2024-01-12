package utils

import (
	"fmt"
	"strings"
)

type stringLib struct {
}

type StringLib interface {
	Stringify(a []int) string
}

func NewStringLib() *stringLib {
	return &stringLib{}
}

func (s *stringLib) Stringify(a []int) string {
	return strings.Trim(fmt.Sprint(a), "[]")
}
