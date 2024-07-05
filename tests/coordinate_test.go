package tests

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	constants "github.com/UltimateThread/geos-go/core/constants"
	geom "github.com/UltimateThread/geos-go/core/geom"
)

func TestConstructor3D(t *testing.T) {
	c := geom.NewCoordinateXYZ(350.2, 4566.8, 5266.3)
	assert.Equal(t, c.X, 350.2)
	assert.Equal(t, c.Y, 4566.8)
	assert.Equal(t, c.Z, 5266.3)
}

func TestConstructor2D(t *testing.T) {
	c := geom.NewCoordinateXY(350.2, 4566.8)
	assert.Equal(t, c.X, 350.2)
	assert.Equal(t, c.Y, 4566.8)
	assert.True(t, math.IsNaN(c.Z))
}

func TestDefaultConstructor(t *testing.T) {
	c := geom.DefaultCoordinateXY()
	assert.Equal(t, c.X, 0.0)
	assert.Equal(t, c.Y, 0.0)
	assert.True(t, math.IsNaN(c.Z))
}

func TestCopyConstructor3D(t *testing.T) {
	orig := geom.NewCoordinateXYZ(350.2, 4566.8, 5266.3)
	c := geom.NewCoordinateFromCoordinate(orig)
	assert.Equal(t, c.X, 350.2)
	assert.Equal(t, c.Y, 4566.8)
	assert.Equal(t, c.Z, 5266.3)
}

func TestSetCoordinate(t *testing.T) {
	orig := geom.NewCoordinateXYZ(350.2, 4566.8, 5266.3)
	c := geom.DefaultCoordinateXY()
	c.SetCoordinate(orig)
	assert.Equal(t, c.X, 350.2)
	assert.Equal(t, c.Y, 4566.8)
	assert.Equal(t, c.Z, 5266.3)
}

func TestGetOrdinate(t *testing.T) {
	c := geom.NewCoordinateXYZ(350.2, 4566.8, 5266.3)
	assert.Equal(t, *c.GetOrdinate(constants.COORDINATE_X), 350.2)
	assert.Equal(t, *c.GetOrdinate(constants.COORDINATE_Y), 4566.8)
	assert.Equal(t, *c.GetOrdinate(constants.COORDINATE_Z), 5266.3)
}

func TestSetOrdinate(t *testing.T) {
	c := geom.DefaultCoordinateXY()
	c.SetOrdinate(constants.COORDINATE_X, 111.)
	c.SetOrdinate(constants.COORDINATE_Y, 222.)
	c.SetOrdinate(constants.COORDINATE_Z, 333.)
	assert.Equal(t, *c.GetOrdinate(constants.COORDINATE_X), 111.0)
	assert.Equal(t, *c.GetOrdinate(constants.COORDINATE_Y), 222.0)
	assert.Equal(t, *c.GetOrdinate(constants.COORDINATE_Z), 333.0)
}

func TestEquals2D(t *testing.T) {
	c1 := geom.NewCoordinateXYZ(1, 2, 3)
	c2 := geom.NewCoordinateXYZ(1, 2, 3)
	assert.True(t, c1.Equals2D(c2))

	c3 := geom.NewCoordinateXYZ(1, 22, 3)
	assert.True(t, !c1.Equals2D(c3))
}

func TestEquals3D(t *testing.T) {
	c1 := geom.NewCoordinateXYZ(1., 2., 3.)
	c2 := geom.NewCoordinateXYZ(1., 2., 3.)
	assert.True(t, c1.Equals3D(c2))

	c3 := geom.NewCoordinateXYZ(1., 22., 3.)
	assert.True(t, !c1.Equals3D(c3))
}

func TestEquals2DWithinTolerance(t *testing.T) {
	c := geom.NewCoordinateXYZ(100.0, 200.0, 50.0)
	a_bit_off := geom.NewCoordinateXYZ(100.1, 200.1, 50.0)
	assert.True(t, c.Equals2DWithTolerance(a_bit_off, 0.2))
}

func TestEqualsInZ(t *testing.T) {
	c := geom.NewCoordinateXYZ(100.0, 200.0, 50.0)
	with_same_z := geom.NewCoordinateXYZ(100.1, 200.1, 50.1)
	assert.True(t, c.EqualInZ(with_same_z, 0.2))
}

func TestCompareTo(t *testing.T) {
	lowest := geom.NewCoordinateXYZ(10.0, 100.0, 50.0)
	highest := geom.NewCoordinateXYZ(20.0, 100.0, 50.0)
	equal_to_highest := geom.NewCoordinateXYZ(20.0, 100.0, 50.0)
	higher_still := geom.NewCoordinateXYZ(20.0, 200.0, 50.0)

	assert.Equal(t, -1, lowest.CompareTo(highest))
	assert.Equal(t, 1, highest.CompareTo(lowest))
	assert.Equal(t, -1, highest.CompareTo(higher_still))
	assert.Equal(t, 0, highest.CompareTo(equal_to_highest))
}

func TestToString(t *testing.T) {
	expected_result := "(100.100000, 200.200000, 50.300000)"
	actual_result := geom.NewCoordinateXYZ(100.1, 200.2, 50.3).ToString()
	assert.Equal(t, expected_result, actual_result)
}

func TestClone(t *testing.T) {
	c := geom.NewCoordinateXYZ(100.0, 200.0, 50.0)
	clone := c.Clone()
	assert.True(t, c.Equals3D(clone))
}

func TestDistance(t *testing.T) {
	coord1 := geom.NewCoordinateXYZ(0.0, 0.0, 0.0)
	coord2 := geom.NewCoordinateXYZ(100.0, 200.0, 50.0)
	distance := coord1.Distance(coord2)
	assert.Equal(t, distance, 223.60679774997897)
}

func TestDistance3D(t *testing.T) {
	coord1 := geom.NewCoordinateXYZ(0.0, 0.0, 0.0)
	coord2 := geom.NewCoordinateXYZ(100.0, 200.0, 50.0)
	distance := coord1.Distance3D(coord2)
	assert.Equal(t, distance, 229.128784747792)
}

func TestCoordinateXY(t *testing.T) {
	xy := geom.NewCoordinateXY(1.0, 1.0)          // 2D
	coord := geom.NewCoordinateFromCoordinate(xy) // copy
	assert.True(t, xy.X == coord.X && xy.Y == coord.Y)

	coord = geom.NewCoordinateXYZ(1.0, 1.0, 1.0) // 2.5d
	xy = geom.NewCoordinateFromCoordinate(coord) // copy
	assert.True(t, xy.X == coord.X && xy.Y == coord.Y)
}

func TestCoordinateXYM(t *testing.T) {
	xym := geom.DefaultCoordinateXYM()

	xym.M = 1.0
	assert.Equal(t, 1.0, xym.M)

	coord := geom.NewCoordinateFromCoordinate(xym) // copy
	assert.True(t, xym.X == coord.X && xym.Y == coord.Y)

	coord = geom.NewCoordinateXYZ(1.0, 1.0, 1.0)  // 2.5d
	xym = geom.NewCoordinateFromCoordinate(coord) // copy
	assert.True(t, xym.X == coord.X && xym.Y == coord.Y)
}

func TestCoordinateXYZM(t *testing.T) {
	xyzm := geom.DefaultCoordinateXYZM()
	xyzm.Z = 1.0
	assert.Equal(t, 1.0, xyzm.Z)
	xyzm.M = 1.0
	assert.Equal(t, 1.0, xyzm.M)

	coord := geom.NewCoordinateFromCoordinate(xyzm) // copy
	assert.Equal(t, xyzm.X, coord.X)
	assert.Equal(t, xyzm.Y, coord.Y)
	assert.Equal(t, xyzm.Z, coord.Z)
	assert.Equal(t, xyzm.M, coord.M)
	assert.True(t, xyzm.EqualInCoordinateZ(coord, 0.000001))

	coord = geom.NewCoordinateXYZ(1.0, 1.0, 1.0)   // 2.5d
	xyzm = geom.NewCoordinateFromCoordinate(coord) // copy
	assert.Equal(t, xyzm.X, coord.X)
	assert.Equal(t, xyzm.Y, coord.Y)
	assert.Equal(t, xyzm.Z, coord.Z)
	assert.True(t, math.IsNaN(xyzm.M) && math.IsNaN(coord.M))
	assert.True(t, xyzm.EqualInCoordinateZ(coord, 0.000001))
}
