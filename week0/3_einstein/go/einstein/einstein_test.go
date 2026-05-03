package einstein

import (
	"testing"
)

func Test1(t *testing.T) {
	e := Einstein(1)
	eExp := 90000000000000000
	if e != uint64(eExp) {
		t.Errorf("%d should be %d", e, eExp)
	}
}

func Test14(t *testing.T) {
	e := Einstein(14)
	eExp := 1260000000000000000
	if e != uint64(eExp) {
		t.Errorf("%d should be %d", e, eExp)
	}
}

func Test50(t *testing.T) {
	e := Einstein(50)
	eExp := 4500000000000000000
	if e != uint64(eExp) {
		t.Errorf("%d should be %d", e, eExp)
	}
}