package e_test

import (
	"errors"
	"testing"

	"github.com/julienbreux/e"
)

// TestRror test the Rror function
func TestRror(t *testing.T) {
	err := e.Rror("...", errors.New("Oops"))

	if err.Error() != "Oops" {
		t.Error("Oops...")
	}
}
