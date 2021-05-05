package main

import "testing"

func TestGetData(t *testing.T) {
	for i:=0; i<100; i++ {
		if GetData(i) != nil {
			t.Error("Unespected value")
		}
	}
}