package xkcd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func RunCLI() {
	var xkcdCLI CLI

	xkcdCLI.parseArgs()

	var comicJSON XkcdJSON
	url := BuildURL(xkcdCLI.comicNumber)
	if err := xkcdCLI.fetchComicJSON(url, &comicJSON); err != nil {
		log.Fatal(err)
	}

	if xkcdCLI.saveImage {
		err := xkcdCLI.getImage(&comicJSON)
		if err != nil {
			log.Fatal(err)
		}
	}

	if xkcdCLI.outputJSON {
		err := PrettyPrintJSON(comicJSON)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		PrintText(comicJSON)
	}

}

type CLI struct {
	hc          http.Client
	comicNumber int
	saveImage   bool
	outputJSON  bool
}

func (cli *CLI) parseArgs() {
	cli.hc = *http.DefaultClient

	//Get comic number
	flag.IntVar(&cli.comicNumber, "n", latestComic, "The number of comic to be fetched (default value is the latest comic)")

	//Set timeout
	flag.DurationVar(&cli.hc.Timeout, "t", 30*time.Second, "Time limit for request")

	//Whether to save the image
	flag.BoolVar(&cli.saveImage, "save", false, "Save image to current directory (default false)")

	//Print format
	flag.BoolVar(&cli.outputJSON, "json", false, "Output as JSON (default false), otherwise it will be printed as text")

	flag.Parse()
}

func (cli *CLI) fetchComicJSON(url string, data *XkcdJSON) error {
	response, err := cli.hc.Get(url)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("received %d response code", response.StatusCode)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}

func (cli *CLI) getImage(data *XkcdJSON) error {
	response, err := cli.hc.Get(data.Image)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("received %d response code", response.StatusCode)
	}

	CreateImagesFolder("images")
	imageFile, err := os.Create(fmt.Sprintf("images/%s.png", data.Title))
	if err != nil {
		return err
	}

	defer imageFile.Close()
	_, err = io.Copy(imageFile, response.Body)
	if err != nil {
		return err
	}

	return nil
}
