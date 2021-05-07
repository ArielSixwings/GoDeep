package main

import (
	"testing"
	"../src/processStrategy"
	"gocv.io/x/gocv"
)

func TestNormalizeSetParameters(t *testing.T) {
	var (normalize process.Normalize
		normtype gocv.NormType = gocv.NormMinMax
	)

	err := normalize.SetParameters(0.0, 255.0, normtype)

	if err != nil {
		t.Error("Unexpected value")
	}
}