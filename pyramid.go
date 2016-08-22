package pyramid

type Tile struct {
	symbols [2][2]int8
	taken   bool
}

type Pyramid struct {
	tileLevels map[int8][][]*Tile
}

func (pyramid *Pyramid) GetAvailableTiles() []Tile {
	var availableTiles []Tile
	// calls checkTile and fills in availableTiles

	return availableTiles
}

func NewPyramid(numberOfLevels int8) (*Pyramid, int) {
	result := new(Pyramid)
	tileCounter := 0

	result.tileLevels = make(map[int8][][]*Tile)

	for int8 i = 1; i <= numberOfLevels; i++ {
		result.tileLevels[i] = make([][]*Tile, i, i)
		currentLevel := result.tileLevels[i]

		for j := 0; j < i; j++ {
			currentLevel[j] = make([]*Tile, i, i)

			for k := 0; k < i; k++ {
				currentLevel[j][k] = new(Tile)
				tileCounter++
			}
		}
	}

	return result, tileCounter
}

func checkTile(tile Tile, level int8) bool {
	// Tile checking logic goes here
	return true
}

func dequeue(collection *[]Tile) Tile {
	length := len(*collection)
	result := (*collection)[length-1]
	*collection = (*collection)[:length-1]
	return result
}
