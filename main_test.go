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

func TestMain(m *testing.M) {
	// можно добавить setup, если нужно

	code := m.Run() // запускаем все тесты

	// можно добавить teardown, если нужно

	os.Exit(code)
}
