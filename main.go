package main

import (
	"codeberg.org/peterzam/BurmeseBible/utils"
	"encoding/json"
	"io/ioutil"

	"github.com/bmaupin/go-epub"
)

func main() {

	EPUB_NAME := "BurmeseBible.epub"
	e := epub.NewEpub("Burmese Bible")
	cover_image_path, err := e.AddImage("src/cover.png", "cover.png")
	if err != nil {
		panic(err)
	}
	cover_css_path, err := e.AddCSS("src/cover.css", "")
	if err != nil {
		panic(err)
	}
	e.SetCover(cover_image_path, cover_css_path)
	e.SetAuthor("Adoniram Judson")
	e.SetDescription("1835 Judson Burmese Bible")
	e.SetLang("my")

	// Read bible
	file, err := ioutil.ReadFile("src/judson.json")
	if err != nil {
		panic(err)
	}

	var bible utils.Bible
	err = json.Unmarshal(file, &bible)
	if err != nil {
		panic(err)
	}

	// Make EPUB
	err = utils.MakeEpub(e, bible, EPUB_NAME)
	if err != nil {
		panic(err)
	}

	// Fix EPUB by adding index
	err = utils.AddIndex(&bible, EPUB_NAME)
	if err != nil {
		panic(err)
	}
}
