package main

import (
    "testing"
    "runtime"
)

func TestStart(t *testing.T) {
	want := "Starting..."
	got := startApp()

	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestOS(t *testing.T){
    want :="linux"
    got := clearScreen(runtime.GOOS)
    if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}