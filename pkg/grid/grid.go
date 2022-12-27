package grid

import cell "github.com/Vallghall/golmie/pkg/Cell"

type Grid struct {
	Rows *[50][50]*cell.Cell
}

func New(cellSize int) *Grid {
	grid := &Grid{
		Rows: new([50][50]*cell.Cell),
	}
	for i, row := range grid.Rows {
		for j := range row {
			grid.Rows[i][j] = cell.New(cellSize)
		}
	}

	return grid
}

func (g *Grid) NextGeneration() {
	next := new([50][50]*cell.Cell)
	for i, row := range g.Rows {
		for j := range row {
			next[i][j] = g.aliveNeighbours(i, j)
		}
	}

	g.Rows = next
}

func (g *Grid) aliveNeighbours(i, j int) *cell.Cell {
	count := 0
	neighbours := make([]*cell.Cell, 0)
	rows := g.Rows

	if i-1 >= 0 {
		neighbours = append(neighbours, rows[i-1][j])

		if j-1 >= 0 {
			neighbours = append(neighbours, rows[i-1][j-1])
		}

		if j+1 < len(rows) {
			neighbours = append(neighbours, rows[i-1][j+1])
		}
	}

	if i+1 < len(rows) {
		neighbours = append(neighbours, rows[i+1][j])

		if j-1 >= 0 {
			neighbours = append(neighbours, rows[i+1][j-1])
		}

		if j+1 < len(rows) {
			neighbours = append(neighbours, rows[i+1][j+1])
		}
	}

	if j+1 < len(rows) {
		neighbours = append(neighbours, rows[i][j+1])
	}

	if j-1 >= 0 {
		neighbours = append(neighbours, rows[i][j-1])
	}

	for _, n := range neighbours {
		if n.Alive {
			count++
		}
	}

	current := g.Rows[i][j]
	if (current.Alive && (count < 2 || count > 3)) || (!current.Alive && count != 3) {
		return &cell.Cell{
			Alive: false,
			Size:  current.Size,
		}
	}

	return &cell.Cell{
		Alive: true,
		Size:  current.Size,
	}
}
