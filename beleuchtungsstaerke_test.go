package main

import (
	"math"
	"strconv"
	"testing"
)

func TestCalcIlluminance(t *testing.T) {
	tests := []struct {
		name   string
		fields IlluminanceParameters
		want   float64
	}{
		{
			name: "Test1",
			fields: IlluminanceParameters{
				OpeningAngle:   sr2rad(1.67 * math.Pi),
				LuminousFlux:   3545,
				Distance:       1.67,
				AngleOfIncline: math.Atan(1.15 / 1.67),
			},
			want: 199.543947,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			methodResult := tt.fields.CalcIlluminance()
			if strconv.FormatFloat(methodResult, 'f', 6, 64) != strconv.FormatFloat(tt.want, 'f', 6, 64) {
				t.Errorf("Expected: %f, want: %f", tt.want, methodResult)
			}
		})

	}
}
