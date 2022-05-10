package utils

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func AddIndex(bible *Bible, filename string) error {

	var temp_filename string = "." + filename
	err := UnZip(filename, temp_filename)
	if err != nil {
		return err
	}

	// Version 2 toc.nav
	var v2_index V2Index   // all-index
	var v2_navmap V2NavMap //books
	toc_file, _ := ioutil.ReadFile(temp_filename + "/EPUB/toc.ncx")
	err = xml.Unmarshal(FormatXML(toc_file), &v2_index)
	if err != nil {
		return err
	}
	for i := range bible.Books {
		var v2_navpoint []V2NavPoint //chapters
		n := bible.Books[i].Name
		for j := 0; j < (len(bible.Books[i].Chapters)); j++ {
			c := strconv.Itoa(j + 1)
			v2_navpoint = append(v2_navpoint, V2NavPoint{
				ID: c,
				V2NavLabel: V2NavLabel{
					Text: c,
				},
				Content: Content{
					Src: "xhtml/" + n + ".xhtml#" + c,
				},
			})
		}
		v2_navmap.V2NavPoint = append(v2_navmap.V2NavPoint, V2NavPoint{
			ID: n + ".xhtml",
			V2NavLabel: V2NavLabel{
				Text: Kyans[i],
			},
			Content: Content{
				Src: "xhtml/" + n + ".xhtml",
			},
			V2NavPoint: v2_navpoint,
		})
	}
	v2_index.V2NavMap = v2_navmap
	result, err := xml.MarshalIndent(v2_index, "", "    ")
	if err != nil {
		return err
	}
	result = append([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"), result...)
	err = ioutil.WriteFile(temp_filename+"/EPUB/toc.ncx", result, 0777)
	if err != nil {
		return err
	}

	// Version 3 nav.xhtml
	var v3_index V3Index
	var v3_ul V3Ul
	nav_file, _ := ioutil.ReadFile(temp_filename + "/EPUB/nav.xhtml")
	err = xml.Unmarshal(FormatXML(nav_file), &v3_index)
	if err != nil {
		return err
	}

	for i := range bible.Books {
		var v3_li_chapters []V3Li
		n := bible.Books[i].Name

		for j := 0; j < (len(bible.Books[i].Chapters)); j++ {
			c := strconv.Itoa(j + 1)
			v3_li_chapters = append(v3_li_chapters, V3Li{
				V3A: V3A{
					Text: c,
					Href: "xhtml/" + n + ".xhtml#" + c,
				},
			})
		}

		v3_ul.V3Li = append(v3_ul.V3Li, V3Li{
			V3A: V3A{
				Text: Kyans[i],
				Href: "xhtml/" + n + ".xhtml",
			},
			V3Ul: V3Ul{
				V3Li: v3_li_chapters,
			},
		})
	}

	v3_index.Body.Nav.V3Ul.V3Li = v3_ul.V3Li

	result, err = xml.MarshalIndent(v3_index, "", "    ")
	if err != nil {
		return err
	}
	result = append([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE html>\n"), result...)
	err = ioutil.WriteFile(temp_filename+"/EPUB/nav.xhtml", result, 0744)
	if err != nil {
		return err
	}

	err = Zip(temp_filename, filename)
	if err != nil {
		return err
	}
	os.RemoveAll(temp_filename)

	return err
}

func FormatXML(b []byte) []byte {
	var Replacer = strings.NewReplacer("&#xA;", "", "&#x9;", "", "\n", "", "\t", "")
	result := Replacer.Replace(string(b))
	return []byte(result)
}
