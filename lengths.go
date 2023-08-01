// Package lengths provides functionality for measuring and displaying lengths.

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

// Common length units.
//
// To count the number of units in a Length, divide it by the unit:
//
//	meter := 10 * units.Meter
//	fmt.Print(uint64(meter/units.Millimeter)) // prints 10000
//
// To convert an integer number of units to a Length, multiply it by the unit:
//
//	meters := 10
//	fmt.Print(units.Length(meters)*units.Meter) // prints 10m
//
// To convert a floating point number of units to a Length, use the dedicated
// FromXXX function where XXX is the unit.
const (
	Nanometer  Length = 1
	Micrometer        = 1e3 * Nanometer
	Millimeter        = 1e6 * Nanometer
	Centimeter        = 1e7 * Nanometer
	Meter             = 1e9 * Nanometer
	Kilometer         = 1e12 * Nanometer
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

// Kilometers returns the length as a floating point number of kilometers.
func (l Length) Kilometers() float64 {
	return float64(l/Kilometer) + float64(l%Kilometer)/1e12
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

// Kilometers returns a length from a floating point number of kilometers. The
// length's precision is floored to the closest nanometer.
func Kilometers(f float64) Length {
	return Length(f * float64(Kilometer))
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
		return "0"
	case l < Micrometer:
		return fmt.Sprintf("%dnm", l/Nanometer)
	case l < Millimeter:
		return fmt.Sprintf("%sμm", formatFloat(l.Micrometers()))
	case l < Centimeter:
		return fmt.Sprintf("%smm", formatFloat(l.Millimeters()))
	case l < Meter:
		return fmt.Sprintf("%scm", formatFloat(l.Centimeters()))
	case l < Kilometer:
		return fmt.Sprintf("%sm", formatFloat(l.Meters()))
	default:
		return fmt.Sprintf("%skm", formatFloat(l.Kilometers()))
	}
}
