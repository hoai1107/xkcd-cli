package xkcd

import "fmt"

const latestComic = 0

type XkcdJSON struct {
	Month      string `json:"month"`
	Number     int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Image      string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func BuildURL(comicNumber int) string {
	if comicNumber == latestComic {
		return "https://xkcd.com/info.0.json"
	}

	return fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicNumber)
}
