package main

import (
	"reflect"
	"testing"
)

func TestNeighbours(t *testing.T) {
	var tests = []struct {
		cell Cell
		want *CellSet
	}{
		{
			Cell{2, 4}, CellSetFrom(
				Cell{1, 3}, Cell{1, 4}, Cell{1, 5},
				Cell{2, 3}, Cell{2, 5},
				Cell{3, 3}, Cell{3, 4}, Cell{3, 5}),
		},
	}

	for _, c := range tests {
		got := neighbours(&c.cell)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("neighbours(%q) == %v, want %v", c.cell, got, c.want)
		}
	}
}
