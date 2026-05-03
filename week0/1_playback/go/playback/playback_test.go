package playback

import "testing"

func TestThis_is_CS50(t *testing.T) {
	txt := Playback("This is CS50")
	if txt != "This...is...CS50" {
		t.Errorf("%s should be playbacked", txt)
	}
}