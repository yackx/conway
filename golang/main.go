package main

/*
A simple Conway's Game of Life in Go.

No cycle detection. Infinite grid. No GUI.

Copyright (C) 2014-2025 Youri Ackx under GNU General Public License.
See the LICENSE file and [http://www.gnu.org/licenses/].
*/

import (
	"fmt"
)

func main() {
	// Basic grid
	basic := CellSetFrom(
		Cell{0, 0}, Cell{1, 0}, Cell{2, 0},
		Cell{1, 1}, Cell{2, 1})
	/*
		// Blinker
		blinker := CellSetFrom(
				Cell{1, 1}, Cell{2, 1}, Cell{3, 1})

		// Glider
		glider :=  CellSetFrom(
				Cell{1, 0}, Cell{2, 1}, Cell{0, 2},
				Cell{1, 2}, Cell{2, 2})

		// Die hard
		diehard := CellSetFrom(
				Cell{1, 2}, Cell{2, 2}, Cell{2, 3},
				Cell{7, 1}, Cell{6, 3}, Cell{7, 3},
				Cell{8, 3})
	*/

	grid := basic
	fmt.Println(grid.ToString())
	for !grid.Empty() {
		grid = Next(grid)
		fmt.Println(grid.ToString())
	}
}
