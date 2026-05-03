package faces

import "testing"

func TestHello(t *testing.T) {
	txt := Faces("Hello :)")
	if txt != "Hello 🙂" {
		t.Errorf("%s should have a smile face emoticon", txt)
	}
}

func TestGoodbye(t *testing.T) {
	txt := Faces("Goodbye :(")
	if txt != "Goodbye 🙁" {
		t.Errorf("%s should have a frowning face emoticon", txt)
	}
}

func TestHelloGoodbye(t *testing.T) {
	txt := Faces("Hello :) Goodbye :(")
	if txt != "Hello 🙂 Goodbye 🙁" {
		t.Errorf("%s should have a frowning face and smila face emoticon", txt)
	}
}