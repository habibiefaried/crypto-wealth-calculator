package main

import (
	"fmt"
	"math"
	"testing"
)

func TestConvertStringToFloat(t *testing.T) {
	tests := []struct {
		numStr       string
		decimalShift int
		want         float64
	}{
		{"17013825326990646", 18, 0.017013825326990646},
		{"1000000000000000000", 18, 1.0},
		{"123456789012345678", 18, 0.123456789012345678},
		{"0", 18, 0.0},
		{"260655465325976763", 18, 0.260655465325976763},
		{"115956089072092773783", 18, 115.956089072092773783},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		got, err := ConvertStringToFloat(tt.numStr, tt.decimalShift)
		if err != nil {
			t.Errorf("ConvertStringToFloat(%s, %d) error = %v, want no error", tt.numStr, tt.decimalShift, err)
		}
		// Since we are dealing with floating-point numbers, use a small delta for comparison to account for any possible rounding errors.
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("ConvertStringToFloat(%s, %d) = %v, want %v", tt.numStr, tt.decimalShift, got, tt.want)
		}

		fmt.Printf("%.5f\n", got)
	}
}
