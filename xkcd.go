package xkcd

import (
	"http"
	"json"
	"os"
	"fmt"
	"io"
	"strings"
	"strconv"
)

func extractFilename(url string) string {
	splat := strings.Split(url, "/", -1)
	return splat[len(splat)-1]
}

type error string

func (e error) String() string { return string(e) }

const (
	Comic404 error = "There is no comic #404"
	NoSuch   error = "Nonexistant comic"
)

type Date struct {
	Year       int64
	Month, Day int
}

type Comic struct {
	num           uint
	title         string
	url, filename string
	alt           string
	date          Date
	img           io.ReadCloser
}

func Get(num uint) (*Comic, os.Error) {
	if num == 404 {
		return nil, Comic404
	}
	var url string
	if num == 0 {
		url = "http://xkcd.com/info.0.json"
	} else {
		url = fmt.Sprintf("http://xkcd.com/%d/info.0.json", num)
	}
	response, _, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 404 {
		return nil, NoSuch
	}
	decoder := json.NewDecoder(response.Body)
	var jsonVal interface{}
	err = decoder.Decode(&jsonVal)
	if err != nil {
		return nil, err
	}
	jsonMap := jsonVal.(map[string]interface{})
	comic := new(Comic)
	if num == 0 {
		comic.num = uint(jsonMap["num"].(float64))
	} else {
		comic.num = num
	}
	comic.title = jsonMap["title"].(string)
	comic.url = jsonMap["img"].(string)
	comic.filename = extractFilename(url)
	comic.alt = jsonMap["alt"].(string)
	year, _ := strconv.Atoi64(jsonMap["year"].(string))
	month, _ := strconv.Atoi(jsonMap["month"].(string))
	day, _ := strconv.Atoi(jsonMap["day"].(string))
	comic.date = Date{year, month, day}
	//Note that comic.img is intentionally left as nil to
	//avoid an extra download if only metadata are required.
	return comic, nil
}

func (c *Comic) Number() uint { return c.num }

func (c *Comic) Title() string { return c.title }

func (c *Comic) ImageURL() string { return c.url }

func (c *Comic) ImageFilename() string { return c.filename }

func (c *Comic) AltText() string { return c.alt }

func (c *Comic) Date() Date { return c.date }

//UpdateImage() re-downloads the comic image, in case it has changed
//(this is very unlikely, but it happens occasionally).
func (c *Comic) UpdateImage() {
	if c.img != nil {
		c.img.Close()
	}
	response, _, _ := http.Get(c.url)
	c.img = response.Body
}

//Image() returns the image; remember to Close() it when you're done with it.
func (c *Comic) Image() io.ReadCloser {
	if c.img == nil {
		c.UpdateImage()
	}
	return c.img
}
