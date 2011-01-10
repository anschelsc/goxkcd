package main

import (
	"http"
	"json"
	"fmt"
	"os"
)

type Comic struct {
	Img   string
	Title string
	Year  string
	Month string
	Day   string
	Num   int
	Alt   string
}

func Get(num int) (*Comic, os.Error) {
	response, _, err := http.Get(fmt.Sprintf("http://xkcd.com/%d/info.0.json", num))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	c := new(Comic)
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(c)
	return c, err
}
