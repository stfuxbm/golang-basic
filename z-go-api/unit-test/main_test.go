package main

import "testing"

func TestTambah(t *testing.T) {
	hasil := Tambah(2, 4)
	if hasil != 6 {
		t.Errorf("Expected 6, got %d", hasil)
	}
}
