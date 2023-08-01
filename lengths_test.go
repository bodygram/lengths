package lengths

import (
	"math"
	"testing"
)

// floatEqual returns true if a and b are equal up to the 5 decimal; it returns
// false otherwise.
func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-5
}

func TestFromMetrics(t *testing.T) {
	testCases := []struct {
		f      float64
		wantMM Length
		wantCM Length
		wantM  Length
	}{
		{
			f:      0,
			wantMM: 0 * Nanometer,
			wantCM: 0 * Nanometer,
			wantM:  0 * Nanometer,
		},
		{
			f:      1e-9,
			wantMM: 0 * Nanometer,
			wantCM: 0 * Nanometer,
			wantM:  1 * Nanometer,
		},
		{
			f:      1e-8,
			wantMM: 0 * Nanometer,
			wantCM: 0 * Nanometer,
			wantM:  1e1 * Nanometer,
		},
		{
			f:      1e-7,
			wantMM: 0 * Nanometer,
			wantCM: 1 * Nanometer,
			wantM:  1e2 * Nanometer,
		},
		{
			f:      1e-6,
			wantMM: 1 * Nanometer,
			wantCM: 1e1 * Nanometer,
			wantM:  1e3 * Nanometer,
		},
		{
			f:      1e-5,
			wantMM: 1e1 * Nanometer,
			wantCM: 1e2 * Nanometer,
			wantM:  1e4 * Nanometer,
		},
		{
			f:      1e-4,
			wantMM: 1e2 * Nanometer,
			wantCM: 1e3 * Nanometer,
			wantM:  1e5 * Nanometer,
		},
		{
			f:      1e-3,
			wantMM: 1e3 * Nanometer,
			wantCM: 1e4 * Nanometer,
			wantM:  1e6 * Nanometer,
		},
		{
			f:      1e-2,
			wantMM: 1e4 * Nanometer,
			wantCM: 1e5 * Nanometer,
			wantM:  1e7 * Nanometer,
		},
		{
			f:      1e-1,
			wantMM: 1e5 * Nanometer,
			wantCM: 1e6 * Nanometer,
			wantM:  1e8 * Nanometer,
		},
		{
			f:      1,
			wantMM: 1e6 * Nanometer,
			wantCM: 1e7 * Nanometer,
			wantM:  1e9 * Nanometer,
		},
		{
			f:      1e1,
			wantMM: 1e7 * Nanometer,
			wantCM: 1e8 * Nanometer,
			wantM:  1e10 * Nanometer,
		},
		{
			f:      1e2,
			wantMM: 1e8 * Nanometer,
			wantCM: 1e9 * Nanometer,
			wantM:  1e11 * Nanometer,
		},
		{
			f:      1e3,
			wantMM: 1e9 * Nanometer,
			wantCM: 1e10 * Nanometer,
			wantM:  1e12 * Nanometer,
		},
	}

	for _, tc := range testCases {
		if got := Millimeters(tc.f); got != tc.wantMM {
			t.Errorf("FromMillimeters(): got %q, want %q", got, tc.wantMM)
		}
		if got := Centimeters(tc.f); got != tc.wantCM {
			t.Errorf("FromCentimeters(): got %q, want %q", got, tc.wantCM)
		}
		if got := Meters(tc.f); got != tc.wantM {
			t.Errorf("FromMeters(): got %q, want %q", got, tc.wantM)
		}
	}
}

func TestFromInches(t *testing.T) {
	testCases := []struct {
		f    float64
		want Length
	}{
		{
			f:    0,
			want: 0 * Nanometer,
		},
		{
			f:    1e-7,
			want: 2 * Nanometer,
		},
		{
			f:    1e-6,
			want: 25 * Nanometer,
		},
		{
			f:    1e-5,
			want: 254 * Nanometer,
		},
		{
			f:    1e-4,
			want: 2540 * Nanometer,
		},
		{
			f:    1e-3,
			want: 25400 * Nanometer,
		},
		{
			f:    1e-2,
			want: 254000 * Nanometer,
		},
		{
			f:    1e-1,
			want: 2540000 * Nanometer,
		},
		{
			f:    1,
			want: 25400000 * Nanometer,
		},
		{
			f:    1e1,
			want: 254000000 * Nanometer,
		},
		{
			f:    1e2,
			want: 2540000000 * Nanometer,
		},
	}

	for _, tc := range testCases {
		if got := Inches(tc.f); got != tc.want {
			t.Errorf("FromInches(): got %q, want %q", got, tc.want)
		}
	}
}

func TestFromFeet(t *testing.T) {
	testCases := []struct {
		f    float64
		want Length
	}{
		{
			f:    0,
			want: 0 * Nanometer,
		},
		{
			f:    1e-7,
			want: 30 * Nanometer,
		},
		{
			f:    1e-6,
			want: 304 * Nanometer,
		},
		{
			f:    1e-5,
			want: 3048 * Nanometer,
		},
		{
			f:    1e-4,
			want: 30480 * Nanometer,
		},
		{
			f:    1e-3,
			want: 304800 * Nanometer,
		},
		{
			f:    1e-2,
			want: 3048000 * Nanometer,
		},
		{
			f:    1e-1,
			want: 30480000 * Nanometer,
		},
		{
			f:    1,
			want: 304800000 * Nanometer,
		},
		{
			f:    1e1,
			want: 3048000000 * Nanometer,
		},
		{
			f:    1e2,
			want: 30480000000 * Nanometer,
		},
	}

	for _, tc := range testCases {
		if got := Feet(tc.f); got != tc.want {
			t.Errorf("FromFeet(): got %q, want %q", got, tc.want)
		}
	}
}

func TestFeetAndInches(t *testing.T) {
	testCases := []struct {
		l          Length
		wantFeet   float64
		wantInches float64
	}{
		{
			l:          0 * Nanometer,
			wantFeet:   0,
			wantInches: 0,
		},
		{
			l:          254 * Micrometer,
			wantFeet:   0,
			wantInches: 0.01,
		},
		{
			l:          178 * Centimeter,
			wantFeet:   5,
			wantInches: 10.07874,
		},
	}

	for _, tc := range testCases {
		gotFeet, gotInches := tc.l.FeetAndInches()

		if !floatEqual(gotFeet, tc.wantFeet) || !floatEqual(gotInches, tc.wantInches) {
			t.Errorf(
				"String(): got %f, %f, want %f, %f",
				gotFeet,
				gotInches,
				tc.wantFeet,
				tc.wantInches,
			)
		}
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		l    Length
		want string
	}{
		{
			l:    0 * Nanometer,
			want: "0nm",
		},
		{
			l:    1 * Nanometer,
			want: "1nm",
		},
		{
			l:    12 * Nanometer,
			want: "12nm",
		},
		{
			l:    123 * Nanometer,
			want: "123nm",
		},
		{
			l:    1234 * Nanometer,
			want: "1.234μm",
		},
		{
			l:    12345 * Nanometer,
			want: "12.345μm",
		},
		{
			l:    123456 * Nanometer,
			want: "123.456μm",
		},
		{
			l:    1234567 * Nanometer,
			want: "1.234567mm",
		},
		{
			l:    76543210 * Nanometer,
			want: "7.654321cm",
		},
		{
			l:    765432100 * Nanometer,
			want: "76.54321cm",
		},
		{
			l:    7654321000 * Nanometer,
			want: "7.654321m",
		},
		{
			l:    76543210000 * Nanometer,
			want: "76.54321m",
		},
		{
			l:    765432100000 * Nanometer,
			want: "765.4321m",
		},
		{
			l:    7654321000000 * Nanometer,
			want: "7654.321m",
		},
		{
			l:    76543210000000 * Nanometer,
			want: "76543.21m",
		},
		{
			l:    765432100000000 * Nanometer,
			want: "765432.1m",
		},
		{
			l:    7654321000000000 * Nanometer,
			want: "7654321m",
		},
	}

	for _, tc := range testCases {
		if got := tc.l.String(); got != tc.want {
			t.Errorf("String(): got %q, want %q", got, tc.want)
		}
	}
}
