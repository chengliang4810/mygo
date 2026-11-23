package ui

import "testing"

func TestAccessAfterListStateReset(t *testing.T) {
	var list ListState
	tt := NewTester(func(c *Context) {
		List(c, &list, 1, func(int) { Text(c, "Row") }).Grow(1)
	}, 300, 400)
	list = ListState{}
	tt.rt.accessTree()
}
