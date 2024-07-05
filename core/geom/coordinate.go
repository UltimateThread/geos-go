package geos

import (
	"fmt"
	"math"

	constants "github.com/UltimateThread/geos-go/core/constants"
	util "github.com/UltimateThread/geos-go/core/util"
)

type Coordinate struct {
	/**
	 * The x-ordinate.
	 */
	X float64

	/**
	 * The y-ordinate.
	 */
	Y float64

	/**
	 * The z-ordinate.
	 */
	Z float64

	/**
	 * The m-ordinate.
	 */
	M float64
}

/**
 *  Constructs a <code>Coordinate</code> at (0,0,NaN).
 */
func DefaultCoordinateXY() *Coordinate {
	c := new(Coordinate)
	c.X = 0
	c.Y = 0
	c.Z = math.NaN()
	c.M = math.NaN()
	return c
}

func DefaultCoordinateXYZ() *Coordinate {
	c := new(Coordinate)
	c.X = 0
	c.Y = 0
	c.Z = 0
	c.M = math.NaN()
	return c
}

func DefaultCoordinateXYM() *Coordinate {
	c := new(Coordinate)
	c.X = 0
	c.Y = 0
	c.Z = math.NaN()
	c.M = 0
	return c
}

func DefaultCoordinateXYZM() *Coordinate {
	c := new(Coordinate)
	c.X = 0
	c.Y = 0
	c.Z = 0
	c.M = 0
	return c
}

/**
 *  Constructs a <code>Coordinate</code> at (x,y,NaN).
 *
 *@param  x  the x-value
 *@param  y  the y-value
 */
func NewCoordinateXY(x float64, y float64) *Coordinate {
	c := new(Coordinate)
	c.X = x
	c.Y = y
	c.Z = math.NaN()
	c.M = math.NaN()
	return c
}

/**
 *  Constructs a <code>Coordinate</code> at (x,y,z).
 *
 *@param  x  the x-ordinate
 *@param  y  the y-ordinate
 *@param  z  the z-ordinate
 */
func NewCoordinateXYZ(x float64, y float64, z float64) *Coordinate {
	c := new(Coordinate)
	c.X = x
	c.Y = y
	c.Z = z
	c.M = math.NaN()
	return c
}

/**
 *  Constructs a <code>Coordinate</code> at (x,y,m).
 *
 *@param  x  the x-ordinate
 *@param  y  the y-ordinate
 *@param  m  the m-ordinate
 */
func NewCoordinateXYM(x float64, y float64, m float64) *Coordinate {
	c := new(Coordinate)
	c.X = x
	c.Y = y
	c.Z = math.NaN()
	c.M = m
	return c
}

/**
 *  Constructs a <code>Coordinate</code> at (x,y,z, m).
 *
 *@param  x  the x-ordinate
 *@param  y  the y-ordinate
 *@param  z  the z-ordinate
 *@param  m  the m-ordinate
 */
func NewCoordinateXYZM(x float64, y float64, z float64, m float64) *Coordinate {
	c := new(Coordinate)
	c.X = x
	c.Y = y
	c.Z = z
	c.M = m
	return c
}

/**
 *  Constructs a <code>Coordinate</code> having the same (x,y,z) values as
 *  <code>other</code>.
 *
 *@param  c  the <code>Coordinate</code> to copy.
 */
func NewCoordinateFromCoordinate(existing *Coordinate) *Coordinate {
	c := new(Coordinate)
	c.X = existing.X
	c.Y = existing.Y
	c.Z = existing.Z
	c.M = existing.M
	return c
}

func (coord *Coordinate) Clone() *Coordinate {
	return NewCoordinateFromCoordinate(coord)
}

func (coord *Coordinate) SetCoordinate(other *Coordinate) {
	coord.X = other.X
	coord.Y = other.Y
	coord.Z = other.Z
	coord.M = other.M
}

/**
 * Gets the ordinate value for the given index.
 *
 * The base implementation supports values for the index are
 * {@link #X}, {@link #Y}, and {@link #Z}.
 *
 * @param ordinateIndex the ordinate index
 * @return the value of the ordinate
 * @throws IllegalArgumentException if the index is not valid
 */
func (coord *Coordinate) GetOrdinate(ordinateIndex int) *float64 {
	switch ordinateIndex {
	case constants.COORDINATE_X:
		return &coord.X
	case constants.COORDINATE_Y:
		return &coord.Y
	case constants.COORDINATE_Z:
		return &coord.Z
	case constants.COORDINATE_M:
		return &coord.M
	default:
		return nil
	}
}

/**
 * Sets the ordinate for the given index
 * to a given value.
 *
 * The base implementation supported values for the index are
 * {@link #X}, {@link #Y}, and {@link #Z}.
 *
 * @param ordinateIndex the ordinate index
 * @param value the value to set
 * @throws IllegalArgumentException if the index is not valid
 */
func (coord *Coordinate) SetOrdinate(ordinateIndex int, value float64) {
	switch ordinateIndex {
	case constants.COORDINATE_X:
		coord.X = value
	case constants.COORDINATE_Y:
		coord.Y = value
	case constants.COORDINATE_Z:
		coord.Z = value
	case constants.COORDINATE_M:
		coord.M = value
	}
}

/**
 * Tests if the coordinate has valid X and Y ordinate values.
 * An ordinate value is valid iff it is finite.
 *
 * @return true if the coordinate is valid
 * @see Double#isFinite(double)
 */
func (coord *Coordinate) IsValid() bool {
	if math.IsInf(coord.X, 0) {
		return false
	}
	if math.IsInf(coord.Y, 0) {
		return false
	}
	return true
}

func (coord *Coordinate) IsXY() bool {
	return !math.IsNaN(coord.X) && !math.IsNaN(coord.Y) && math.IsNaN(coord.Z) && math.IsNaN(coord.M)
}

func (coord *Coordinate) IsXYM() bool {
	return !math.IsNaN(coord.X) && !math.IsNaN(coord.Y) && math.IsNaN(coord.Z) && !math.IsNaN(coord.M)
}

func (coord *Coordinate) IsXYZM() bool {
	return !math.IsNaN(coord.X) && !math.IsNaN(coord.Y) && !math.IsNaN(coord.Z) && !math.IsNaN(coord.M)
}

/**
 *  Returns whether the planar projections of the two <code>Coordinate</code>s
 *  are equal.
 *
 *@param  other  a <code>Coordinate</code> with which to do the 2D comparison.
 *@return        <code>true</code> if the x- and y-coordinates are equal; the
 *      z-coordinates do not have to be equal.
 */
func (coord *Coordinate) Equals2D(other *Coordinate) bool {
	if coord.X != other.X {
		return false
	}
	if coord.Y != other.Y {
		return false
	}
	return true
}

/**
 * Tests if another Coordinate has the same values for the X and Y ordinates,
 * within a specified tolerance value.
 * The Z ordinate is ignored.
 *
 *@param c a <code>Coordinate</code> with which to do the 2D comparison.
 *@param tolerance the tolerance value to use
 *@return true if <code>other</code> is a <code>Coordinate</code>
 *      with the same values for X and Y.
 */
func (coord *Coordinate) Equals2DWithTolerance(c *Coordinate, tolerance float64) bool {
	if !util.EqualsWithTolerance(coord.X, c.X, tolerance) {
		return false
	}
	if !util.EqualsWithTolerance(coord.Y, c.Y, tolerance) {
		return false
	}
	return true
}

/**
 * Tests if another coordinate has the same values for the X, Y and Z ordinates.
 *
 *@param other a <code>Coordinate</code> with which to do the 3D comparison.
 *@return true if <code>other</code> is a <code>Coordinate</code>
 *      with the same values for X, Y and Z.
 */
func (coord *Coordinate) Equals3D(other *Coordinate) bool {
	return (coord.X == other.X) && (coord.Y == other.Y) && ((coord.Z == other.Z) || (math.IsNaN(coord.Z) && math.IsNaN(other.Z)))
}

/**
 * Tests if another coordinate has the same value for Z, within a tolerance.
 *
 * @param c a coordinate
 * @param tolerance the tolerance value
 * @return true if the Z ordinates are within the given tolerance
 */
func (coord *Coordinate) EqualInZ(c *Coordinate, tolerance float64) bool {
	return util.EqualsWithTolerance(coord.Z, c.Z, tolerance)
}

/**
 * Tests if another coordinate has the same value for Z, within a tolerance.
 *
 * @param c a coordinate
 * @param tolerance the tolerance value
 * @return true if the Z ordinates are within the given tolerance
 */
func (coord *Coordinate) EqualInCoordinateZ(c *Coordinate, tolerance float64) bool {
	return util.EqualsWithTolerance(coord.Z, c.Z, tolerance)
}

/**
 *  Compares this {@link Coordinate} with the specified {@link Coordinate} for order.
 *  This method ignores the z value when making the comparison.
 *  Returns:
 *  <UL>
 *    <LI> -1 : this.x &lt; other.x || ((this.x == other.x) &amp;&amp; (this.y &lt; other.y))
 *    <LI> 0 : this.x == other.x &amp;&amp; this.y = other.y
 *    <LI> 1 : this.x &gt; other.x || ((this.x == other.x) &amp;&amp; (this.y &gt; other.y))
 *
 *  </UL>
 *  Note: This method assumes that ordinate values
 * are valid numbers.  NaN values are not handled correctly.
 *
 *@param  o  the <code>Coordinate</code> with which this <code>Coordinate</code>
 *      is being compared
 *@return    -1, zero, or 1 as this <code>Coordinate</code>
 *      is less than, equal to, or greater than the specified <code>Coordinate</code>
 */
func (coord *Coordinate) CompareTo(other *Coordinate) int {
	if coord.X < other.X {
		return -1
	}
	if coord.X > other.X {
		return 1
	}
	if coord.Y < other.Y {
		return -1
	}
	if coord.Y > other.Y {
		return 1
	}
	return 0
}

/**
 * Computes the 2-dimensional Euclidean distance to another location.
 * The Z-ordinate is ignored.
 *
 * @param c a point
 * @return the 2-dimensional Euclidean distance between the locations
 */
func (coord *Coordinate) Distance(c *Coordinate) float64 {
	dx := coord.X - c.X
	dy := coord.Y - c.Y

	return math.Hypot(dx, dy)
}

/**
 * Computes the 3-dimensional Euclidean distance to another location.
 *
 * @param c a coordinate
 * @return the 3-dimensional Euclidean distance between the locations
 */
func (coord *Coordinate) Distance3D(c *Coordinate) float64 {
	dx := coord.X - c.X
	dy := coord.Y - c.Y
	dz := coord.Z - c.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (coord *Coordinate) ToString() string {
	return fmt.Sprintf("(%f, %f, %f)", coord.X, coord.Y, coord.Z)
}
