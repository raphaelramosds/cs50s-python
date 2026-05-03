package faces

import "strings"

func Faces(txt string) string {
	withSmileFace := strings.ReplaceAll(txt, ":)", "🙂");
	withFrowningFace := strings.ReplaceAll(withSmileFace, ":(", "🙁");
	return withFrowningFace;
}