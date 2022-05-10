package utils

import "encoding/xml"

type Bible struct {
	Books []struct {
		Name     string `json:"name"`
		Chapters []struct {
			Verses []struct {
				Text string `json:"text"`
			} `json:"verses"`
		} `json:"chapters"`
	} `json:"books"`
}

type V2Index struct {
	XMLName xml.Name `xml:"ncx"`
	Xmlns   string   `xml:"xmlns,attr"`
	Version string   `xml:"version,attr"`
	Head    struct {
		Meta struct {
			Name    string `xml:"name,attr"`
			Content string `xml:"content,attr"`
		} `xml:"meta"`
	} `xml:"head"`
	DocTitle struct {
		Text string `xml:"text"`
	} `xml:"docTitle"`
	V2NavMap V2NavMap `xml:"navMap"`
}

type V2NavMap struct {
	XMLName    xml.Name     `xml:"navMap"`
	V2NavPoint []V2NavPoint `xml:"navPoint"`
}

type V2NavPoint struct {
	ID         string       `xml:"id,attr"`
	V2NavLabel V2NavLabel   `xml:"navLabel"`
	Content    Content      `xml:"content"`
	V2NavPoint []V2NavPoint `xml:"navPoint"`
}

type V2NavLabel struct {
	Text string `xml:"text"`
}

type Content struct {
	Src string `xml:"src,attr"`
}

type V3Index struct {
	XMLName xml.Name `xml:"html"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Epub    string   `xml:"epub,attr"`
	Head    struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
	} `xml:"head"`
	Body struct {
		Text string `xml:",chardata"`
		Nav  struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			H1   string `xml:"h1"`
			V3Ul V3Ul   `xml:"ul"`
		} `xml:"nav"`
	} `xml:"body"`
}

type V3Ul struct {
	V3Li []V3Li `xml:"li"`
}

type V3Li struct {
	V3A  V3A  `xml:"a"`
	V3Ul V3Ul `xml:"ul"`
}

type V3A struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
}
