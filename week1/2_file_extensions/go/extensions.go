package main

import (
	"regexp"
	"strings"
)

func Extensions(filename string) string {
	// Reference: https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/MIME_types/Common_types
	filename = strings.TrimSpace(strings.ToLower(filename))
	re := regexp.MustCompile(`.*\.(\w+)$`)
	match := re.FindStringSubmatch(filename)
	mimetype := ""
	if len(match) > 0 {
		switch match[1] {
		case "jpg":
			mimetype = "image/jpeg"
		case "jpeg":
			mimetype = "image/jpeg"
		case "png":
			mimetype = "image/png"
		case "gif":
			mimetype = "image/gif"
		case "pdf":
			mimetype = "application/pdf"
		case "zip":
			mimetype = "application/zip"
		case "txt":
			mimetype = "text/plain"
		default:
			mimetype = "application/octet-stream"
		}
	}
	return mimetype
}