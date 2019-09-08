package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AnissiaSub struct {
	Ep     string `json:"s"`
	Time   string `json:"d"`
	Url    string `json:"a"`
	Author string `json:"n"`
}

func FindAnissia(id int, author string) string {
	queryUrl := fmt.Sprintf("https://www.anissia.net/anitime/cap?i=%d", id)
	resp, err := http.Get(queryUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	subs := []AnissiaSub{}
	err = json.Unmarshal(body, &subs)
	if err != nil {
		panic(err)
	}

	for _, sub := range subs {
		if sub.Author == author {
			return sub.Url
		}
	}

	panic("No target")
}
