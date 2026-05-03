package playback

import "strings"

func Playback(txt string) string {
	return strings.Replace(txt, " ", "...",-1)
}