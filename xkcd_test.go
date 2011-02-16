package xkcd

import "testing"

const (
	title = "Good Code"
	num   = 844
	year  = "2011"
	month = "1"
	day   = "7"
	alt   = "You can either hang out in the Android Loop or the HURD loop."
	img   = "http://imgs.xkcd.com/comics/good_code.png"
)

func TestMain(t *testing.T) {
	c, err := Get(num)
	if err != nil {
		t.Errorf("Error %s returned by Get().", err)
	}
	if c.Title != title {
		t.Errorf("Title should be %s, got %s.", title, c.Title)
	}
	if c.Num != num {
		t.Errorf("Num should be %d, got %d.", num, c.Num)
	}
	if c.Year != year {
		t.Errorf("Year should be %s, got %s.", year, c.Year)
	}
	if c.Month != month {
		t.Errorf("Month should be %s, got %s.", month, c.Month)
	}
	if c.Day != day {
		t.Errorf("Day should be %s, got %s.", day, c.Day)
	}
	if c.Alt != alt {
		t.Errorf("Alt should be %s, got %s.", alt, c.Alt)
	}
	if c.Img != img {
		t.Errorf("Img should be %s, got %s.", img, c.Img)
	}
}

//The following comic can break unsanitized scripts:
const (
	wtitle = "("
	wnum   = 859
	wyear  = "2011"
	wmonth = "2"
	wday   = "11"
	walt   = `Brains aside, I wonder how many poorly-written xkcd.com-parsing scripts will break on this title (or \;;"\''{\<<[' this mouseover text."`
	wimg   = "http://imgs.xkcd.com/comics/(.png"
)

func TestWeird(t *testing.T) {
	c, err := Get(wnum)
	if err != nil {
		t.Errorf("Error %s returned by Get().", err)
	}
	if c.Title != wtitle {
		t.Errorf("Title should be %s, got %s.", wtitle, c.Title)
	}
	if c.Num != wnum {
		t.Errorf("Num should be %d, got %d.", wnum, c.Num)
	}
	if c.Year != wyear {
		t.Errorf("Year should be %s, got %s.", wyear, c.Year)
	}
	if c.Month != wmonth {
		t.Errorf("Month should be %s, got %s.", wmonth, c.Month)
	}
	if c.Day != wday {
		t.Errorf("Day should be %s, got %s.", wday, c.Day)
	}
	if c.Alt != walt {
		t.Errorf("Alt should be %s, got %s.", walt, c.Alt)
	}
	if c.Img != wimg {
		t.Errorf("Img should be %s, got %s.", wimg, c.Img)
	}
}
