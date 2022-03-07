package main

import (
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://nba.stats.qq.com/team/list.htm")
	panicErr(err)
	encode := determineEncoding(response.Body)
	reader := transform.NewReader(response.Body, encode.NewDecoder())
	all, err := ioutil.ReadAll(reader)
	panicErr(err)

	findNBATeams(all, encode.NewDecoder())
}
