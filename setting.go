package main

import "gopkg.in/yaml.v2"

var sampleYAML = `
directory:
  source: adf
  target: asdf

subscribe:
  - title: hitori
    subtitle:
      anissia: 123
      author: 코잔디
`

type Setting struct {
	Subscribe []Subscribe
	Directory Directory
}

type Directory struct {
	Source string
	Target string
}

type Subtitle struct {
	Anissia int
	Author  string
}

type Subscribe struct {
	Source   string
	Subtitle Subtitle
	Title    string
}

func ReadSetting() Setting {
	setting := Setting{}
	err := yaml.Unmarshal([]byte(sampleYAML), &setting)
	if err != nil {
		panic(err)
	}
	return setting
}
