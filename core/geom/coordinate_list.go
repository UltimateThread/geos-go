package geos

type CoordinateList struct {
	Coordinates []Coordinate
}

/**
 * Constructs a new list without any coordinates
 */
func DefaultCoordinateList() *CoordinateList {
	cl := new(CoordinateList)
	return cl
}

/**
 * Constructs a new list from an array of Coordinates, allowing repeated points.
 * (I.e. this constructor produces a {@link CoordinateList} with exactly the same set of points
 * as the input array.)
 *
 * @param coord the initial coordinates
 */
func NewCoordinateList(coords []Coordinate) *CoordinateList {
	new := new(CoordinateList)
	new.AddCoordinateListRepeated(coords, true)
	return new
}

/**
 * Constructs a new list from an array of Coordinates,
 * allowing caller to specify if repeated points are to be removed.
 *
 * @param coord the array of coordinates to load into the list
 * @param allowRepeated if <code>false</code>, repeated points are removed
 */
func NewCoordinateListWithRepeated(coords []Coordinate, allowRepeated bool) *CoordinateList {
	new := new(CoordinateList)
	new.AddCoordinateListRepeated(coords, allowRepeated)
	return new
}

func (coordinateList *CoordinateList) GetCoordinate(i int) *Coordinate {
	if i > len(coordinateList.Coordinates)-1 {
		return nil
	}
	coord := coordinateList.Coordinates[i]
	return &coord
}

/**
 * Adds a section of an array of coordinates to the list.
 * @param coord The coordinates
 * @param allowRepeated if set to false, repeated coordinates are collapsed
 * @param start the index to start from
 * @param end the index to add up to but not including
 * @return true (as by general collection contract)
 */
func (coordinateList *CoordinateList) AddCoordinateListRepeatedStartEnd(
	coords []Coordinate,
	allowRepeated bool,
	start uint,
	end uint,
) bool {
	inc := 1
	if start > end {
		inc = -1
	}

	i := start
	for i != end {
		coordinateList.AddCoordinateRepeated(&coords[i], allowRepeated)
		i += uint(inc)
	}
	return true
}

/**
 * Adds an array of coordinates to the list.
 * @param coord The coordinates
 * @param allowRepeated if set to false, repeated coordinates are collapsed
 * @param direction if false, the array is added in reverse order
 * @return true (as by general collection contract)
 */
func (coordinateList *CoordinateList) AddCoordinateListRepeatedDirection(
	coords []Coordinate,
	allowRepeated bool,
	direction bool,
) bool {
	if direction {
		for i := 0; i < len(coords); i++ {
			coordinateList.AddCoordinateRepeated(&coords[i], allowRepeated)
		}
	} else {
		i := len(coords)
		for i >= 0 {
			coordinateList.AddCoordinateRepeated(&coords[i], allowRepeated)
			i -= 1
		}
	}
	return true
}

/**
 * Adds an array of coordinates to the list.
 * @param coord The coordinates
 * @param allowRepeated if set to false, repeated coordinates are collapsed
 * @return true (as by general collection contract)
 */
func (coordinateList *CoordinateList) AddCoordinateListRepeated(
	coords []Coordinate,
	allowRepeated bool,
) bool {
	coordinateList.AddCoordinateListRepeatedDirection(coords, allowRepeated, true)
	return true
}

/**
 * Adds a coordinate to the end of the list.
 *
 * @param coord The coordinates
 * @param allowRepeated if set to false, repeated coordinates are collapsed
 */
func (coordinateList *CoordinateList) AddCoordinateRepeated(coord *Coordinate, allowRepeated bool) {
	// don't add duplicate coordinates
	if !allowRepeated {
		if len(coordinateList.Coordinates) >= 1 {
			index := len(coordinateList.Coordinates) - 1
			last := coordinateList.GetCoordinate(index)
			if last != nil {
				if last.Equals2D(coord) {
					return
				}
			}
		}
	}
	coordinateList.Coordinates = append(coordinateList.Coordinates, *coord)
}

/**
 * Inserts the specified coordinate at the specified position in this list.
 *
 * @param i the position at which to insert
 * @param coord the coordinate to insert
 * @param allowRepeated if set to false, repeated coordinates are collapsed
 */
func (coordinateList *CoordinateList) AddIndexCoordinateRepeated(
	i int,
	coord Coordinate,
	allowRepeated bool,
) {
	// don't add duplicate coordinates
	if !allowRepeated {
		size := len(coordinateList.Coordinates)
		if size > 0 {
			if i > 0 {
				prev := coordinateList.GetCoordinate(i - 1)
				if prev != nil {
					if prev.Equals2D(&coord) {
						return
					}
				}
			}
			if i < size {
				next := coordinateList.GetCoordinate(i)
				if next != nil {
					if next.Equals2D(&coord) {
						return
					}
				}
			}
		}
	}
	coordinateList.Coordinates[i] = coord
}

/** Add an array of coordinates
 * @param coll The coordinates
 * @param allowRepeated if set to false, repeated coordinates are collapsed
 * @return true (as by general collection contract)
 */
func (coordinateList *CoordinateList) AddAll(coll []Coordinate, allowRepeated bool) bool {
	is_changed := false
	for _, coordinate := range coll {
		coordinateList.AddCoordinateRepeated(&coordinate, allowRepeated)
		is_changed = true
	}
	return is_changed
}

/**
 * Ensure this coordList is a ring, by adding the start point if necessary
 */
func (coordinateList *CoordinateList) CloseRing() {
	if len(coordinateList.Coordinates) > 0 {
		duplicate := coordinateList.GetCoordinate(0)
		if duplicate != nil {
			coordinateList.AddCoordinateRepeated(duplicate, false)
		}
	}
}

/** Returns the Coordinates in this collection.
 *
 * @return the coordinates
 */
func (coordinateList *CoordinateList) ToCoordinateArray() []Coordinate {
	new_slice := make([]Coordinate, len(coordinateList.Coordinates))
	copy(new_slice, coordinateList.Coordinates)
	return new_slice
}

/**
 * Creates an array containing the coordinates in this list,
 * oriented in the given direction (forward or reverse).
 *
 * @param isForward true if the direction is forward, false for reverse
 * @return an oriented array of coordinates
 */
func (coordinateList *CoordinateList) ToCoordinateArrayForward(isForward bool) []Coordinate {
	if isForward {
		new_slice := make([]Coordinate, len(coordinateList.Coordinates))
		copy(new_slice, coordinateList.Coordinates)
		return new_slice
	}
	// construct reversed array
	size := len(coordinateList.Coordinates)
	pts := make([]Coordinate, size)
	for i := 0; i < size; i++ {
		coord := coordinateList.GetCoordinate(size - i - 1)
		if coord != nil {
			pts[i] = *coord
		}
	}
	return pts
}
