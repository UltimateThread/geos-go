package tests

import (
	"testing"

	geom "github.com/UltimateThread/geos-go/core/geom"
	"github.com/stretchr/testify/assert"
)

func TestForward(t *testing.T) {
	coords1 := []float64{0., 0., 1., 1., 2., 2.}
	coords2 := []float64{0., 0., 1., 1., 2., 2.}
	coordList := coord_list(coords1)
	check_value(t,
		coordList.ToCoordinateArrayForward(true),
		coords2,
	)
}

func TestReverse(t *testing.T) {
	coords1 := []float64{0., 0., 1., 1., 2., 2.}
	coords2 := []float64{2., 2., 1., 1., 0., 0.}
	coordList := coord_list(coords1)
	check_value(
		t,
		coordList.ToCoordinateArrayForward(false),
		coords2,
	)
}

func TestReverseEmpty(t *testing.T) {
	var empty []float64
	var empty2 []float64
	coordList := coord_list(empty)
	check_value(
		t,
		coordList.ToCoordinateArrayForward(false),
		empty2,
	)
}

func check_value(t *testing.T, coord_array []geom.Coordinate, ords []float64) {
	assert.Equal(t, len(coord_array)*2, len(ords))

	i := 0
	for i < len(coord_array) {
		pt := coord_array[i]
		assert.Equal(t, pt.X, ords[2*i])
		assert.Equal(t, pt.Y, ords[2*i+1])
		i += 2
	}
}

func coord_list(ords []float64) geom.CoordinateList {
	cl := geom.DefaultCoordinateList()
	i := 0
	for i < len(ords) {
		coord := geom.NewCoordinateXY(ords[i], ords[i+1])
		cl.AddCoordinateRepeated(coord, false)
		i += 2
	}
	return *cl
}
