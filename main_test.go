package main

import (
	"os"
	"testing"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7
	res := MaxInt(a, b)
	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}

// TestMain нужен только для setup/teardown
func TestMain(m *testing.M) {
	// setup code (если нужно)

	code := m.Run() // запускаем все тесты

	// teardown code (если нужно)

	os.Exit(code)
}
