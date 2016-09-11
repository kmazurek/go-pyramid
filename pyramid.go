package pyramid

type Tile struct {
	symbols [2][2]int8
	taken   bool
}

type Pyramid struct {
	numberOfLevels int8
	tileLevels     map[int8][][]*Tile
}

func (pyramid *Pyramid) GetAvailableTiles() []Tile {
	var availableTiles []Tile
	// calls checkTile and fills in availableTiles
	for level, tiles := range pyramid.tileLevels {
		for i := 0; i < int(level); i++ {
			for j := 0; j < int(level); j++ {
				if pyramid.checkTile(i, j, level) {
					availableTiles = append(availableTiles, *(tiles[i][j]))
				}
			}
		}
	}

	return availableTiles
}

func (pyramid *Pyramid) getNeighbourhoodThreshold() int8 {
	return int8(2)
}

func NewPyramid(numberOfLevels int8) (*Pyramid, int) {
	result := new(Pyramid)
	tileCounter := 0

	result.numberOfLevels = numberOfLevels
	result.tileLevels = make(map[int8][][]*Tile)

	for i := int8(1); i <= numberOfLevels; i++ {
		result.tileLevels[i] = make([][]*Tile, i, i)
		currentLevel := result.tileLevels[i]

		for j := int8(0); j < i; j++ {
			currentLevel[j] = make([]*Tile, i, i)

			for k := int8(0); k < i; k++ {
				currentLevel[j][k] = new(Tile)
				tileCounter++
			}
		}
	}

	return result, tileCounter
}

func (pyramid *Pyramid) checkTile(row, column int, levelKey int8) bool {
	currentLevel := pyramid.tileLevels[levelKey]
	tile := currentLevel[row][column]

	if levelKey == 1 {
		return !tile.taken
	}

	if hasTilesAbove(row, column, levelKey, pyramid) {
		return false
	}

	return isThresholdFulfilled(row, column, levelKey, pyramid)
}

func hasTilesAbove(row, column int, level int8, pyramid *Pyramid) bool {
	levelAbove := pyramid.tileLevels[level-1]

	for i := row; i >= row-1; i-- {
		for j := column; j >= column-1; j-- {
			if isPermittedCell(i, j, level-1) {
				if levelAbove[i][j].taken == false {
					return true
				}
			}
		}
	}

	return false
}

func isPermittedCell(row, column int, level int8) bool {
	return (row >= 0 && row < int(level)) && (column >= 0 && column < int(level))
}

func isThresholdFulfilled(row, column int, level int8, pyramid *Pyramid) bool {
	currentLevel := pyramid.tileLevels[level]
	freeCellCount := int8(0)
	modifiers := []int{-1, 1}

	for _, modifier := range modifiers {
		if isPermittedCell(row+modifier, column, level) {
			if currentLevel[row+modifier][column].taken {
				freeCellCount++
			}
		} else {
			freeCellCount++
		}

		if isPermittedCell(row, column+modifier, level) {
			if currentLevel[row][column+modifier].taken {
				freeCellCount++
			}
		} else {
			freeCellCount++
		}
	}

	return freeCellCount >= pyramid.getNeighbourhoodThreshold()
}

func dequeue(collection *[]Tile) Tile {
	length := len(*collection)
	result := (*collection)[length-1]
	*collection = (*collection)[:length-1]
	return result
}
