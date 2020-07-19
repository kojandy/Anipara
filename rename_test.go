package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var exampleTitles = []string{
	"[Ohys-Raws] Princess Connect! ReDive  - 13 END (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive - 12 (BS11 1920x1080 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 12 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive - 11 (AT-X 1920x1080 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 11 (AT-X 1280x720 x264 AAC).mp4",
	"Princess Connect! ReDive  - 10 (BS11 1920x1080 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 10 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 09 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 08 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 07 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 06 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 05 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 04 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 03 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 02 (BS11 1280x720 x264 AAC).mp4",
	"[Ohys-Raws] Princess Connect! ReDive  - 01 (BS11 1280x720 x264 AAC).mp4",
}

var exampleParsed = []Filename{
	{title: "Princess Connect! ReDive", ep: 13, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 12, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 12, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 11, source: "AT-X"},
	{title: "Princess Connect! ReDive", ep: 11, source: "AT-X"},
	{title: "Princess Connect! ReDive", ep: 10, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 10, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 9, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 8, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 7, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 6, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 5, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 4, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 3, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 2, source: "BS11"},
	{title: "Princess Connect! ReDive", ep: 1, source: "BS11"},
}

func Test__deduplicateWhitespace(t *testing.T) {
	tests := [][]string{
		{"        ", " "},
		{"this   is      test", "this is test"},
		{" this is test ", " this is test "},
		{"tttt", "tttt"},
		{"tt    tt", "tt tt"},
	}

	for _, test := range tests {
		res := deduplicateWhitespace(test[0])
		if test[1] != res {
			t.Error("deduplicate failed:", test[0], "->", res)
		}
	}
}

func Test__parseFilename(t *testing.T) {
	for idx, filename := range exampleTitles {
		parsed, err := parseFilename(filename)
		if err != nil {
			t.Fatal(err)
		}

		if parsed != exampleParsed[idx] {
			t.Error("parse failed:", filename, parsed)
		}
	}
}

func TestRename(t *testing.T) {
	src, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(src)

	dst, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dst)

	for _, title := range exampleTitles {
		ioutil.WriteFile(filepath.Join(src, title), []byte{}, 664)
	}

	Rename(src, dst)
}
