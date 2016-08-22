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

	return availableTiles
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
	// Tile checking logic goes here
	currentLevel := pyramid.tileLevels[levelKey]
	tile := currentLevel[row][column]

	if levelKey == 1 {
		return !tile.taken
	} else {
		if hasTilesAbove(row, column, levelKey, pyramid) {
			return false
		}
	}

	return true
}

func hasTilesAbove(row, column int, level int8, pyramid *Pyramid) bool {
	levelAbove := pyramid.tileLevels[level-1]

	for i := row; i > i-2; i-- {
		for j := column; j > j-2; j-- {
			if isPermittedIndex(i, level-1) && isPermittedIndex(j, level-1) {
				if levelAbove[i][j].taken == false {
					return true
				}
			}
		}
	}

	return false
}

func isPermittedIndex(index int, level int8) bool {
	return index >= 0 && index < int(level)
}

func dequeue(collection *[]Tile) Tile {
	length := len(*collection)
	result := (*collection)[length-1]
	*collection = (*collection)[:length-1]
	return result
}
