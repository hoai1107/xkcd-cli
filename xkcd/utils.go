package xkcd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func PrettyPrintJSON(v interface{}) error {
	output, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	fmt.Print(string(output))

	return nil
}

func PrintText(comic XkcdJSON) {
	date := fmt.Sprintf("%s/%s/%s", comic.Day, comic.Month, comic.Year)

	fmt.Printf("Publish date: %s\n"+
		"Title: %s\n"+
		"Description: %s\n"+
		"Image link: %s", date, comic.Title, comic.Alt, comic.Image)
}

func CreateImagesFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
}
