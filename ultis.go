package main

import "fmt"

type multiString []string

func (m *multiString) String() string {
	if m == nil {
		return "[]"
	}
	return fmt.Sprintf("%v", *m)
}

func (m *multiString) Set(value string) error {
	if m == nil {
		*m = make([]string, 0)
	}
	*m = append(*m, value)
	return nil
}
