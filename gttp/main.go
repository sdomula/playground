package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

func main() {
	log.SetFlags(0)
	onlyHeader := flag.Bool("header", false, "only print the header")
	follow := flag.Bool("follow", false, "follow redirects")
	flag.Parse()

	var method, urlString string
	switch flag.NArg() {
	case 2:
		method = flag.Arg(0)
		urlString = flag.Arg(1)
	case 1:
		method = "GET"
		urlString = flag.Arg(0)
	default:
		flag.Usage()
		os.Exit(2)
	}

	url, err := url.ParseRequestURI(urlString)
	if err != nil {
		log.Fatalf("invalid url %s", urlString)
	}

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if !*follow {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	r, err := formatResponse(res, *onlyHeader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(r, "\n"))
}

func formatResponse(res *http.Response, onlyHeader bool) ([]string, error) {
	r := appendHeader([]string{}, res)
	if onlyHeader {
		return r, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, fmt.Errorf("could not read response body: %v", err)
	}
	for _, ct := range res.Header["Content-Type"] {
		if strings.Contains(ct, "application/json") {
			var pretty bytes.Buffer
			if err = json.Indent(&pretty, body, "", "\t"); err != nil {
				return r, fmt.Errorf("could not pretty print json: %v", err)
			}
			body = pretty.Bytes()
			break
		}
	}
	r = append(r, string(body))
	return r, nil
}

func appendHeader(s []string, res *http.Response) []string {
	for k, v := range res.Header {
		s = append(s, fmt.Sprintf("%s: %s", k, strings.Join(v, " ")))
	}
	sort.Strings(s)
	h := []string{fmt.Sprintf("%s %s", res.Proto, res.Status)}
	s = append(h, s...)
	return append(s, "\n")
}
