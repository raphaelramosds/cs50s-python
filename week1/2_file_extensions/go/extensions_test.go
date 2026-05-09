package main

import "testing"

func TestJpg(t *testing.T) {
	mime := Extensions("happy.jpg")
	if mime != "image/jpeg" {
		t.Errorf("Extensions(\"happy.jpg\") = %q, want \"image/jpeg\"", mime)
	}
}

func TestGif(t *testing.T) {
	mime := Extensions("funny.gif")
	if mime != "image/gif" {
		t.Errorf("Extensions(\"funny.gif\") = %q, want \"image/gif\"", mime)
	}
}

func TestJpeg(t *testing.T) {
	mime := Extensions("photo.jpeg")
	if mime != "image/jpeg" {
		t.Errorf("Extensions(\"photo.jpeg\") = %q, want \"image/jpeg\"", mime)
	}
}

func TestPng(t *testing.T) {
	mime := Extensions("icon.png")
	if mime != "image/png" {
		t.Errorf("Extensions(\"icon.png\") = %q, want \"image/png\"", mime)
	}
}

func TestPdf(t *testing.T) {
	mime := Extensions("document.pdf")
	if mime != "application/pdf" {
		t.Errorf("Extensions(\"document.pdf\") = %q, want \"application/pdf\"", mime)
	}
}

func TestTxt(t *testing.T) {
	mime := Extensions("notes.txt")
	if mime != "text/plain" {
		t.Errorf("Extensions(\"notes.txt\") = %q, want \"text/plain\"", mime)
	}
}

func TestZip(t *testing.T) {
	mime := Extensions("archive.zip")
	if mime != "application/zip" {
		t.Errorf("Extensions(\"archive.zip\") = %q, want \"application/zip\"", mime)
	}
}

func TestCaseInsensitive(t *testing.T) {
	mime := Extensions("PHOTO.JPEG")
	if mime != "image/jpeg" {
		t.Errorf("Extensions(\"PHOTO.JPEG\") = %q, want \"image/jpeg\"", mime)
	}
}

func TestSpaceInsensitive(t *testing.T) {
	mime := Extensions("  photo.jpeg  ")
	if mime != "image/jpeg" {
		t.Errorf("Extensions(\"  photo.jpeg  \") = %q, want \"image/jpeg\"", mime)
	}
}

func TestUnknownExtension(t *testing.T) {
	mime := Extensions("unknown.xyz")
	if mime != "application/octet-stream" {
		t.Errorf("Extensions(\"unknown.xyz\") = %q, want \"application/octet-stream\"", mime)
	}
}