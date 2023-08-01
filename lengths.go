// Package lengths provides a Length type and common length units.
package lengths

import (
	"fmt"
	"strconv"
)

// A Length represents the extent of something from end to end as an uint64
// nanometer count (as a length cannot be negative). The representation limits
// the largest representable length to approximately 18 gigameters (which is
// more than 40 times the distance between Earth and the Moon).
type Length uint64

// Common length units. There is no definition for units larger than meter as
// Bodygram doesn't have a need for it at the time this package was written.
//
// To count the number of units in a Length, divide it by the unit:
//
//	meter := units.Meter
//	fmt.Print(uint64(meter/units.Millimeter)) // prints 1000
//
// To convert an integer number of units to a Length, multiply it by the unit:
//
//	meters := 10
//	fmt.Print(units.Length(meters)*units.Meter) // prints 10.000m
//
// To convert a floating point number of units to a Length, use the dedicated
// FromXXX function where XXX is the unit.
const (
	Nanometer  Length = 1
	Micrometer        = 1e3 * Nanometer
	Millimeter        = 1e6 * Nanometer
	Centimeter        = 1e7 * Nanometer
	Meter             = 1e9 * Nanometer
	Inch              = 254e5 * Nanometer
	Foot              = 3048e5 * Nanometer
)

// Micrometers returns the length as a floating point number of micrometers.
func (l Length) Micrometers() float64 {
	return float64(l/Micrometer) + float64(l%Micrometer)/1e3
}

// Millimeters returns the length as a floating point number of millimeters.
func (l Length) Millimeters() float64 {
	return float64(l/Millimeter) + float64(l%Millimeter)/1e6
}

// Centimeters returns the length as a floating point number of centimeters.
func (l Length) Centimeters() float64 {
	return float64(l/Centimeter) + float64(l%Centimeter)/1e7
}

// Meters returns the length as a floating point number of meters.
func (l Length) Meters() float64 {
	return float64(l/Meter) + float64(l%Meter)/1e9
}

// Inches returns the length as a floating point number of inches.
func (l Length) Inches() float64 {
	return float64(l/Inch) + float64(l%Inch)/254e5
}

// FeetAndInches returns the length as a number of feet and inches.
//
// The number of feet is always an integer but is returned as a floating point
// number for consistency with how inches are returned.
func (l Length) FeetAndInches() (float64, float64) {
	feet := l / Foot
	inches := l % Foot
	return float64(feet), inches.Inches()
}

// Millimeters returns a length from a floating point number of millimeters.
// The length's precision is floored to the closest nanometer.
func Millimeters(f float64) Length {
	return Length(f * float64(Millimeter))
}

// Centimeters returns a length from a floating point number of centimeters.
// The length's precision is floored to the closest nanometer.
func Centimeters(f float64) Length {
	return Length(f * float64(Centimeter))
}

// Meters returns a length from a floating point number of meters. The
// length's precision is floored to the closest nanometer.
func Meters(f float64) Length {
	return Length(f * float64(Meter))
}

// Inches returns a length from a floating point number of inches. The
// length's precision is floored to the closest nanometer.
func Inches(f float64) Length {
	return Length(f * float64(Inch))
}

// Feet returns a length from a floating point number of feet. The length's
// precision is floored to the closest nanometer.
func Feet(f float64) Length {
	return Length(f * float64(Foot))
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func (l Length) String() string {
	switch {
	case l < Nanometer:
		return "0nm"
	case l < Micrometer:
		return fmt.Sprintf("%dnm", l/Nanometer)
	case l < Millimeter:
		return fmt.Sprintf("%sμm", formatFloat(l.Micrometers()))
	case l < Centimeter:
		return fmt.Sprintf("%smm", formatFloat(l.Millimeters()))
	case l < Meter:
		return fmt.Sprintf("%scm", formatFloat(l.Centimeters()))
	default:
		return fmt.Sprintf("%sm", formatFloat(l.Meters()))
	}
}