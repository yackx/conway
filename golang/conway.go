package main

/*
A simple Conway's Game of Life in Go.

No cycle detection. Infinite grid. No GUI.

Copyright (C) 2014-2025 Youri Ackx under GNU General Public License.
See the LICENSE file and [http://www.gnu.org/licenses/].
*/

// Return all the neighbours of the given cell, living or not.
func neighbours(cell *Cell) *CellSet {
	n := NewCellSet()
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				n.Add(Cell{i + cell.X, j + cell.Y})
			}
		}
	}
	return n
}

// Next Moves to the next state
func Next(grid *CellSet) *CellSet {
	newGrid := NewCellSet()
	candidates := NewCellSet()

	// Retrieve all cells to check - living cells and theirs neighbours
	for _, living := range grid.Cells() {
		candidates.Add(living)
		neighbours := neighbours(&living)
		for _, neighbour := range neighbours.Cells() {
			candidates.Add(neighbour)
		}
	}

	// Check each candidate against the game's rules
	for _, cell := range candidates.Cells() {
		livingNeighboursCount := len(neighbours(&cell).Intersect(grid).Cells())
		isAlive := grid.Contains(cell)

		if (isAlive && (livingNeighboursCount == 2 || livingNeighboursCount == 3)) ||
			(!isAlive && livingNeighboursCount == 3) {
			newGrid.Add(cell)
		}
	}

	return newGrid
}
