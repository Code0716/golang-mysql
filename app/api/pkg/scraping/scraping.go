package scraping

import (
	"io/ioutil"
	"net/http"
)

//Scrape func スクレイピングする関数
func Scrape() string {
	// TODO スクレイピングパッケージをいれるところから。
	url := "http://example.com/"

	resp, err := http.Get(url)
	if err != nil {
		panic(err) //
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(byteArray)
}
