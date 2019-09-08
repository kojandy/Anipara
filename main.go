package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var sampleYAML = `
directory:
  source: adf
  target: asdf

subscribe:
  - title: dumbbell
    subtitle:
      anissia: 4403
      author: 카이란
`

func main() {
	setting := Setting{}
	err := yaml.Unmarshal([]byte(sampleYAML), &setting)
	if err != nil {
		panic(err)
	}
	fmt.Println(setting.Subscribe[0].Subtitle.FindSubtitle())
}
