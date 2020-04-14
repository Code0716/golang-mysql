package scraping

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//scrape func スクレイピングする関数
func scrape() string {
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

//GetScrape TODO
//スクレイピング　TODO
func GetScrape(ginContext *gin.Context) {
	data := scrape()
	fmt.Println(data)
	//ginContext.JSON(http.StatusOK, data)

}
