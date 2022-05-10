package utils

import (
	"fmt"

	"github.com/bmaupin/go-epub"
)

func MakeEpub(e *epub.Epub, bible Bible, epub_file_out_path string) error {

	for _, book := range bible.Books {
		var text string
		for i, chapter := range book.Chapters {
			text = text + fmt.Sprintf("<p style=\"text-align:center;\" id=\"%d\">[%s]</p>", i+1, Number2Ganan(i+1))
			for j, verse := range chapter.Verses {
				text = text + fmt.Sprintf("<sup>[%s]</sup>\n%s\n", Number2Ganan(j+1), verse.Text)
			}
		}
		_, err := e.AddSection(text, book.Name, book.Name+".xhtml", "")
		if err != nil {
			return err
		}
	}
	e.Write(epub_file_out_path)
	return nil
}
