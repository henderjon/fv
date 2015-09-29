package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"time"
)

var (
	ref  string
	help bool
	next, prev int
)

var api = "http://www.esvapi.org/v2/rest/passageQuery"

var opts = map[string]string{
	"key":                        "IP",
	"output-format":              "plain-text",
	"include-short-copyright":    "0",
	"include-copyright":          "0",
	"include-audio-link":         "0",
	"include-word-ids":           "0",
	"include-verse-ids":          "0",
	"include-headings":           "0",
	"include-subheadings":        "0",
	"include-footnote-links":     "0",
	"include-footnotes":          "0",
	"include-verse-numbers":      "1",
	"include-passage-references": "1",
}

func main() {

	flag.StringVar(&ref, "ref", "", "(optional) the reference to retrieve.")
	flag.BoolVar(&help, "help", false, "(optional) trigger the help")
	flag.IntVar(&next, "next", 0, "(optional) the verse for this week +x to view")
	flag.IntVar(&prev, "prev", 0, "(optional) the verse for this week -x to view")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	if ref == "" {
		week := diff()
		week += next
		week -= prev
		ref = Base[week]
	}

	fmt.Println("")
	fmt.Println(string(ping(ref)))
	fmt.Println("")
}

func ping(ref string) []byte {

	vals := &url.Values{}
	vals.Set("passage", ref)
	for k, v := range opts {
		vals.Set(k, v)
	}

	url, err := url.Parse(api)
	if err != nil {
		log.Fatal(err)
	}

	url.RawQuery = vals.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)

	// req.Header.Set("", "")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()
	return content

}

func diff() int {
	// Mon, 02 Jan 2006 15:04:05 -0700
	format := "Jan 2, 2006"
	t, err := time.Parse(format, "Jan 10, 2011")
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Floor(time.Since(t).Hours() / 24 / 7))
}
