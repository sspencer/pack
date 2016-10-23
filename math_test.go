package pack

import (
	"testing"
)

func TestMathMin(t *testing.T) {
	if min(5, 10) != 5 {
		t.Error("WOW! How did you screw up min???")
	}
}

func TestMathMax(t *testing.T) {
	if max(5, 10) != 10 {
		t.Error("WOW! How did you screw up max???")
	}
}
