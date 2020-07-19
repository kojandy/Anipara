package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Filename struct {
	title  string
	ep     int
	source string
}

func deduplicateWhitespace(s string) string {
	var buf strings.Builder
	var last rune
	for i, r := range s {
		if i == 0 || r != ' ' || r != last {
			buf.WriteRune(r)
			last = r
		}
	}
	return buf.String()
}

func parseFilename(filename string) (Filename, error) {
	filename = deduplicateWhitespace(filename)
	filename = regexp.MustCompile("^\\[.*\\] ").ReplaceAllString(filename, "")
	filename = regexp.MustCompile(".mp4$").ReplaceAllString(filename, "")
	title := regexp.MustCompile("(.*) -.*").ReplaceAllString(filename, "$1")
	ep, err := strconv.Atoi(regexp.MustCompile(".*- ([0-9]+).*").ReplaceAllString(filename, "$1"))
	if err != nil {
		return Filename{}, err
	}
	source := regexp.MustCompile(".*\\((.*?) .*").ReplaceAllString(filename, "$1")
	return Filename{title, ep, source}, nil
}

func Rename(src string, dst string) error {
	// srcDir, err := os.Open(src)
	// if err != nil {
	// 	return err
	// }
	// names, err := srcDir.Readdirnames(0)
	// if err != nil {
	// 	return err
	// }
	return nil
}
