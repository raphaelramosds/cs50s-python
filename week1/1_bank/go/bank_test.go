package main

import "testing"

func TestHello(t *testing.T) {
	output := Bank("Hello")
	if output != "$0" {
		t.Error()
	}
}

func TestHelloComma(t *testing.T) {
	output := Bank("Hello,")
	if output != "$0" {
		t.Error()
	}
}

func TestHowYouDoing(t *testing.T) {
	output := Bank("How you doing?")
	if output != "$20" {
		t.Error(output)
	}
}

func TestWhatsHappening(t *testing.T) {
	output := Bank("What's happening?")
	if output != "$100" {
		t.Error(output)
	}
}
