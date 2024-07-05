package geos

const (
	/** Standard ordinate index value for, where X is 0 */
	COORDINATE_X = 0

	/** Standard ordinate index value for, where Y is 1 */
	COORDINATE_Y = 1

	/**
	 * Standard ordinate index value for, where Z is 2.
	 *
	 * <p>This constant assumes XYZM coordinate sequence definition, please check this assumption
	 * using {@link CoordinateSequence#getDimension()} and {@link CoordinateSequence#getMeasures()}
	 * before use.
	 */
	COORDINATE_Z = 2

	/**
	 * Standard ordinate index value for, where M is 3.
	 *
	 * <p>This constant assumes XYZM coordinate sequence definition, please check this assumption
	 * using {@link CoordinateSequence#getDimension()} and {@link CoordinateSequence#getMeasures()}
	 * before use.
	 */
	COORDINATE_M = 3
)
