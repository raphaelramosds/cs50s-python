package indoor

import "testing"

func TestHELLO(t *testing.T) {
	txt := Indoor("HELLO")
	if txt != "hello" {
		t.Errorf("%s should be lowercase", txt)
	}
}

func TestTHIS_IS_CS50(t *testing.T) {
	txt := Indoor("THIS IS CS50")
	if txt != "this is cs50" {
		t.Errorf("%s should be lowercase", txt)
	}
}

func Test50(t *testing.T) {
	txt := Indoor("50")
	if txt != "50" {
		t.Errorf("%s should be lowercase", txt)
	}
}
