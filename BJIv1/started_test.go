package main

import (
	"runtime"
	"testing"
)

func TestStart(t *testing.T) {
	want := "Starting..."
	got := startApp()

	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestOS(t *testing.T) {
	want1 := "linux"
	want2 := "darwin"
	want3 := "windows"
	got := clearScreen(runtime.GOOS)
	if want1 != got || want2 != got || want3 != got {
		t.Errorf("wanted %v, %v, or %v, but got %v", want1, want2, want3, got)
	}
}
