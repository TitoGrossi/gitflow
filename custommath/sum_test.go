package custommath_test

import (
	"testing"

	"github.com/TitoGrossi/gitflow/custommath"
)

func TestSum(t *testing.T) {
	total := custommath.Sum(15, 15)
	if total != 30 {
		t.Errorf("Got unexpected result: Got %d, Expected %d", total, 30)
	}
}

func TestSubtraction(t *testing.T) {
	total := custommath.Subtraction(15, 10)
	if total != 5 {
		t.Errorf("Got unexpected result: Got %d, Expected %d", total, 5)
	}
}
