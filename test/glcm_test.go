package main

import (
	"testing"
	"../src/ProcessStrategy"
)

func TestGlcmSetParameters(t *testing.T) {
	var glcm process.GLCM

	err := glcm.SetParameters(1,0)

	if err != nil {
		t.Error("Unexpected value")
	}
}